<script lang="ts">
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

<div>{@html data.items[33].first_name}</div>

{#each data.items as item (item.user_id)}
<ul class="mt-4 pl-6 text-lg">
	<li class="font-bold"> first_name: {item.first_name} </li>
	<li class="font-bold">last_name: {item.last_name}</li>
	<li class="font-bold">email: {item.email}</li>
	<li class="font-bold">cell: {item.cell}</li>
</ul>
{/each}
