<script>
    import { isVisible } from "../../stores";
    import { createEventDispatcher } from "svelte";
    import { fade } from "svelte/transition";
    import Loading from "./Loading.svelte";

    const dispatch = createEventDispatcher();
    const submit = () => dispatch("submit");

    export let id;
    const url = "http://localhost:8000/api/v1/projects/" + id;

    let project;
    let message;

    $: title = "";
    $: description = "";
    $: date = new Date();
    $: link = "";
    $: image = "";
    let username;

    const updateProject = async () => {
        const images =
            image !== "" && image !== undefined
                ? image.split(",").map((e) => e.trim())
                : [];

        const links =
            link !== "" && image !== undefined
                ? link.split(",").map((e) => e.trim())
                : [];

        await fetch(url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                username,
                title,
                description,
                date,
                images,
                links,
            }),
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                message = data.message;
                isVisible.set(false);
                getProject();
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    const addProject = async () => {
        const url = "http://localhost:8000/api/v1/projects";
        const images =
            image !== "" && image !== undefined
                ? image.split(",").map((e) => e.trim())
                : [];

        const links =
            link !== "" && image !== undefined
                ? link.split(",").map((e) => e.trim())
                : [];
        await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                title,
                description,
                date,
                images,
                links,
            }),
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                message = data.message;
                isVisible.set(false);
                getProject();
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    const getProject = async () => {
        if (id !== undefined) {
            await fetch(url, {
                method: "GET",
                credentials: "include",
            })
                .then((response) => response.json())
                .then((data) => (project = data.project))
                .then(() => {
                    title = project.title;
                    description = project.description;
                    date = project.date;
                    if (project.links !== null) link = project.links.toString();
                    if (project.images !== null)
                        image = project.images.toString();
                    username = project.username;
                })
                .catch((error) => {
                    console.error("Error:", error);
                });
        }
    };

    let promise = getProject();
</script>

<div
    class="grid h-screen w-full fixed top-0 right-0 left-0 items-center z-10 bg-[rgba(0,0,0,0.5)]"
    out:fade={{ duration: 300 }}
    in:fade={{ duration: 300 }}
>
    <div class="place-items-center">
        <div
            class="relative ml-auto mr-auto text-black dark:text-white border rounded-lg shadow-md dark:border-gray-500 border-gray-700 max-w-[600px] p-4 bg-gray-200 dark:bg-slate-900"
        >
            <!-- svelte-ignore a11y-missing-attribute -->
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <a
                on:click={() => isVisible.set(false)}
                class="absolute right-4 top-2 text-red-500 font-bold text-xl close"
            >
                x
            </a>
            {#await promise}
                <Loading />
            {:then project}
                <form class="max-w-xl m-auto">
                    <div class="mb-6">
                        <label
                            for="title"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >Title</label
                        >
                        <input
                            type="text"
                            id="title"
                            name="title"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            placeholder="Title"
                            bind:value={title}
                        />
                    </div>
                    <div class="mb-6">
                        <label
                            for="date"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >Date</label
                        >
                        <input
                            type="date"
                            id="date"
                            name="date"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            placeholder="Date"
                            bind:value={date}
                        />
                    </div>
                    <div class="mb-6">
                        <label
                            for="description"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >Description</label
                        >
                        <textarea
                            type="text"
                            id="description"
                            name="description"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            bind:value={description}
                        />
                    </div>

                    <div class="mb-6">
                        <label
                            for="links"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >Links</label
                        >
                        <textarea
                            type="text"
                            id="links"
                            name="links"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            bind:value={link}
                        />
                    </div>

                    <div class="mb-6">
                        <label
                            for="images"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >Images</label
                        >
                        <textarea
                            type="text"
                            id="images"
                            name="images"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            bind:value={image}
                        />
                    </div>
                    <button
                        type="button"
                        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                        on:click={() => {
                            id !== undefined ? updateProject() : addProject();
                            submit();
                        }}>Submit</button
                    >
                </form>
            {/await}
        </div>
    </div>
</div>

<style>
    textarea {
        resize: none;
    }

    a {
        cursor: pointer;
    }
</style>
