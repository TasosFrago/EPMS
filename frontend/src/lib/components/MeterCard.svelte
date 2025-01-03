<script lang="ts" module>
	export interface MeterCardInfo {
		supply_id: string;
		address: string;
		date: string;
		amount: string;
		is_paid: boolean;
	}
</script>

<script lang="ts">
	let { supply_id, address, date, amount, is_paid }: MeterCardInfo = $props();
</script>

{#snippet PowerIcon()}
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
			d="m3.75 13.5 10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75Z"
		/>
	</svg>
{/snippet}

<div
	class="flex min-h-[350px] max-w-[225px] flex-col justify-between rounded-2xl border-[2px] border-black px-3 py-3"
>
	<div class="top-header">
		<ul>
			<li class="text-sm text-[#FF0000]">
				{@render PowerIcon()}
				<span>{supply_id}</span>
			</li>
			<li class="text-black">{address}</li>
		</ul>
	</div>
	<div class="middle ml-1 border-l-[2px] border-l-black px-1 pr-2">
		<ul>
			<li class="text-md mb-3 font-bold">Invoice</li>
			<li class="mb-1 flex flex-row justify-between text-sm">
				<span>Expires:</span>
				<span>{date}</span>
			</li>
			<li class="flex flex-row justify-between text-sm">
				<span>Amount:</span>
				<span class="font-bold">{amount} €</span>
			</li>
		</ul>
	</div>
	<div class="bottom-footer mx-1">
		<!-- TODO: Add pay redirect path -->
		<button
			class="w-full items-center rounded-xl p-1 font-bold text-white"
			style="background: {is_paid ? 'green' : 'black'}"
		>
			{is_paid ? 'Paid' : 'Pay'}
		</button>
	</div>
</div>
