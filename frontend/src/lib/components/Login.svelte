<script module lang="ts">
	export interface LogInData {
		email: string;
		password: string;
	}
</script>

<script lang="ts">
	import type { Writable } from 'svelte/store';

	interface Props {
		title: string;
		data: Writable<LogInData>;
		handleSubmit: () => Promise<void>;
	}

	let { title, data, handleSubmit }: Props = $props();
	let showPassword: boolean = $state(false);
</script>

{#snippet EyeIconON()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-6"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"
		/>
		<path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
	</svg>
{/snippet}

{#snippet EyeIconOFF()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-6"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88"
		/>
	</svg>
{/snippet}

<h1 class="mx-auto my-5 text-center text-3xl font-bold">{title}</h1>

<div
	class="mt-10 flex min-h-full max-w-md flex-col justify-center px-6 py-12 sm:mx-auto sm:w-full sm:max-w-sm lg:px-8"
>
	<form onsubmit={handleSubmit} class="space-y-4">
		<div class="mb-6">
			<label for="email" class="block text-sm/6 font-medium text-gray-900">Email</label>
			<input
				type="email"
				id="email"
				placeholder="Enter your email"
				bind:value={$data.email}
				class="mt-2 w-full rounded-lg border border-gray-300 py-2 pl-3 outline-none focus:ring-indigo-600"
				required
			/>
		</div>
		<div class="relative mb-4">
			<label for="password" class="mb-1 block text-sm font-medium text-gray-700">Password</label>
			<input
				type={showPassword ? 'text' : 'password'}
				id="password"
				placeholder="Enter your password"
				bind:value={$data.password}
				class="mt-2 w-full rounded-lg border border-gray-300 py-2 pl-3 outline-none focus:ring-indigo-600"
				required
			/>
			<button
				type="button"
				onclick={() => (showPassword = !showPassword)}
				aria-label="toggle-password-visibility"
				class="insert-y-0 absolute right-0 px-3 text-gray-500 hover:text-gray-700 focus:outline-none"
			>
				{#if showPassword}
					{@render EyeIconOFF()}
				{:else}
					{@render EyeIconON()}
				{/if}
			</button>
		</div>

		<button
			type="submit"
			class="w-full rounded-md bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
			aria-label="submit"
		>
			Login
		</button>
	</form>
</div>
