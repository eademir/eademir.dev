<script>
	import { currentPath } from "../../../stores.js";
	import { Link } from "svelte-navigator";

	export let item = "";
	export let route;
</script>

<li>
	<Link
		to={route}
		class="block py-2 pl-3 pr-4 font-semibold font-serif text-slate-900 bg-blue-700 rounded md:bg-transparent md:p-0 dark:text-gray-200"
		aria-current="page"
		>{item}
		<!--<div class:active="{route === $currentPath}">
		
	</div>-->
	</Link>
</li>

<style>
</style>
