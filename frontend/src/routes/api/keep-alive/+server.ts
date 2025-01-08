import { apiUrl } from '$lib/settings';
import { json } from '@sveltejs/kit';

export async function GET() {
	const randomChoise = (Math.floor(Math.random() * 2) + 1) == 1 ? 'one' : 'two';
	try {
		console.log("Random ping on: " + randomChoise);
		const response = await fetch(apiUrl(`/keep-alive/${randomChoise}`));

		if (!response.ok) {
			throw new Error(`Error pinging API: ${response.statusText}`);
		}

		return json({ message: 'Pinged API successfully' });
	} catch (error) {
		return json({ error: error.message }, { status: 500 });
	}
}
