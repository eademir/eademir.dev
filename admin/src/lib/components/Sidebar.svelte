<script>
    import { navigate } from "svelte-navigator";
    ("use strict");
    import { activePage } from "../../stores";

    const setIndex = (index) => {
        activePage.set(index);
    };

    let logout = async () =>
        await fetch("/api/v1/logout", {
            method: "POST",
            credentials: "include",
            redirect: "follow",
        })
            .then((res) => {
                if (res.ok) {
                    navigate("/", { replace: true });
                }
            })
            .catch((error) => {
                console.error("Error:", error);
            });
</script>

<aside class="w-full" aria-label="Sidebar">
    <div class="overflow-y-auto py-4 px-3 bg-gray-50 rounded dark:bg-gray-800">
        <ul class="space-y-2">
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(0)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path d="M21.21 15.89A10 10 0 1 1 8 2.83" /><path
                            d="M22 12A10 10 0 0 0 12 2v10z"
                        /></svg
                    >
                    <span class="ml-3">Dashboard</span>
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(1)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path
                            d="M14 2H6a2 2 0 0 0-2 2v16c0 1.1.9 2 2 2h12a2 2 0 0 0 2-2V8l-6-6z"
                        /><path d="M14 3v5h5M16 13H8M16 17H8M10 9H8" /></svg
                    >
                    <span class="flex-1 ml-3 whitespace-nowrap">Blogs</span>
                    <span
                        class="inline-flex justify-center items-center px-2 ml-3 text-sm font-medium text-gray-800 bg-gray-200 rounded-full dark:bg-gray-700 dark:text-gray-300"
                    />
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(2)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path
                            d="M3 15v4c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2v-4M17 9l-5 5-5-5M12 12.8V2.5"
                        /></svg
                    >
                    <span class="flex-1 ml-3 whitespace-nowrap">Inbox</span>
                    <span
                        class="inline-flex justify-center items-center p-3 ml-3 w-3 h-3 text-sm font-medium text-blue-600 bg-blue-200 rounded-full dark:bg-blue-900 dark:text-blue-200"
                    />
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(3)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path
                            d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"
                        /><circle cx="9" cy="7" r="4" /><path
                            d="M23 21v-2a4 4 0 0 0-3-3.87"
                        /><path d="M16 3.13a4 4 0 0 1 0 7.75" /></svg
                    >

                    <span class="flex-1 ml-3 whitespace-nowrap">User</span>
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(4)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><rect
                            x="9"
                            y="9"
                            width="13"
                            height="13"
                            rx="2"
                            ry="2"
                        /><path
                            d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"
                        /></svg
                    >
                    <span class="flex-1 ml-3 whitespace-nowrap">Projects</span>
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => setIndex(5)}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path
                            d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"
                        /><circle cx="12" cy="7" r="4" /></svg
                    >
                    <span class="flex-1 ml-3 whitespace-nowrap">About Me</span>
                </a>
            </li>
            <li>
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => logout()}
                    on:keypress
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                    <svg
                        class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        ><path d="M18.36 6.64a9 9 0 1 1-12.73 0" /><line
                            x1="12"
                            y1="2"
                            x2="12"
                            y2="12"
                        /></svg
                    >
                    <span class="flex-1 ml-3 whitespace-nowrap">Sign Out</span>
                </a>
            </li>
        </ul>
    </div>
</aside>

<style>
    a {
        cursor: pointer;
    }
</style>
