<script lang="ts">
	import '../app.css';
	import Popup from '$lib/components/Popup.svelte';
	import { type PopupStoreT } from '$lib/stores';
	import { PopupStatus } from '$lib/types';
	import { setContext } from 'svelte';

	let { children } = $props();

	let popup: PopupStoreT = $state({
		show: false,
		msg: '',
		status: PopupStatus.INFO
	});

	function setPopup(p: PopupStoreT) {
		popup = p;

		setTimeout(() => {
			popup.show = false;
		}, 5000);
	}
	setContext('popup', setPopup);
</script>

<Popup show={popup.show} color={popup.status}>
	<div>
		<span>{popup.msg}</span>
	</div>
</Popup>

{@render children()}
