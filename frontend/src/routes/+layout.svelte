<script lang="ts">
	import '../app.css';
	let { children } = $props();
	let username = $state('MARINA');
</script>

<div class="flex h-screen">
	<nav
		class="flex flex-col justify-between bg-gradient-to-br from-[#FF4000] from-10% via-[#1F0A70] via-[#C34623] via-30% to-[#28057B]"
	>
		<ul>
			<li>
				<h2 class="sm:text-md mx-auto my-5 text-center text-xl font-bold lg:text-2xl">
					Welcome {username}!
				</h2>
			</li>
		</ul>
		<ul>
			<li><h3 class="mb-4">Sign Out</h3></li>
		</ul>
	</nav>

	<div class="flex-1 overflow-auto">
		{@render children()}
	</div>
</div>

<style>
</style>
