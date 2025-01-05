import type { LayoutServerLoad } from "./$types"


export const load: LayoutServerLoad = ({ locals }) => {
	console.log("\nLoading Layout/ ")
	console.log("Locals", JSON.stringify(locals, null, 2))
	console.log("user " + locals.user)
	if (locals.error) {
		console.log("Inside / " + locals.error.msg);
		//cookies.set('error_message', locals.error.msg, { path: '/', maxAge: 5 });
		//if (locals.error.shouldRedirect && locals.error.shouldRedirect.flag) {
		//	redirect(302, locals.error.shouldRedirect.path);
		//}
		return { error: locals.error }
	}
	//if (locals.errorMsg && locals.errorMsg != "") {
	//	console.log("Inside /" + locals.errorMsg)
	//	return { msg: locals.errorMsg }
	//}
	//return { msg: null }
}
