<script lang="ts">
	import '../app.css';
	import Popup from '$lib/components/Popup.svelte';
	import { type PopupStoreT } from '$lib/stores';
	import { PopupStatus } from '$lib/types';
	import { onDestroy, onMount, setContext, type Snippet } from 'svelte';
	import type { LayoutData } from './$types';
	import { goto, invalidateAll } from '$app/navigation';

	import { dev } from '$app/environment';
	import { injectAnalytics } from '@vercel/analytics/sveltekit';
	import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';

	injectAnalytics({ mode: dev ? 'development' : 'production' });
	injectSpeedInsights();

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
		console.log(data);
		if (data.error) {
			console.log('We have error');
			setPopup({
				show: true,
				msg: data.error.msg,
				status: PopupStatus.ERROR
			});
			if (data.error.shouldRedirect && data.error.shouldRedirect.flag)
				goto(data.error.shouldRedirect.path);
		}

		const handlePopState = () => {
			invalidateAll();
		};
		window.addEventListener('popstate', handlePopState);
		onDestroy(() => {
			window.removeEventListener('popstate', handlePopState);
		});
	});
</script>

<Popup show={popup.show} color={popup.status}>
	<div>
		<span>{popup.msg}</span>
	</div>
</Popup>

{@render children()}
