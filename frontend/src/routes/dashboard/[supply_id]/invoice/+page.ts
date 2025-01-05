import { apiUrl } from '$lib/settings';
import { UnauthorizedUserError, InternalServerError } from '$lib/types';
import type { PageServerLoad } from './$types';
import type { Invoice } from '$lib/components/InvoiceList.svelte'

export const load: PageServerLoad = async ({ data, params }) => {
	console.log("Inside page Load")
	const locals = await data.localsD;
	console.log("locals not in server :" + JSON.stringify(locals, null, 2))
	console.log("Url params: " + JSON.stringify(params, null, 2))
	try {
		if (locals.error) {
			throw new UnauthorizedUserError("Unauthorized User", { flag: true, path: "/" })
		}
		if (locals.user && locals.token) {
			const invoices: Invoice[] = await getInvoiceList(locals.token, locals.user.user_id, params.supply_id);
			return { invoices: invoices }

		}
	} catch (err) {
		console.log("Error in page")
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

const getInvoiceList = async (token: string, user_id: number, supply_id: number): Promise<Invoice[]> => {
	console.log("User_id: " + user_id);
	console.log("supply_id: " + supply_id);
	const response = await fetch(apiUrl(`/consumer/${user_id}/meters/${supply_id}/invoices/`), {
		method: 'GET',
		headers: {
			'Authorization': 'Bearer ' + token
		}
	});
	const data = await response.json();
	if (response.ok) {
		return data
			.map(({
				invoice_id,
				issue_date,
				expiry_date,
				current_cost,
				total_paid,
				is_paid
			}: {
				invoice_id: number;
				issue_date: string;
				expiry_date: string;
				current_cost: number;
				total_paid: number;
				is_paid: boolean;
			}) => {
				return {
					invoice_id: invoice_id.toString().padStart(6, '0'),
					issue_date: issue_date.replace(/-/g, '/'),
					expiry_date: expiry_date.replace(/-/g, '/'),
					current_cost: current_cost.toFixed(2).toString(),
					paid_amount: total_paid.toFixed(2).toString(),
					is_paid: is_paid,
				}
			})
			.sort((a: Invoice, b: Invoice) => new Date(b.issue_date).getTime() - new Date(a.issue_date).getTime())
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
