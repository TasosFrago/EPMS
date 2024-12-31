<script lang="ts">
	import { writable, type Writable } from 'svelte/store';
	import { apiUrl, debugLog } from '$lib/settings';
	import Modal from '$lib/components/Modal.svelte';

	interface SignUpData {
		first_name: string;
		last_name: string;
		password: string;
		email: string;
		cell: string;
		landline?: string;
	}

	const formData: Writable<SignUpData> = writable({
		first_name: '',
		last_name: '',
		password: '',
		email: '',
		cell: '',
		landline: ''
	});
	let seePassword = $state(false);
	let errorMessage = $state('');
	let showPopup = $state(false);

	const handleSignUp = async (): Promise<void> => {
		const data = $formData;
		debugLog(data);
		try {
			const response = await fetch(apiUrl('/auth/signup/consumer'), {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(data)
			});

			if (response.ok) {
				debugLog('Sign-up successful:' + (await response.json()));
				showPopup = false;
			} else {
				switch (response.status) {
					case 400: // Bad Request
						errorMessage = 'Invalid data given. Check your input';
						break;
					case 409: // Conflict
						errorMessage = 'Consumer with this data exists already';
						break;
					case 500: // Internal Server error
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

<h1 class="mx-auto my-5 text-center text-3xl font-bold">Sign Up As Consumer</h1>

<div class="flex justify-center">
	<div class="form-container">
		<label for="name">First Name</label>
		<input
			type="text"
			id="name"
			bind:value={$formData.first_name}
			placeholder="Enter your name"
			required
		/>

		<label for="last_name">Last Name</label>
		<input
			type="text"
			id="last_name"
			bind:value={$formData.last_name}
			placeholder="Enter your last name"
			required
		/>

		<label for="password" class="password-label">Password</label>
		<div class="password-container">
			<input
				type={seePassword ? 'text' : 'password'}
				id="password"
				bind:value={$formData.password}
				placeholder="Enter your password"
				class="password-input"
				required
			/>
			<button
				type="button"
				class="toggle-password"
				onclick={() => (seePassword = !seePassword)}
				aria-label="Toggle password visibility"
			>
				{seePassword ? 'Hide' : 'show'}
			</button>
		</div>

		<label for="email">Email</label>
		<input
			type="email"
			id="email"
			bind:value={$formData.email}
			placeholder="Enter your email"
			required
		/>

		<label for="cell">Cell</label>
		<input
			type="tel"
			id="cell"
			bind:value={$formData.cell}
			placeholder="Enter your number"
			required
		/>

		<label for="landline">Landline</label>
		<input
			type="tel"
			id="landline"
			bind:value={$formData.landline}
			placeholder="Enter your landline"
		/>

		<button type="submit" onclick={handleSignUp}>Sign Up</button>
	</div>
</div>

<style>
	.form-container {
		max-width: 400px;
		width: 90%;
		margin: 3%;
		padding: 16px;
		border: 1px solid #fdfcfc;
		border-radius: 16px;
		background: #d3d0d0;
	}

	.password-container {
		display: grid-template-columns;
		align-items: center;
		margin: 0 auto;
	}

	.password-label {
		display: block;
		margin-bottom: 0px;
		font-weight: bold;
		border-radius: 8px;
	}

	.password-input {
		width: 80%;
		flex-grow: 1;
		margin-right: 10px;
		padding: 10px;
		border: 1px solid #181717;
		border-radius: 8px;
	}

	.toggle-password {
		width: 15%;
		background: none;
		color: #007bff;
		border: 1px solid #181717;
		border-radius: 8px;
		padding: 10px;
		cursor: pointer;
		display: inline-flex;
		justify-content: center;
		flex-shrink: 0;
	}

	.toggle-password:hover {
		background-color: #84abc5;
	}

	.toggle-password:focus {
		outline: none;
		box-shadow: 0 0 4px #613e5e;
	}

	label {
		display: block;
		margin-bottom: 10px;
		font-weight: bold;
		border-radius: 8px;
	}

	input {
		width: 100%;
		padding: 10px;
		margin-bottom: 16px;
		border: 1px solid #181717;
		border-radius: 8px;
	}

	button {
		width: 100%;
		padding: 12px;
		background-color: #007bff;
		color: white;
		border: none;
		border-radius: 8px;
		margin-top: 16px;
		cursor: pointer;
	}

	button:hover {
		background-color: #295ecf;
	}
</style>
