<script module lang="ts">
	export interface LogInData {
		email: string;
		password: string;
	}
</script>

<script lang="ts">
	interface Props {
		title: string;
		data: LogInData;
	}

	let { title, data }: Props = $props();
</script>
