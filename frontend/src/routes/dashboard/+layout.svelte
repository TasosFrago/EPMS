<script lang="ts">
	import { getContext, onDestroy, onMount, type Snippet } from 'svelte';
	import type { LayoutData } from './$types';
	import { type UserData } from './+layout.server';
	import type { PopupStoreT } from '$lib/stores';
	import { PopupStatus } from '$lib/types';
	import { goto } from '$app/navigation';
	import { browser } from '$app/environment';
	import { CookieManager } from '$lib';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();
	let user: UserData = $state({
		user_id: 0,
		first_name: '',
		last_name: ''
	});
	let loading = $state(true);
	let showNav = $state(true);

	console.log(data);
	if (data && !data.error) {
		user = data.loadData;
		loading = false;
	} else if (data.error) {
		const setPopup: (p: PopupStoreT) => void = getContext('popup');
		setPopup({
			show: true,
			msg: data.error.msg,
			status: PopupStatus.ERROR
		});
		if (data.error.shouldRedirect && data.error.shouldRedirect.flag && browser)
			goto(data.error.shouldRedirect.path);
	}
	// Function to check screen width and update showNav
	const updateNavVisibility = () => {
		showNav = window.innerWidth >= 1024; // Auto-open if screen is lg (1024px) or wider
	};

	// Run when component mounts
	onMount(() => {
		if (browser) {
			updateNavVisibility(); // Set initial state based on screen width
			window.addEventListener('resize', updateNavVisibility); // Listen for screen resizes
		}
	});

	// Cleanup event listener when component is destroyed
	onDestroy(() => {
		if (browser) window.removeEventListener('resize', updateNavVisibility);
	});

	//const [loadData, err]: Result<UserData, ErrorHandler> = data.loadData;
	//if (loadData) user = loadData;
	//if (err != null) {
	//	console.log(err);
	//	if (browser) goto('/');
	//}
	//$effect(() => {
	//	if (!err) loading = false;
	//});
</script>

{#snippet HamburgerBtn()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-7"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
		/>
	</svg>
{/snippet}

{#if loading}
	<!-- #FIX: Make a better loading module -->
	<h1>Loading...</h1>
{:else}
	<div class="flex h-screen">
		<nav
			class="flex flex-col bg-[linear-gradient(to_bottom_left,_#FF4000_0%,_#C34623_24%,_#1F0A70_69%,_#28057B_100%)] transition-all duration-300 {showNav
				? 'absolute z-50 h-full w-full'
				: 'w-9'} lg:relative lg:h-full lg:w-56"
		>
			<button
				id="hamburger-btn"
				title="Back"
				class="self-center pt-2 text-center text-white lg:hidden"
				onclick={() => {
					showNav = !showNav;
				}}
			>
				{@render HamburgerBtn()}
			</button>

			<ul class="{showNav ? 'block' : 'hidden'} flex flex-1 flex-col justify-between">
				<li>
					<h1 class="sm:text-md mx-2 my-5 text-center text-xl font-bold lg:text-2xl">
						Welcome <br />
						{user.first_name + ' ' + user.last_name}!
					</h1>
				</li>
				<li class="px-5">
					<button
						class="mb-4 w-full items-center rounded-xl bg-sky-600 p-1 text-center font-bold text-white hover:bg-red-500"
						onclick={() => CookieManager.delete('jwt')}
					>
						Sign Out
					</button>
				</li>
			</ul>
		</nav>

		<div class="flex-1 overflow-auto">
			{@render children?.()}
		</div>
	</div>
{/if}
