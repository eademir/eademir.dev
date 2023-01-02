<script>
	import NavItem from "./NavItem.svelte";
	import Logo from "../../../assets/logo.svg";
	import { fly } from "svelte/transition";

	let x = 0;
	$: isNavVisible = x > 768 ? true : false;

	function closeNav() {
		if (x <= 768) {
			isNavVisible = false;
		}
	}
</script>

<svelte:window bind:innerWidth={x} />

<nav
	class="bg-white px-2 dark:bg-gray-900 fixed w-full z-20 top-0 left-0 border-b border-gray-200 dark:border-gray-600"
>
	<div
		class="container flex flex-wrap items-center justify-between mx-auto relative sm:px-4 py-2.5"
	>
		<a href="/" class="flex items-center z-10 dark:bg-gray-900 bg-white">
			<img src={Logo} class="h-6 mr-3 sm:h-9" alt="eademir Logo" />
			<span
				class="self-center font-serif text-xl font-semibold whitespace-nowrap dark:text-white"
				>EADEMIR</span
			>
		</a>
		<div class="flex md:order-2 z-10">
			<button
				type="button"
				class="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
				on:click={() => (isNavVisible = !isNavVisible)}
			>
				{#if isNavVisible}
					<span class="sr-only">Close main menu</span>
					<svg
						class="w-6 h-6"
						aria-hidden="true"
						xmlns="http://www.w3.org/2000/svg"
						fill="currentColor"
						viewBox="0 0 16 16"
					>
						<path
							d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
						/>
					</svg>
				{:else}
					<span class="sr-only">Open main menu</span>
					<svg
						class="w-6 h-6"
						aria-hidden="true"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
							clip-rule="evenodd"
						/></svg
					>
				{/if}
			</button>
		</div>
		{#if isNavVisible}
			<div
				class="items-center justify-between w-full md:flex md:w-auto md:order-1 z-0 absolute md:relative top-14 md:top-0"
				out:fly={{ y: -300 }}
				in:fly={{ y: -300 }}
			>
				<ul
					class="flex flex-col p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
					on:click={() => closeNav()}
				>
					<NavItem route="/blogs" item="Blogs" />
					<NavItem route="/" item="About Me" />
					<NavItem route="/projects" item="Projects" />
					<NavItem route="/contact" item="Contact" />
				</ul>
			</div>
		{/if}
	</div>
</nav>

<style>
</style>
