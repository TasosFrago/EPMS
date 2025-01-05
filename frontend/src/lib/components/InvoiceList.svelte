<script lang="ts" module>
	export interface Invoice {
		invoice_id: string;
		issue_date: string;
		expiry_date: string;
		current_cost: number;
		paid_amount: number;
		is_paid: boolean;
	}
	//export type Invoices = Invoice[];
</script>

<script lang="ts">
	let { invoiceList }: { invoiceList: Invoice[] } = $props();
</script>

{#snippet verify()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-4"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
		/>
	</svg>
{/snippet}

<div class="m-1 h-[336px] w-fit overflow-y-auto border-[2.5px] border-black">
	<table class="max-w-lg table-auto">
		<thead
			class="sticky top-0 z-10 max-w-lg border-separate border-[2.5px] border-black bg-[#DDDFE0]"
			style="border: 2.5px solid black;"
		>
			<tr class="h-[48px]">
				<th class="px-2 text-center">#ID</th>
				<th class="px-2 text-center">Issued</th>
				<th class="px-2 text-center">Expires</th>
				<th class="px-2 text-center">Cost</th>
				<th class="px-2 text-center">Paid Amount</th>
				<th class="px-2 text-center"></th>
			</tr>
		</thead>
		<tbody class="divide-y divide-gray-300">
			{#each invoiceList as invoice}
				<tr class="h-[48px] odd:bg-white even:bg-[#F4EFEF]">
					<td class="px-2 text-center">{invoice.invoice_id}</td>
					<td class="px-2 text-center">{invoice.issue_date}</td>
					<td class="px-2 text-center">{invoice.expiry_date}</td>
					<td class="px-2 text-center">{invoice.current_cost}</td>
					<td class="px-2 text-center">{invoice.paid_amount}</td>
					<td class="px-2 text-center">
						{#if invoice.is_paid}
							{@render verify()}
						{:else}
							-
						{/if}
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
