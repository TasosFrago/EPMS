<script lang="ts">
	import type { Invoice } from '$lib/components/InvoiceList.svelte';
	import type { PopupStoreT } from '$lib/stores';
	import { getContext } from 'svelte';
	import type { PageData } from './$types';
	import { PopupStatus, UnauthorizedUserError } from '$lib/types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import InvoiceList from '$lib/components/InvoiceList.svelte';
	import { page } from '$app/state';
	import { apiUrl } from '$lib/settings';
	import { CookieManager } from '$lib';
	import Modal from '$lib/components/Modal.svelte';
	//import type { Provider } from './+page.ts';
	import { writable, type Writable } from 'svelte/store';

	interface Provider {
		name: string;
		phone?: string;
		email?: string;
	}

	let { data }: { data: PageData } = $props();

	let loading = $state(true);
	let invoiceList: Invoice[] = $state([]);
	let user_id = $state(0);

	let selectedProvider: Provider | null = $state(null);

	const Supply_id = page.params.supply_id.padStart(5, '0');
	console.log(page.params);

	if (data.user_id) user_id = data.user_id;

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
	const handleClick = (invoice_id: number) => {
		if (browser && data.user_id) getInvoice(user_id, invoice_id);
		invoiceDetailsShow = true;
	};

	const paymentClick = (current_cost: number, providerName: string) => {
		console.log('Provider' + providerName);
		selectedProvider = { name: providerName };
	};

	interface PaymentData {
		providerName: string;
		amount: number;
	}
	const formData: Writable<PaymentData> = writable({
		providerName: '',
		amount: 0
	});

	interface InvoiceDetails {
		invoice_id: string;
		total: number;
		current_cost: number;
		receiver: number;
		meter: number;
		provider: string;
		name: string;
		month: string;
		year: number;
	}

	let invoiceDetails: InvoiceDetails | null = $state(null);
	let invoiceDetailsShow = $state(false);

	const SubmitPayment = async () => {
		const formD = $formData;
	};

	const getInvoice = async (user_id: number, invoice_id: number) => {
		console.log(user_id);
		try {
			const token = CookieManager.get('jwt');
			if (token == null) {
				throw new UnauthorizedUserError('Empty Cookies', { flag: true, path: '/' });
			}
			const response = await fetch(apiUrl(`/consumer/${user_id}/invoice/${invoice_id}`), {
				method: 'GET',
				headers: {
					Authorization: 'Bearer ' + token
				}
			});
			const dataInv = await response.json();
			if (response.ok) {
				invoiceDetails = {
					invoice_id: dataInv.invoice_id.toString().padStart(6, '0'),
					total: dataInv.total,
					current_cost: dataInv.current_cost,
					receiver: dataInv.receiver,
					meter: dataInv.meter,
					provider: dataInv.provider,
					name: dataInv.name,
					month: dataInv.month,
					year: dataInv.year
				};
			}
		} catch (err) {
			console.error(err);
		}
	};

	const convertToDate = (month: string, year: number): string => {
		const monthNames = [
			'January',
			'February',
			'March',
			'April',
			'May',
			'June',
			'July',
			'August',
			'September',
			'October',
			'November',
			'December'
		];
		const monthIndex = monthNames.indexOf(month);
		console.log(month);
		const date = new Date(year, monthIndex);
		console.log(date);
		const formatedDate = date.toISOString().slice(0, 10).replace(/-/g, '/');
		return formatedDate;
	};
</script>

{#if invoiceDetails}
	<Modal
		show={invoiceDetailsShow}
		ShowHandler={() => {
			invoiceDetailsShow = false;
		}}
	>
		<div class="max-w-full flex-grow">
			<ul class="m-2">
				<li class="flex flex-row justify-between text-xl">
					<span class="mr-6">
						Invoice ID: <span class="font-bold text-blue-600">{invoiceDetails?.invoice_id}</span>
					</span>
					<span>{convertToDate(invoiceDetails.month, invoiceDetails.year)}</span>
				</li>
				<hr class="my-2 mr-1 border-t-[1.5px] border-black" />

				<li class="flex flex-row justify-between text-xl">
					<span class="mr-4">Plan:</span>
					<span>{invoiceDetails.name}</span>
				</li>
				<li class="flex flex-row justify-between text-xl">
					<span>Provider:</span>
					<span>{invoiceDetails.provider}</span>
				</li>
				<li class="flex flex-row justify-between text-xl">
					<span>Total:</span>
					<span>{invoiceDetails.total} €</span>
				</li>
				<li class="flex flex-row justify-between text-xl">
					<span>Current Cost:</span>
					<span>{invoiceDetails.current_cost} €</span>
				</li>
			</ul>
		</div>
	</Modal>
{/if}

{#snippet BackArrow()}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="size-8"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="m11.25 9-3 3m0 0 3 3m-3-3h7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
		/>
	</svg>
{/snippet}

{#if loading}
	<!-- FIX: Add better loading here -->
	<h1>Loading...</h1>
{:else}
	<div class="flex flex-col overflow-y-auto">
		<div class=" mt-28 flex flex-col p-3 md:flex-row">
			<div class="w-full">
				<div class=" flex flex-row items-center">
					<button
						class="transition-transform hover:-translate-x-1 hover:-translate-y-1 hover:text-red-500 hover:shadow-black"
						onclick={() => {
							history.back();
						}}
					>
						{@render BackArrow()}
					</button>
					<div class="ml-8 border-l-[9px] border-l-red-500">
						<div
							class="inline-block border-[3px] border-l-0 border-black px-2 py-[0.8px] text-2xl font-bold text-[#FF0000]"
						>
							{Supply_id}
						</div>
					</div>
				</div>

				<hr class="my-2 mr-1 border-t-[1.5px] border-black" />

				<div class="overflow-x-scroll p-2 md:overflow-x-hidden">
					<InvoiceList {invoiceList} {handleClick} {paymentClick} />
				</div>
			</div>
			<div class="w-full">
				<div class="ml-10 border-l-[9px] border-l-green-600">
					<div
						class="inline-block border-[3px] border-l-0 border-black px-2 py-[0.8px] text-2xl font-bold"
					>
						Pay
					</div>
				</div>

				<hr class="my-2 mr-1 border-t-[1.5px] border-black" />

				<div class="p-3">
					<form onsubmit={SubmitPayment}>
						<div class="flex flex-row items-center justify-between">
							<label for="amount-to-pay" class="text-lg">Amount to pay</label>
							<div class="mr-5 lg:mr-10">
								<input
									class="mx-2 mt-1 h-8 w-20 rounded-lg"
									type="number"
									bind:value={$formData.amount}
									required
								/>
								<span>€</span>
							</div>
						</div>

						<div class="mt-4 flex flex-row items-center justify-between">
							<label for="provider" class="text-lg">Select Provider</label>

							<select
								class="mr-5 flex h-8 appearance-none items-center rounded-lg text-left align-top leading-4 lg:mr-10"
								bind:value={selectedProvider}
							>
								{#if data.providers}
									{#each data.providers as provider}
										<option class="mb-5 h-6 text-center text-sm" value={provider}>
											{provider.name}
										</option>
									{/each}
								{/if}
							</select>
						</div>
						<button
							type="submit"
							aria-label="submit"
							class="mt-4 w-full rounded-xl bg-blue-600 px-5 py-2 font-bold text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
						>
							Pay Provider
						</button>
					</form>
				</div>
			</div>
		</div>
		<div class="mt-20 w-full p-3">
			<div class="ml-10 border-l-[9px] border-l-green-600">
				<div
					class="inline-block border-[3px] border-l-0 border-black px-2 py-[0.8px] text-2xl font-bold"
				>
					Plan
				</div>
			</div>

			<hr class="my-2 mr-1 border-t-[1.5px] border-black" />

			<div class="overflow-x-auto p-2">
				<InvoiceList {invoiceList} {handleClick} {paymentClick} />
			</div>
		</div>
	</div>
{/if}
