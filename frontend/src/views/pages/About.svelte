<script>
    import { onMount, onDestroy } from "svelte";
    import { currentPath } from "../../stores";
    import Loading from "../components/loading/Loading.svelte";

    onMount(() => ($currentPath = window.location.pathname));
    const getAbout = async () => {
        var response = await fetch("http://localhost:8000/");
        var result = await response.json();
        return result.about;
    };

    let promise = getAbout();
    onDestroy(() => ($currentPath = null));
</script>

<svelte:head>
        {#await promise then about}
        <title>About - eademir.dev</title>
        <meta name="description" content={about.detail} />
        <meta name="keywords" content="eray, software, developer, flutter, golang, javascript, svelte, gin gonic" />
        <meta name="author" content="Eray" />
        {/await}
</svelte:head>

<div
    class="flex flex-col m-auto items-center w-full 2xl:w-2/3 md:flex-row p-4 gap-4"
>
    {#await promise}
        <Loading />
    {:then about}
        <div>
            {#each about.detail.split("\n") as text}
                <p
                    class="text-slate-900 dark:text-gray-200 font-normal font-sans"
                >
                    {text}
                </p>
                <br />
            {/each}
        </div>
        {#if about.image_url !== "no"}
            <img src={about.image_url} alt="eray aydemir" width="300px" />
        {/if}
    {:catch err}
        <h1>{err}</h1>
    {/await}
</div>

<style>
</style>
