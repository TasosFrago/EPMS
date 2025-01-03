import { type Handle } from '@sveltejs/kit';


export const handle: Handle = async ({ event, resolve }) => {
	const errorMsgCookie = event.cookies.get('errorMsg');
	if (errorMsgCookie && errorMsgCookie != "") {
		event.locals.errorMsg = errorMsgCookie;
		console.log(event.locals);
		event.cookies.delete('errorMsg', {
			path: "/",
		})
	}
	const response = await resolve(event);
	return response;
}
