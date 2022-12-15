<script>
    "use strict";
    import { showToast, message } from "../../stores";
    import { fade, fly } from "svelte/transition";
    import Logo from "../assets/logo.jpeg";

    let m = "";

    setTimeout(() => showToast.set(false), 3000);

    message.subscribe((value) => (m = value));
</script>

<div
    id="toast-default"
    class="flex absolute top-5 right-5 items-center p-4 space-x-4 w-full max-w-xs text-gray-500 bg-white rounded-lg divide-x divide-gray-200 shadow dark:text-gray-400 dark:divide-gray-700 space-x dark:bg-gray-800"
    role="alert"
    out:fade
    in:fly={{ y: -50 }}
>
    <div
        class="inline-flex flex-shrink-0 justify-center items-center w-8 h-8 text-blue-500 bg-blue-100 rounded-lg dark:bg-blue-800 dark:text-blue-200"
    >
        <img class="mx-auto w-8" src={Logo} alt="logo" />
        <span class="sr-only">Logo</span>
    </div>
    <div class="ml-3 text-sm font-normal w-full capitalize">{m}</div>
    <button
        type="button"
        class="ml-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700"
        aria-label="Close"
        on:click={() => showToast.set(false)}
    >
        <span class="sr-only">Close</span>
        <svg
            aria-hidden="true"
            class="w-5 h-5"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
            ><path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
            /></svg
        >
    </button>
</div>
