<script>
    import Loading from "../components/loading/Loading.svelte";
    import { onMount, onDestroy } from "svelte";
    import { currentPath } from "../../stores";

    const getProjects = async () =>
        await fetch("/api/v1/projects")
            .then((response) => response.json())
            .then((data) => {
                return data.projects;
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    let promise = getProjects();

    onMount(() => ($currentPath = window.location.pathname));
    onDestroy(() => ($currentPath = null));
</script>

<svelte:head>
    {#await promise then project}
        <title>Blogs - eademir.dev</title>
        <meta name="description" content={project[0].description} />
        <meta
            name="keywords"
            content="eray, software, developer, flutter, golang, javascript, svelte, gin gonic, projects"
        />
        <meta name="author" content="Eray" />
    {/await}
</svelte:head>

<div class="grid grid-cols-1 gap-4 text-black dark:text-white p-4">
    {#await promise}
        <Loading />
    {:then projects}
        {#each projects as project}
            <div
                class="flex flex-col p-4 m-auto w-full 2xl:w-2/3 bg-white border rounded-lg shadow-md dark:border-gray-700 dark:bg-gray-800"
            >
                <h5
                    class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white font-sans"
                >
                    {project.title}
                    <span class="font-thin text-sm align-top font-serif"
                        >({new Date(
                            Date.parse(project.date)
                        ).getFullYear()})</span
                    >
                </h5>
                {#each project.description.split("\n") as text}
                    <p
                        class="mb-3 font-normal text-gray-700 dark:text-gray-400 font-sans"
                    >
                        {text}
                    </p>
                {/each}
                {#if project.links !== null}
                    <div class="flex flex-col gap-4 mt-4">
                        {#each project.links as link}
                            <a class="font-sans italic underline" href={link}
                                >{link}</a
                            >
                        {/each}
                    </div>
                {/if}

                {#if project.images !== null}
                    <div class="flex flex-row gap-4 mt-4">
                        {#each project.images as image}
                            <img
                                class="inline-block w-36"
                                src={image}
                                alt={image}
                            />
                        {/each}
                    </div>
                {/if}
            </div>
        {/each}
    {/await}
</div>

<style>
</style>
