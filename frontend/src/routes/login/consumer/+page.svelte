<script lang="ts">
	import { writable, type Writable } from 'svelte/store';
	import Login, { type LogInData } from '$lib/components/Login.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { apiUrl, debugLog } from '$lib/settings';
	import { CookieManager, type CookieOptions } from '$lib';

	let data: Writable<LogInData> = writable({
		email: '',
		password: ''
	});
	let errorMessage = $state('');
	let showPopup = $state(false);

	const handleLogin = async (): Promise<void> => {
		const loginData = $data;
		debugLog(loginData);
		try {
			const response = await fetch(apiUrl('/auth/login/consumer'), {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(loginData)
			});

			if (response.ok) {
				const responseData = await response.json();
				console.log(responseData);
				debugLog('Login successful:' + responseData);
				showPopup = false;

				const token = responseData.data.token;
				console.log(token);
				CookieManager.set('jwt', token, {
					maxAge: 24 * 60 * 60,
					secure: true,
					sameSite: 'Lax'
				} as CookieOptions);
				window.location.href = '/dashboard';
			} else {
				switch (response.status) {
					case 401: // Unauthorized
						errorMessage = 'Invalid Email. Please try again.';
						break;
					case 500: // Internal Server Error
						errorMessage = 'Server error. Please try again';
						break;
					default:
						errorMessage = 'Unexpected error happend. Please try again';
						break;
				}
				showPopup = true;
			}
		} catch (error) {
			console.error('Error:', error);
			errorMessage = 'Network error. Please retry later';
			showPopup = true;
		}
	};
</script>

<Modal show={showPopup}>
	<div class=" text-center text-xl">
		{errorMessage}
	</div>
</Modal>

<Login title={'Login Consumer'} {data} handleSubmit={handleLogin} />
