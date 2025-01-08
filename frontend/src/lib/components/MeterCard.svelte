<script lang="ts" module>
	export interface MeterCardInfo {
		supply_id: string;
		address: string;
		date?: string;
		amount?: string;
		is_paid?: boolean;
	}
</script>

<script lang="ts">
	let {
		meterInfo,
		handleClick,
		handleButtonClick
	}: { meterInfo: MeterCardInfo; handleClick: () => void; handleButtonClick: () => void } =
		$props();

	function OnClickF() {
		if (!meterInfo.is_paid) {
			handleButtonClick();
		}
	}
</script>

{#snippet PowerIcon()}
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
			d="m3.75 13.5 10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75Z"
		/>
	</svg>
{/snippet}

<div
	class="custom-shadow mb-2 ml-1 mt-2 flex min-h-[350px] w-[225px] transform flex-col rounded-2xl border-[2px] border-black px-3 py-3 transition-transform hover:-translate-x-1 hover:-translate-y-1 hover:shadow-black sm:max-w-[225px]"
>
	<div
		class="top-header"
		role="button"
		tabindex="0"
		onclick={handleClick}
		onkeyup={handleClick}
		title="invoice"
	>
		<ul>
			<li
				class="text-md flex flex-row items-center text-[#FF0000] underline-offset-4 hover:underline"
			>
				{@render PowerIcon()}
				<span>{meterInfo.supply_id}</span>
			</li>
			<li class="text-black underline-offset-4 hover:underline">{meterInfo.address}</li>
		</ul>
	</div>
	<div
		class="relative h-full w-full"
		role="button"
		tabindex="0"
		onclick={handleClick}
		onkeyup={handleClick}
		title="invoice"
	>
		<div class="absolute bottom-12 w-full">
			<div class=" flex max-h-fit flex-col justify-end px-1">
				<div class="middle w-full border-l-[2px] border-l-black px-1">
					<ul>
						<li class="text-md mb-3 font-bold">Invoice</li>
						<li class="mb-1 flex flex-row justify-between text-sm">
							<span>Expires:</span>
							<span>{meterInfo.date ? meterInfo.date : '-'}</span>
						</li>
						<li class="flex flex-row justify-between text-sm">
							<span>Amount:</span>
							<span class="font-bold">{meterInfo.amount ? meterInfo.amount + '€' : '-'}</span>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
	<div class="bottom-footer mx-1 w-full self-center">
		<!-- FIX: Add pay redirect path -->
		<button
			class="w-full items-center rounded-xl p-1 font-bold text-white"
			style="background: {meterInfo.is_paid ? 'green' : 'black'}"
			onclick={OnClickF}
			title="pay"
		>
			{meterInfo.is_paid == null ? '-' : meterInfo.is_paid ? 'Paid' : 'Pay'}
		</button>
	</div>
</div>

<style>
	.custom-shadow {
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease;
	}

	.custom-shadow:hover {
		box-shadow:
			rgba(0, 0, 0, 0.4) 3px 3px,
			rgba(0, 0, 0, 0.3) 6px 6px;
		/*rgba(240, 46, 170, 0.4) 3px 3px,*/
		/*rgba(240, 46, 170, 0.3) 6px 6px;*/
		/*transform: translateX(-0.5rem) translateY(-0.5rem);*/
	}
</style>
