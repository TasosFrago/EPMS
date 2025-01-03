<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { LayoutData } from './$types';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();
	console.log(data.userData.user_id);
</script>

<div class="flex h-screen">
	<nav
		class="flex flex-col justify-between bg-gradient-to-bl from-[#FF4000] via-[#1F0A70] via-[#C34623] via-65% via-[26%] to-[#28057B]"
	>
		<ul>
			<li>
				<h1 class="sm:text-md mx-2 my-5 text-center text-xl font-bold lg:text-2xl">
					Welcome <br />
					{data.userData.first_name + ' ' + data.userData.last_name}!
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

<style>
</style>
