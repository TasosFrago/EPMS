import type { LayoutServerLoad } from "./$types"


export const load: LayoutServerLoad = ({ locals }): { msg: string | null } => {
	if (locals.errorMsg && locals.errorMsg != "") {
		console.log("Inside /" + locals.errorMsg)
		return { msg: locals.errorMsg }
	}
	return { msg: null }
}
