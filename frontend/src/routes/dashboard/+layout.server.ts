import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { apiUrl } from "$lib/settings";
import { UnauthorizedUserError, InternalServerError } from "$lib/types";


export interface UserData {
	user_id: number;
	first_name: string;
	last_name: string;
}

export const load: LayoutServerLoad = async ({ cookies }): Promise<{ loadData: UserData }> => {
	let resUserData: UserData;
	const msg = "You are not logged in. Please log in or sign up"

	const token = cookies.get("jwt")
	if (token == null) {
		cookies.set('errorMsg', msg, {
			httpOnly: true,
			secure: true,
			maxAge: 60,
			path: '/',
		})
		redirect(307, "/");
		//error(400, "Empty cookies")
		//return { loadData: [null, ErrorHandler.REDIRECT] }
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
			resUserData = {
				user_id: await data.user_id,
				first_name: await userDatad.first_name,
				last_name: await userDatad.last_name
			};
			return { loadData: resUserData };
		}
	} catch (err) {
		console.log("This is the error" + err);
		cookies.set('errorMsg', msg, {
			httpOnly: true,
			secure: true,
			maxAge: 60,
			path: '/',
		})
		//popupStore.set(msg);
		redirect(307, "/");
		//error(500, "Unexpected error");
		//return { loadData: [null, ErrorHandler.INTERNAL_SERVER_ERROR] }
	}
	//return { loadData: [null, ErrorHandler.INTERNAL_SERVER_ERROR] }
	//error(500, "Unexpected error");
	redirect(307, "/");
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
