<script module>
</script>

<script lang="ts">
	import { type Snippet } from 'svelte';
	import { fade } from 'svelte/transition';
	import { PopupStatus } from '$lib/types';

	interface Props {
		show: boolean;
		color?: PopupStatus;
		children: Snippet;
	}

	let { show, color, children }: Props = $props();
</script>

{#if show}
	<div
		class="fixed right-4 top-4 ml-2 rounded-lg p-4 shadow-md transition-all"
		style="background-color: {color || PopupStatus.INFO};"
		role="alert"
	>
		<div
			class="text-white"
			style="text-shadow: 1px 1px 1.5px black, -1px -1px 1.5px black, 1px -1px 1.5px black, -1px 1px 1.5px black;"
			transition:fade
		>
			{@render children?.()}
		</div>
	</div>
{/if}
