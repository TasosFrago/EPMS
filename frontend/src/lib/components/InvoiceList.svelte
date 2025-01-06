<script lang="ts" module>
	export interface Invoice {
		invoice_id: string;
		provider: string;
		issue_date: string;
		expiry_date: string;
		current_cost: number;
		paid_amount: number;
		is_paid: boolean;
	}
	//export type Invoices = Invoice[];
</script>

<script lang="ts">
	let {
		invoiceList,
		handleClick,
		paymentClick
	}: {
		invoiceList: Invoice[];
		handleClick: (invoice_id: number) => void;
		paymentClick: (current_cost: number, providerName: string) => void;
	} = $props();
</script>

{#snippet verify()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-5 text-green-500"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
		/>
	</svg>
{/snippet}

{#snippet BackArrow()}
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
			d="m12.75 15 3-3m0 0-3-3m3 3h-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
		/>
	</svg>
{/snippet}

<div class="m-1 h-[336px] w-fit overflow-y-auto border-[2.5px] border-black">
	<table class="max-w-lg table-auto">
		<thead class="sticky top-0 max-w-lg border-separate bg-[#DDDFE0]">
			<tr class="h-[48px]" style="box-shadow: inset 0px -2px 2px rgba(0, 0, 0, 1);">
				<th class="px-2 text-center">#ID</th>
				<th class="px-2 text-center">Issued</th>
				<th class="px-2 text-center">Expires</th>
				<th class="px-2 text-center">Cost</th>
				<th class="px-2 text-center">Paid Amount</th>
				<th class="px-2 text-center"></th>
				<th class="px-2 text-center">Pay</th>
			</tr>
		</thead>
		<tbody class="divide-y divide-gray-300">
			{#each invoiceList as invoice}
				<tr class="h-[48px] odd:bg-white even:bg-[#F4EFEF] hover:bg-gray-300">
					<td
						class="px-3 text-center text-sm"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{invoice.invoice_id}
					</td>
					<td
						class="px-3 text-center text-sm"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{invoice.issue_date}
					</td>
					<td
						class="px-3 text-center text-sm"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{invoice.expiry_date}
					</td>
					<td
						class="min-w-28 whitespace-nowrap px-3 text-center text-sm"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{invoice.current_cost} €
					</td>
					<td
						class="min-w-28 whitespace-nowrap px-3 text-center text-sm"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{invoice.paid_amount} €
					</td>
					<td
						class="w-3 px-3 text-center"
						role="button"
						onclick={() => {
							handleClick(parseInt(invoice.invoice_id));
						}}
						title="Pay invoice"
					>
						{#if invoice.is_paid}
							{@render verify()}
						{:else}
							-
						{/if}
					</td>
					<td
						class="w-3 px-3 text-center"
						role="button"
						onclick={() => {
							paymentClick(invoice.current_cost, invoice.provider);
						}}
					>
						{#if !invoice.is_paid}
							{@render BackArrow()}
						{:else}
							-
						{/if}
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
