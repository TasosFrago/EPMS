<script lang="ts">
	import MeterCard, { type MeterCardInfo } from '$lib/components/MeterCard.svelte';
	//import { type PopupStoreT } from '$lib/stores';
	//import { PopupStatus } from '$lib/types';
	//import { getContext } from 'svelte';

	//const setPopup: (p: PopupStoreT) => void = getContext('popup');
	//setPopup({
	//	show: true,
	//	msg: 'Heloooooooooo',
	//	status: PopupStatus.ERROR
	//});

	let meterCardInfo: MeterCardInfo = $state({
		supply_id: '',
		address: '',
		date: '',
		amount: '',
		is_paid: false
	});
</script>

<div class="ml-4">
	<div class="ml-10 mt-28 border-l-[9px] border-l-blue-500">
		<div
			class="inline-block border-[3px] border-l-0 border-black px-2 py-[0.8px] text-2xl font-bold"
		>
			Meters
		</div>
	</div>

	<hr class="my-2 mr-6 border-t-[1.5px] border-black" />

	<MeterCard
		supply_id={meterCardInfo.supply_id}
		address={meterCardInfo.address}
		date={meterCardInfo.date}
		amount={meterCardInfo.amount}
		is_paid={meterCardInfo.is_paid}
	/>
</div>
