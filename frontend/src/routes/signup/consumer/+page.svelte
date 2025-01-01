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

{#snippet EyeON()}
<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
    <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z" />
    <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
  </svg>
{/snippet}

{#snippet EyeOff()}
<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
    <path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88" />
  </svg>
{/snippet}

<Modal show={showPopup}>
	<div class=" text-center text-xl">
		{errorMessage}
	</div>
</Modal>

<h1 class="mx-auto mt-5 text-center text-3xl font-bold">Sign Up As Consumer</h1>

<div class="flex justify-center">
	<form onsubmit={handleSignUp} class="form-container bg-gray-200">
		<label for="name" class="pl-1">First Name</label>
		<input
			type="text"
			id="name"
			bind:value={$formData.first_name}
			placeholder="Enter your name"
			required
		/>

		<label for="last_name" class="pl-1">Last Name</label>
		<input
			type="text"
			id="last_name"
			bind:value={$formData.last_name}
			placeholder="Enter your last name"
			required
		/>

		<label for="email" class="pl-1">Email</label>
		<input
			type="email"
			id="email"
			bind:value={$formData.email}
			placeholder="Enter your email"
			required
		/>

		<label for="password" class="password-label pl-1">Password</label>
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
				{#if seePassword}
                    {@render EyeON()}
                {:else}
                    {@render EyeOff()}
                {/if}

			</button>
		</div>

		<label for="cell" class="pl-1">Cell</label>
		<input
			type="tel"
			id="cell"
			bind:value={$formData.cell}
			placeholder="Enter your number"
			required
		/>

		<label for="landline" class="pl-1">Landline</label>
		<input
			type="tel"
			id="landline"
			bind:value={$formData.landline}
			placeholder="Enter your landline"
		/>

		<button
			type="submit"
			aria-label="submit"
			class="w-full rounded-md bg-blue-600 px-5 py-2 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
		>
			Sign Up
		</button>
	</form>
</div>

<style>
	.form-container {
		max-width: 28rem;
		width: 90%;
		margin: 3%;
		padding: 2rem;
		border: 1px solid #fdfcfc;
		border-radius: 16px;
		/*background: #d3d0d0;*/
	}

	.password-container {
		display: flex;
		align-items: center;
        justify-content: center;
		flex-direction: row;
		margin: 0 auto;
	}

	.password-label {
		display: block;
		margin-bottom: 10px;
		font-weight: bold;
		border-radius: 8px;
	}

	.password-input {
		width: 80%;
		flex-grow: 4;
		margin-right: 10px;
		border: 1px solid #181717;
		border-radius: 8px;
		padding: 10px;
	}

	.toggle-password {
		position: relative;
		margin-bottom: 16px;
		width: 15%;
		height: 15%;
		background: none;
		color: #007bff;
		border: 1px solid #181717;
		border-radius: 8px;
		cursor: pointer;
		text-align: center;
		padding: 0.6em;
		display: flex;
		align-items: center;
		justify-content: center;
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
		color: white;
		border: none;
		border-radius: 8px;
		/* margin-top: 16px; */
		cursor: pointer;
	}

	button:hover {
		background-color: #295ecf;
	}
</style>
