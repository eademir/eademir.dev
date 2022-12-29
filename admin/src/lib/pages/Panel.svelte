<script>
    "user strict";
    import { activePage, isPopUpVisable, showToast } from "../../stores";
    import Dashboard from "../components/Dashboard.svelte";
    import BlogList from "../components/BlogList.svelte";
    import Sidebar from "../components/Sidebar.svelte";
    import Users from "../components/Users.svelte";
    import Projects from "../components/Projects.svelte";
    import AboutMe from "../components/AboutMe.svelte";
    import PopUp from "../components/PopUp.svelte";
    import Toast from "../components/Toast.svelte";

    let index = 0;
    let popup = false;
    let toast = false;

    activePage.subscribe((value) => (index = value));
    isPopUpVisable.subscribe((value) => (popup = value));
    showToast.subscribe((value) => (toast = value));
</script>

<div class="gradient-form min-h-screen bg-gray-200 dark:bg-slate-900">
    {#if toast}
        <Toast />
    {/if}
    {#if popup}
        <PopUp />
    {/if}
    <div class="grid grid-flow-row-dense xl:grid-cols-6">
        <div class="p-4 lg:col-span-1">
            <Sidebar />
        </div>
        <div class="xl:col-span-5">
            {#if index === 0}
                <Dashboard />
            {:else if index === 1}
                <BlogList />
            {:else if index === 4}
                <Projects />
            {:else if index === 3}
                <Users />
            {:else if index === 5}
                <AboutMe />
            {/if}
        </div>
    </div>
</div>
