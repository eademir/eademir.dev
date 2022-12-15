<script>
    import { onMount } from "svelte";
    import Loading from "./Loading.svelte";

    $: username = "";
    $: email = "";
    $: new_password = "";
    $: password = "";
    $: message = "";
    $: id = 0;

    async function getUser() {
        const res = await fetch("http://localhost:8000/api/v1/user", {
            method: "GET",
            credentials: "include",
        });
        const json = await res.json();
        if (res.ok) {
            username = json.username;
            email = json.email;
            id = json.id;
        } else {
            throw new Error(json);
        }
    }

    async function updateUser() {
        const res = await fetch("http://localhost:8000/api/v1/user", {
            method: "PUT",
            credentials: "include",
            body: JSON.stringify({
                id,
                username,
                email,
                password,
                new_password,
            }),
        });
        const json = await res.json();

        if (res.ok) {
            message = json.message;
        } else {
            message = json.message;
            throw new Error(json);
        }
    }

    let promise = getUser();

    onMount(() => {
        getUser();
    });
</script>

<div class="text-black dark:text-white p-4">
    {#await promise}
        <Loading/>
    {:then a}
        <div class="max-w-xl m-auto">
            <div class="mb-6">
                <label
                    for="email"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >Email</label
                >
                <input
                    type="email"
                    id="email"
                    name="email"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Email"
                    bind:value={email}
                />
            </div>
            <div class="mb-6">
                <label
                    for="username"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >Username</label
                >
                <input
                    type="username"
                    id="username"
                    name="username"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Username"
                    bind:value={username}
                    disabled
                />
            </div>
            <div class="mb-6">
                <label
                    for="password"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >New Password</label
                >
                <input
                    type="password"
                    id="password"
                    name="password"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    bind:value={new_password}
                />
            </div>

            <div class="mb-6">
                <label
                    for="old_password"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >Old Password</label
                >
                <input
                    type="password"
                    id="old_password"
                    name="old_password"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    bind:value={password}
                    required
                />
            </div>
            <button
                type="submit"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                on:click={() => updateUser()}>Submit</button
            >
        </div>
    {/await}
</div>
