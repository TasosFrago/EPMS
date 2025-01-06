<script lang="ts">
	import { goto } from '$app/navigation';

	const goLogin = () => {
		goto('/login/consumer');
	};
	const goSignUp = () => {
		goto('/signup/consumer');
	};
</script>

<h1 class="mt-10 text-center text-3xl font-bold text-black">Welcome Consumer!</h1>

<div
	class="w-md mx-3 mt-10 flex h-64 min-h-full max-w-md flex-col justify-center space-y-8 rounded-2xl bg-gray-200 px-6 py-7 sm:mx-auto sm:w-full lg:px-8"
>
	<button
		class="mt-3 w-full rounded-xl bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
		onclick={() => goLogin()}>Login</button
	>
	<span class="mt-4 self-center">or</span>
	<button
		class="mt-5 w-full rounded-xl bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
		onclick={() => goSignUp()}
	>
		Sign Up
	</button>
</div>
