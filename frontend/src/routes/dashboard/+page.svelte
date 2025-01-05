<script lang="ts">
	import MeterCard, { type MeterCardInfo } from '$lib/components/MeterCard.svelte';
	import type { PageData } from './$types';
	import { type PopupStoreT } from '$lib/stores';
	import { PopupStatus } from '$lib/types';
	import { getContext } from 'svelte';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';

	//const setPopup: (p: PopupStoreT) => void = getContext('popup');
	//setPopup({
	//	show: true,
	//	msg: 'Heloooooooooo',
	//	status: PopupStatus.ERROR
	//});

	let { data }: { data: PageData } = $props();

	let loading = $state(true);
	let meterCardInfo: MeterCardInfo[] = $state([]);

	function ClickHandler(is_paid_button: boolean, supply_id: string): () => void {
		const invoiceUrl = `/dashboard/${parseInt(supply_id)}/invoice`;
		const paymentUrl = `/dashboard/${parseInt(supply_id)}/payment`;
		if (!is_paid_button && browser) {
			return () => {
				goto(invoiceUrl);
			};
		} else if (is_paid_button && browser) {
			return () => {
				goto(paymentUrl);
			};
		}
		console.error('Undefined is_paid flag');
		return () => {};
	}

	const meterInfo = data.meterData;
	if (data && !data.error && meterInfo) {
		meterCardInfo.push(...meterInfo);
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
</script>

{#if loading}
	<!-- FIX: Add better loading here -->
	<h1>Loading...</h1>
{:else}
	<div class="ml-4">
		<div class="ml-10 mt-28 border-l-[9px] border-l-blue-500">
			<div
				class="inline-block border-[3px] border-l-0 border-black px-2 py-[0.8px] text-2xl font-bold"
			>
				Meters
			</div>
		</div>

		<hr class="my-2 mr-6 border-t-[1.5px] border-black" />

		<div class="scrollbar-hide scrollbar-hide flex w-full flex-row space-x-2 overflow-x-auto">
			<div class="flex flex-nowrap space-x-2">
				{#each meterCardInfo as meter}
					<MeterCard
						meterInfo={meter}
						handleClick={ClickHandler(false, meter.supply_id)}
						handleButtonClick={ClickHandler(true, meter.supply_id)}
					/>
				{/each}
			</div>
		</div>
	</div>
{/if}

<style>
	/* Hide scrollbar */
	.scrollbar-hide::-webkit-scrollbar {
		display: none;
	}

	/* Firefox */
	.scrollbar-hide {
		scrollbar-width: none;
	}
</style>
