<script lang="ts">
	import { type Snippet } from 'svelte';
	import type { LayoutData } from './$types';
	import { type UserData } from './+layout.server';
	import { goto } from '$app/navigation';
	import { type Result, ErrorHandler } from '$lib/errorTypes';
	import { browser } from '$app/environment';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();
	let user: UserData = $state({
		user_id: 0,
		first_name: '',
		last_name: ''
	});
	let loading = $state(true);
	const [loadData, err]: Result<UserData, ErrorHandler> = data.loadData;
	if (loadData) user = loadData;
	if (err != null) {
		console.log(err);
		if (browser) goto('/');
	}
	$effect(() => {
		if (!err) loading = false;
	});
</script>

{#if loading}
	<!-- #FIX: Make a better loading module -->
	<h1>Loading...</h1>
{:else}
	<div class="flex h-screen">
		<nav
			class="flex flex-col justify-between bg-[linear-gradient(to_bottom_left,_#FF4000_0%,_#C34623_24%,_#1F0A70_69%,_#28057B_100%)]"
		>
			<ul>
				<li>
					<h1 class="sm:text-md mx-2 my-5 text-center text-xl font-bold lg:text-2xl">
						Welcome <br />
						{user.first_name + ' ' + user.last_name}!
					</h1>
				</li>
			</ul>
			<ul>
				<li><h3 class="mb-4 text-center text-white">Sign Out</h3></li>
			</ul>
		</nav>

		<div class="flex-1 overflow-auto">
			{@render children?.()}
		</div>
	</div>
{/if}
