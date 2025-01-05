import type { LayoutServerLoad } from "./$types";
import { type HttpError, error, redirect } from "@sveltejs/kit";

import { apiUrl } from "$lib/settings";
import { UnauthorizedUserError, InternalServerError, EmptyCookie } from "$lib/types";


export interface UserData {
	user_id: number;
	first_name: string;
	last_name: string;
}

export const load: LayoutServerLoad = async ({ locals }): Promise<{ loadData: UserData }> => {
	let resUserData: UserData;
	//const msg = "You are not logged in. Please log in or sign up"

	//const token = cookies.get("jwt")
	//if (token == null) {
	//	//throw new EmptyCookie('You are not logged in. Please log in or sign up', { flag: true, path: "/" })
	//	//redirect(307, "/");
	//	//throw error(400, "Empty cookies")
	//	//return { loadData: [null, ErrorHandler.REDIRECT] }
	//}
	console.log("I am herre")

	try {
		if (locals.error) {
			throw new UnauthorizedUserError("Unauthorized User", { flag: true, path: "/" })
		}
		if (locals.user && locals.token) {
			const userData = await getConsumerName(locals.token, locals.user.user_id)
			resUserData = {
				user_id: locals.user.user_id,
				first_name: userData.first_name,
				last_name: userData.last_name
			}
			console.log("This is the res User data " + JSON.stringify(resUserData));
			return { loadData: resUserData }
		}
		throw new InternalServerError("Unexpected error", { flag: true, path: "/" })
	} catch (err) {
		console.log("there is error inside the dashboard")
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
	//try {
	//	const response = await fetch(apiUrl("/auth/me"), {
	//		method: 'GET',
	//		headers: {
	//			'Authorization': 'Bearer ' + token
	//		}
	//	});
	//	const data = await response.json();
	//	if (response.ok && data.user_type == 0) {
	//		const userDatad = await getConsumerName(token, await data.user_id);
	//		if (!userDatad) {
	//			throw new Error("Failed to fetch user data")
	//		}
	//		resUserData = {
	//			user_id: await data.user_id,
	//			first_name: await userDatad.first_name,
	//			last_name: await userDatad.last_name
	//		};
	//		return { loadData: resUserData };
	//	}
	//} catch (err) {
	//	console.log("This is the error" + err);
	//	//cookies.set('errorMsg', msg, {
	//	//	httpOnly: true,
	//	//	secure: true,
	//	//	maxAge: 60,
	//	//	path: '/',
	//	//})
	//	//popupStore.set(msg);
	//	//redirect(307, "/");
	//	throw new Error("Unexpected error");
	//	//return { loadData: [null, ErrorHandler.INTERNAL_SERVER_ERROR] }
	//}
	////return { loadData: [null, ErrorHandler.INTERNAL_SERVER_ERROR] }
	//throw new Error("Unexpected error");
	//redirect(307, "/");
}

const getConsumerName = async (token: string, user_id: number) => {
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
				throw new UnauthorizedUserError(data.warning, { flag: false, path: "" });
			case 500: // Internal Server Status
				throw new InternalServerError(data.error, { flag: false, path: "" });
			default:
				throw new Error(`Unexpected status ${response.status}: ${data.message}`)
		}
	}
}
