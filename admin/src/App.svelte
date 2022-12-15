<script>
  import { Router, Route } from "svelte-navigator";
  import Editor from "./lib/pages/Editor.svelte";
  import LoginPage from "./lib/pages/LoginPage.svelte";
  import Panel from "./lib/pages/Panel.svelte";
</script>

<main>
  <Router primary={false}>
    <Route path="/">
      <LoginPage />
    </Route>
    <Route path="panel">
      <Panel />
    </Route>
    <Route path="blog/:id" let:params>
      <Editor id={params.id} />
    </Route>
    <Route path="blog/create-post">
      <Editor id={undefined} />
    </Route>
  </Router>
</main>

<style>
</style>
