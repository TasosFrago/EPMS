<script lang="ts">
	import '../app.css';
	import Popup from '$lib/components/Popup.svelte';
	import { type PopupStoreT } from '$lib/stores';
	import { PopupStatus } from '$lib/types';
	import { onMount, setContext, type Snippet } from 'svelte';
	import type { LayoutData } from './$types';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();

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

	onMount(() => {
		if (data.msg) {
			popup.show = true;
			popup.msg = data.msg;
			popup.status = PopupStatus.ERROR;
		}
	});
</script>

<Popup show={popup.show} color={popup.status}>
	<div>
		<span>{popup.msg}</span>
	</div>
</Popup>

{@render children()}
