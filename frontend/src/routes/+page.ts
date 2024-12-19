import type { PageLoad } from "./$types";


export const load: PageLoad = async ({ params }) => {
	const res = await fetch("http://localhost:8080/consumers");
	const items = await res.json();

	return { items };
};
