<script>
	import { Router, Route } from "svelte-navigator";
	import Nav from "./views/components/nav/Nav.svelte";
	import About from "./views/pages/About.svelte";
	import Admin from "./views/pages/BlogDetail.svelte";
	import Blog from "./views/pages/Blog.svelte";
	import Contact from "./views/pages/Contact.svelte";
	import Projects from "./views/pages/Projects.svelte";
	import { currentPath } from "./stores";
	import BlogDetail from "./views/pages/BlogDetail.svelte";
	import CommentPolicy from "./views/pages/CommentPolicy.svelte";
</script>

<Router primary={false}>
	<main class="gradient-form min-h-screen bg-gray-200 dark:bg-slate-900">
		<header class="h-[70px]">
			<Nav />
		</header>
		<Route path="blogs">
			<Blog />
		</Route>
		<Route path="projects">
			<Projects />
		</Route>
		<Route path="contact">
			<Contact />
		</Route>
		<Route path="/">
			<About />
		</Route>
		<Route path="blog/:id" let:params>
			<BlogDetail id={params.id} />
		</Route>
		<Route path="comment-policy">
			<CommentPolicy />
		</Route>
	</main>
</Router>

<style>
</style>
