<script>
    import {
        isVisible,
        isPopUpVisable,
        showToast,
        message,
    } from "../../stores";
    import AddProject from "./AddProject.svelte";
    import Loading from "./Loading.svelte";
    import PopUp from "./PopUp.svelte";

    let visible = false;
    let popup = false;
    let id = undefined;

    const getProjects = async () =>
        await fetch("http://localhost:8000/api/v1/projects", {
            method: "GET",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                return data.projects;
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    let promise = getProjects();

    const deleteProject = async () => {
        const url = "http://localhost:8000/api/v1/projects/" + id;

        await fetch(url, {
            method: "DELETE",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                message.set(data.message);
                showToast.set(true);
                isPopUpVisable.set(false);
                handlePromise();
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    function handlePromise() {
        promise = getProjects();
    }

    isVisible.subscribe((value) => (visible = value));

    isPopUpVisable.subscribe((value) => (popup = value));
</script>

<div class="grid grid-cols-1 gap-4 text-black dark:text-white">
    {#if visible}
        <AddProject {id} on:submit={handlePromise} />
    {/if}
    {#if popup}
        <PopUp on:submit={deleteProject} />
    {/if}
    <div class="flow-root p-4">
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-missing-attribute -->
        <a
            on:click={() => {
                id = undefined;
                isVisible.set(true);
            }}
            class="text-center w-[300px] float-right inline-block px-6 py-2.5 bg-green-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-green-600 hover:shadow-lg focus:bg-green-600 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-700 active:shadow-lg transition duration-150 ease-in-out"
        >
            Create A New Project</a
        >
    </div>
    {#await promise}
        <Loading />
    {:then projects}
        {#each projects as project}
            <div
                class="flex flex-col p-4 bg-white border rounded-lg shadow-md md:flex-row dark:border-gray-700 dark:bg-gray-800"
            >
                <div
                    class="flex flex-col p-4 bg-white border rounded-lg shadow-md md:flex-col dark:border-gray-700 dark:bg-gray-800 w-full"
                >
                    <h5
                        class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
                    >
                        {project.title}
                        <span class="font-thin text-sm align-top"
                            >({new Date(
                                Date.parse(project.date)
                            ).getFullYear()})</span
                        >
                    </h5>
                    <p
                        class="mb-3 font-normal text-gray-700 dark:text-gray-400"
                    >
                        {project.description}
                    </p>
                    {#if project.links !== null}
                        <div class="flex flex-row gap-4">
                            {#each project.links as link}
                                <a href={link}>[{link}]</a>
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
                <div class="flex justify-center items-center p-4 h-fill">
                    <button
                        on:click={() => {
                            id = project.id;
                            isVisible.set(true);
                        }}
                        class="focus:outline-none text-white bg-yellow-400 hover:bg-yellow-500 focus:ring-4 focus:ring-yellow-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:focus:ring-yellow-900"
                    >
                        Update
                    </button>
                    <button
                        class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
                        on:click={() => {
                            id = project.id;
                            isPopUpVisable.set(true);
                        }}
                    >
                        Delete
                    </button>
                </div>
            </div>
        {/each}
    {/await}
</div>

<style>
    a {
        cursor: pointer;
    }
</style>
