<script module>
</script>

<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		show: boolean;
		children: Snippet;
	}

	let { show, children }: Props = $props();
</script>

{#if show}
	<div
		class="modal fixed left-0 top-0 z-50 flex h-full w-full items-center justify-center p-8 lg:p-0"
	>
		<!-- Background opacity -->
		<div class="modal-overlay fixed h-full w-full bg-gray-900 opacity-50"></div>

		<div
			class="z-50 mx-auto overflow-y-auto rounded-lg bg-white shadow-xl lg:max-h-full lg:max-w-md"
		>
			<div class="head relative flex items-center bg-white px-8 py-5">
				<div class="mr-4 flex-grow">
					{@render children?.()}
				</div>
				<button
					class="absolute right-2 top-2 ml-7 rounded-full bg-gray-200 p-2 hover:bg-gray-300"
					onclick={() => {
						show = false;
					}}
					aria-label="Close"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						height="15px"
						width="15px"
						viewBox="0 0 24 24"
						fill="#000000"
						><path d="M0 0h24v24H0V0z" fill="none" /><path
							d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12 19 6.41z"
						/></svg
					>
				</button>
			</div>
		</div>
	</div>
{/if}
