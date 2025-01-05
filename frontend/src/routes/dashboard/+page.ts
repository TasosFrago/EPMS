import { apiUrl } from '$lib/settings';
import { InternalServerError, UnauthorizedUserError } from '$lib/types';
import type { PageServerLoad } from './$types';
import type { MeterCardInfo } from '$lib/components/MeterCard.svelte';


export interface MeterInvoiceInfo {
	supply_id: number;
	address: string;
	invoices: {
		date: string;
		amount: number;
		is_paid: boolean;
	}[];
}

export interface Meter {
	supply_id: number;
	address: string;
}
export interface Invoice {
	meter: number;
	current_cost: number;
	issue_date: string;
	expiry_date: string;
	is_paid: boolean;
}

export const load: PageServerLoad = async ({ data }) => {

	const locals = await data.localsD;
	try {
		if (locals.error) {
			throw new UnauthorizedUserError("Unauthorized User", { flag: true, path: "/" })
		}
		if (locals.user && locals.token) {
			const meters = await getMeterList(locals.token, locals.user.user_id);
			const invoices = await getInvoiceList(locals.token, locals.user.user_id);
			const combinedList: MeterInvoiceInfo[] = meters.map((supply: Meter) => {
				return {
					...supply,
					invoices: invoices
						.filter((invoice: Invoice) => invoice.meter === supply.supply_id)
						.map((invoiceD: {
							expiry_date: string;
							current_cost: number;
							is_paid: boolean;
						}) => ({
							date: invoiceD.expiry_date,
							amount: invoiceD.current_cost,
							is_paid: invoiceD.is_paid ?? false
						}))
						.sort((a: { date: string, amount: number, is_paid: boolean }, b: { date: string, amount: number, is_paid: boolean }) => (new Date(b.date).getTime() - new Date(a.date).getTime()))
				}
			});

			const finalList: MeterCardInfo[] = combinedList
				.map(({ supply_id, address, invoices }) => {
					//if (!invoices || invoices.length === 0) return null;

					const latestInvoice = invoices[0];

					return {
						supply_id: supply_id.toString().padStart(5, '0'),
						address: address,
						date: latestInvoice ? latestInvoice.date.replace(/-/g, '/') : null,
						amount: latestInvoice ? latestInvoice.amount.toFixed(2) : null,
						is_paid: latestInvoice ? latestInvoice.is_paid : null
					};
				})
			//.filter(Boolean)

			return { meterData: finalList };
		}
	} catch (err) {
		console.error(err)
		if (locals.error) {
			console.log("Inside layout dashboard " + locals.error.msg);
			return { error: locals.error }
		} else {
			return {
				error: {
					status: 500,
					msg: "Unexpected Error",
					shouldRedirect: {
						flag: true,
						path: "/"
					}
				}
			}
		}
	}

}

const getMeterList = async (token: string, user_id: number) => {
	const response = await fetch(apiUrl("/consumer/" + user_id + "/meters/"), {
		method: 'GET',
		headers: {
			'Authorization': 'Bearer ' + token
		}
	});
	const data = await response.json()
	if (response.ok) {
		return data
			.map(({ supply_id, address }: Meter) => ({ supply_id, address }));
	} else {
		switch (response.status) {
			case 401: // Unauthorized
				throw new UnauthorizedUserError(data.warning, { flag: false, path: "" });
			case 500:
				throw new InternalServerError(data.error, { flag: false, path: "" });
		}
	}
	throw new InternalServerError("Unexpected Error", { flag: false, path: "" });
}

const getInvoiceList = async (token: string, user_id: number) => {
	const response = await fetch(apiUrl("/consumer/" + user_id + "/invoices/"), {
		method: 'GET',
		headers: {
			'Authorization': 'Bearer ' + token
		}
	});
	const data = await response.json()
	if (response.ok) {
		return data
			.map(({ meter, current_cost, issue_date, expiry_date, is_paid }: Invoice) => ({ meter, current_cost, issue_date, expiry_date, is_paid }));
	} else {
		switch (response.status) {
			case 401: // Unauthorized
				throw new UnauthorizedUserError(data.warning, { flag: false, path: "" });
			case 500:
				throw new InternalServerError(data.error, { flag: false, path: "" });
		}
	}
	throw new InternalServerError("Unexpected Error", { flag: false, path: "" });
}
