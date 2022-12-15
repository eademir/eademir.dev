<script>
    //import { message } from '../../stores';

    import Loading from "./Loading.svelte";

    const url = "http://localhost:8000/api/v1/about";

    $: detail = "";
    $: image_url = "";
    let username;
    let message;

    const updateAbout = async () => {
        await fetch(url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                username,
                detail,
                image_url,
            }),
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                message = data.message;
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    const getAbout = async () => {
        await fetch(url, {
            method: "GET",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                image_url = data.about.image_url;
                detail = data.about.detail;
                username = data.about.username;
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    let promise = getAbout();
</script>

<div class="text-black dark:text-white p-4">
    {#await promise}
        <Loading />
    {:then a}
        <div class="max-w-xl m-auto grid gap-4">
            <div>
                <label
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    for="aboutme">About me</label
                >
                <textarea
                    name="aboutme"
                    id="aboutme"
                    class="bg-gray-50 border min-h-[400px] border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    bind:value={detail}
                />
            </div>
            <div>
                <label
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    for="image">Image</label
                >
                <textarea
                    name="image"
                    id="image"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    bind:value={image_url}
                />
            </div>
            <button
                class="mt-4 w-full focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
                on:click={() => updateAbout()}
            >
                Publish
            </button>
        </div>
    {/await}
</div>

<style>
    textarea {
        resize: none;
    }
</style>
