import { isRedirect, redirect, type Redirect, type HandleServerError } from '@sveltejs/kit';
import { type CustomHttpError } from '$lib/types';
import { apiUrl } from '$lib/settings';

console.log("hook loaded")
//let should: { flag: boolean, path: string };
//const count: number = 0;
let err: {
	status: number;
	msg: string;
	shouldRedirect: {
		flag: boolean;
		path: string;
	}
};

//export const handleError: HandleServerError = async ({ error, event, status }) => {
//	console.log("Error Hook loaded \n\n")
//	const typedError = error as CustomHttpError;
//
//	console.log("Inside my handle error hook" + JSON.stringify(error, null, 2))
//	console.log(`status: ${status}, error.status: ${typedError.status}, flag: ${(typedError.status >= 300 && typedError.status <= 399)}`)
//
//	if (!(typedError.status >= 300 && typedError.status <= 399) && typedError.body && !isRedirect(error)) {
//		console.log("set error")
//		err = {
//			status: typedError.status,
//			msg: typedError.body.message,
//			shouldRedirect: typedError.shouldRedirect
//		}
//		event.locals.error = err;
//		console.log("First time loging err to locals: " + JSON.stringify(err, null, 2))
//	}
//
//	if (!isRedirect(error) && typedError.shouldRedirect && typedError.shouldRedirect.flag) {
//		//should = typedError.shouldRedirect;
//		console.log("first")
//		//throw new Redirect(301, typedError.shouldRedirect.path);
//		console.log(">>> About to load locals with: " + JSON.stringify(err, null, 2))
//		event.locals.error = err;
//		console.log(">>> Locals loaded with: " + JSON.stringify(event.locals.error, null, 2))
//		//redirect(307, typedError.shouldRedirect.path);
//	}
//	//if (error.status && (error.status >= 300 && error.status <= 399)) {
//	//	console.log("second")
//	//	redirect(301, error.location);
//	//}
//	//if (count == 1) {
//	//	console.log("redirecting...")
//	//	redirect(307, should.path);
//	//}
//}

export const handle: Handle = async ({ event, resolve }) => {
	console.log("Event: " + JSON.stringify(event, null, 2));

	if (event.route.id.startsWith('/dashboard')) {
		console.log("Inside route")
		const token = event.cookies.get("jwt") as string;
		try {
			const response = await fetch(apiUrl("/auth/me"), {
				method: 'GET',
				headers: {
					'Authorization': 'Bearer ' + token
				}
			});
			const data = await response.json()
			if (response.ok && data.user_type == 0) {
				console.log("data " + data);
				event.locals.user = {
					user_id: data.user_id,
					usrType: data.user_type
				}
				event.locals.token = token;
			} else {
				throw new Error("Unauthorized")
			}
		} catch (error) {
			console.error(error)
			event.locals.error = {
				status: 401,
				msg: "Unauthorized user. Please login",
				shouldRedirect: {
					flag: true,
					path: "/",
				}
			}
			console.log("Events right now: " + JSON.stringify(event.locals))
			console.log("redirecting...")
		}
	}
	//const errorMsgCookie = event.cookies.get('errorMsg');
	//if (errorMsgCookie && errorMsgCookie != "") {
	//	event.locals.errorMsg = errorMsgCookie;
	//	console.log(event.locals);
	//	event.cookies.delete('errorMsg', {
	//		path: "/",
	//	})
	//}
	const response = await resolve(event);
	return response;
}

