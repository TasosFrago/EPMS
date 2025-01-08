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

	interface Plan {
		ID: number;
		Type: string;
		Price: number;
		Name: string;
		Provider: string;
		IssueDate: Date;
		Duration: number;
	}

	let { data }: { data: PageData } = $props();

	let loading = $state(true);
	let invoiceList: Invoice[] = $state([]);
	let user_id = $state(0);

	//let selectedProvider: Provider | null = $state(null);
	let selectProvider: boolean = $state(false);

	const Supply_id = page.params.supply_id.padStart(5, '0');
	console.log(page.params);

	if (data.user_id) user_id = data.user_id;

	const setPopup: (p: PopupStoreT) => void = getContext('popup');
	if (data && !data.error && data.invoices) {
		invoiceList.push(...data.invoices);
		loading = false;
	} else if (data.error) {
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
		$formData.providerName = providerName;
		$formData.amount = current_cost;
		selectProvider = true;
		console.log('Provider ' + $formData);
	};

	interface PaymentData {
		providerName: string;
		amount: number;
	}
	const formData: Writable<PaymentData> = writable({
		providerName: '',
		amount: 1
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
		console.log(formD);
		try {
			const token = CookieManager.get('jwt');
			if (token == null) {
				setPopup({
					show: true,
					msg: 'Unauthorized user',
					status: PopupStatus.ERROR
				});
				if (browser) goto('/');
				return;
			}
			const response = await fetch(apiUrl(`/consumer/${user_id}/meters/${Supply_id}/pays/`), {
				method: 'POST',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					provider: formD.providerName,
					amount: formD.amount
				})
			});
			if (response.ok) {
				setPopup({
					show: true,
					msg: 'Completed Payment to ' + formD.providerName,
					status: PopupStatus.SUCCESS
				});
				return;
			} else {
				let errorMessage: string;
				let redirect: boolean = false;
				switch (response.status) {
					case 401: // Unauthorized
						errorMessage = 'Invalid Email. Please try again.';
						redirect = true;
						break;
					case 400: // Unauthorized
						errorMessage = 'Invalid data given.';
						break;
					case 500: // Internal Server Error
						errorMessage = 'Server error. Please try again';
						break;
					default:
						errorMessage = 'Unexpected error happend. Please try again';
						break;
				}
				setPopup({
					show: true,
					msg: errorMessage,
					status: PopupStatus.ERROR
				});
				if (browser && redirect) goto('/');
			}
		} catch (error) {
			console.error(error);
			setPopup({
				show: true,
				msg: 'Unexpected error occurred',
				status: PopupStatus.ERROR
			});
			if (browser) goto('/');
			return;
		}
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
			<div class="mt-16 w-full md:mt-0">
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
									class="mx-2 mt-1 h-8 w-32 rounded-lg"
									type="number"
									bind:value={$formData.amount}
									min="1"
									defaultvalue="1"
									required
								/>
								<span>€</span>
							</div>
						</div>

						<div class="mt-4 flex flex-row items-center justify-between">
							<label for="provider" class="text-lg">Select Provider</label>

							{#if !selectProvider}
								<select
									class="ml-4 mr-5 flex h-8 appearance-none items-baseline justify-center rounded-lg text-left align-top leading-4 lg:mr-10"
									bind:value={$formData.providerName}
									required
								>
									{#if data.providers}
										{#each data.providers as provider}
											<option class="mb-5 h-6 text-left text-sm" value={provider.name}>
												{provider.name}
											</option>
										{/each}
									{/if}
								</select>
							{:else}
								<div
									class=" mr-5 flex h-8 w-full items-baseline justify-center rounded-lg border-[2px] border-sky-600 pl-2 text-left text-sm lg:mr-10"
								>
									<span class="text-md mt-1 h-full w-full text-left">{$formData.providerName}</span>
								</div>
							{/if}
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

			{#if data.plans && data.plans.length != 0}
				<span>Select plan</span>
				<!-- <div class="mb-10 flex flex-col items-center justify-center"> -->
				<!-- 	<h1 class="m-4 text-lg">You have no plan selected. Please select plan:</h1> -->
				<!-- 	{#each data.plans as plan} -->
				<!-- 		<button class="h-8 w-40 rounded-xl bg-blue-600 text-center font-bold text-white"> -->
				<!-- 			Select plan -->
				<!-- 			<ul> -->
				<!-- 				<li>Id</li> -->
				<!-- 			</ul> -->
				<!-- 		</button> -->
				<!-- 	{/each} -->
				<!-- </div> -->
			{:else}
				<h1 class="m-4 text-lg">You have already selected a plan for the current month</h1>
			{/if}
		</div>
	</div>
{/if}
