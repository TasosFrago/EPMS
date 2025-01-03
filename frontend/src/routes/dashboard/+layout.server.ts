import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
//import { type Cookies } from "@sveltejs/kit";

import { popupStore } from "$lib/stores";
import { apiUrl, debugLog } from "$lib/settings";
import { UnauthorizedUserError, InternalServerError } from "$lib/types";

export interface UserData {
	user_id: number;
	first_name: string;
	last_name: string;
}

export const load: LayoutServerLoad = async ({ cookies }): Promise<{ userData: UserData }> => {
	const msg: string = "You need to Log in";
	let resUserData: UserData;

	const token = cookies.get("jwt")
	if (token == null) {
		popupStore.set(msg);
		throw redirect(307, "/");
	}

	try {
		const response = await fetch(apiUrl("/auth/me"), {
			method: 'GET',
			headers: {
				'Authorization': 'Bearer ' + token
			}
		});
		const data = await response.json();
		if (response.ok && data.user_type == 0) {
			const userDatad = await getConsumerName(token, await data.user_id);
			if (!userDatad) {
				throw new Error("Failed to fetch user data")
			}
			debugLog({
				user_id: data.user_id,
			})
			resUserData = {
				user_id: await data.user_id,
				first_name: await userDatad.first_name,
				last_name: await userDatad.last_name
			};
			return { userData: resUserData };
		} else {
			if (response.status == 401) {
				popupStore.set(msg);
				throw redirect(307, "/");
			}
		}
	} catch (error) {
		console.error(error);
		popupStore.set(msg);
		throw redirect(307, "/");
	}
	throw new Error("Unexpected end of function in load.")
}

const getConsumerName = async (token: string, user_id: number) => {
	try {
		const response = await fetch(apiUrl("/consumer/" + user_id + "/"), {
			method: 'GET',
			headers: {
				'Authorization': 'Bearer ' + token
			}
		});
		const data = await response.json();
		if (response.ok) {
			return {
				first_name: data.first_name,
				last_name: data.last_name
			};
		} else {
			switch (response.status) {
				case 401: // Unauthorized
					throw new UnauthorizedUserError(data.warning, data.code);
				case 500: // Internal Server Status
					throw new InternalServerError(data.error, data.code);
				default:
					throw new Error(`Unexpected status ${response.status}: ${data.message}`)
			}
		}
	} catch (error) {
		console.error(error);
		throw new Error("Error in getConsumerName: " + error);
	}
}
