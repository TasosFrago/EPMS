<script lang="ts">
	import { writable, type Writable } from 'svelte/store';
	import Login, { type LogInData } from '$lib/components/Login.svelte';
	import { apiUrl, debugLog } from '$lib/settings';

	let data: Writable<LogInData> = writable({
		email: '',
		password: ''
	});

	const handleLogin = async (): Promise<void> => {
		const loginData = $data;
		debugLog(loginData);
	};
</script>

<Login title={'Login Consumer'} {data} handleSubmit={handleLogin} />
