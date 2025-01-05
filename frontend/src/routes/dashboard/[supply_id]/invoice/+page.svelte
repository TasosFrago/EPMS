<script lang="ts">
	import type { Invoice } from '$lib/components/InvoiceList.svelte';
	import type { PopupStoreT } from '$lib/stores';
	import { getContext } from 'svelte';
	import type { PageData } from './$types';
	import { PopupStatus } from '$lib/types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import InvoiceList from '$lib/components/InvoiceList.svelte';

	let { data }: { data: PageData } = $props();
	let loading = $state(true);
	let invoiceList: Invoices = $state([]);

	if (data && !data.error && data.invoices) {
		invoiceList.push(...data.invoices);
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
	<InvoiceList {invoiceList} />
{/if}
