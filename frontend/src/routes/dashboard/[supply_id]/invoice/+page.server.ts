import type { PageServerLoad } from './$types';


export const load: PageServerLoad = ({ locals }) => {
	console.log("Locals in page " + JSON.stringify(locals, null, 2))
	return {
		localsD: locals,
	}
}
