<script>
    "use strict";
    import { onMount } from "svelte";
    import { marked } from "marked";
    import hljs from "highlight.js";
    import "highlight.js/styles/stackoverflow-dark.css";
    import { isPopUpVisable } from "../../stores";
    import PopUp from "../components/PopUp.svelte";

    import { useNavigate } from "svelte-navigator";
    import image from "../assets/logo.jpeg";
    const navigate = useNavigate();

    export let id;

    $: content = "";
    $: title = "";
    $: description = "";
    $: image_url = image;
    $: keywords = "";
    let editor;
    let reader;
    let ignoreScroll = false;
    let popup = false;

    isPopUpVisable.subscribe((value) => (popup = value));

    //syncronised scroller
    const scroller = (
        /** @type {{ scrollTop: number; scrollHeight: number; clientHeight: number; }} */ element1,
        /** @type {{ scrollTop: number; scrollHeight: number; }} */ element2
    ) => {
        let ignore = ignoreScroll;
        ignoreScroll = false;
        if (ignore) return;

        ignoreScroll = true;
        const multiplier =
            element1.scrollTop /
            (element1.scrollHeight - element1.clientHeight);
        element2.scrollTop =
            multiplier * (element2.scrollHeight - element1.clientHeight);
    };

    const url = "/api/v1/blogs/" + id;

    const getBlog = async () =>
        await fetch(url, {
            method: "GET",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                content = data["content"];
                title = data["title"];
                description = data["description"];
                keywords = data["keywords"];
                image_url = data["imageURL"];
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    const deleteBlog = async () => {
        await fetch(url, {
            method: "DELETE",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                isPopUpVisable.set(false);
                if (data.message == "success") navigate(-1);
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    const updateBlog = async (is_shared) =>
        await fetch(url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                title,
                description,
                content,
                keywords,
                is_shared,
                image_url,
            }),
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.message == "success") navigate(-1);
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    const createBlog = async (is_shared) =>
        await fetch("/api/v1/blogs", {
            method: "POST",
            body: JSON.stringify({
                title,
                description,
                content,
                keywords,
                is_shared,
                image_url,
            }),
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.message == "success") navigate(-1);
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    async function check() {
        await fetch("/api/v1/dashboard", {
            method: "GET",
            credentials: "include",
        })
            .then((response) => {
                if (!response.ok) {
                    console.log(response.status);
                    navigate("/", { replace: true });
                }
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    }

    onMount(() => {
        if (id != undefined) {
            getBlog();
        }

        check();
        //syncronised scroll for editor and reader parts
        editor.addEventListener("scroll", (/** @type {any} */ _e) => {
            scroller(editor, reader);
        });
        reader.addEventListener("scroll", (/** @type {any} */ _e) => {
            scroller(reader, editor);
        });
    });

    //marked options
    marked.setOptions({
        renderer: new marked.Renderer(),
        highlight: (/** @type {string} */ code, /** @type {string} */ lang) => {
            //detect whether language is contained by the library
            //it crashes immediately if tries to highlight something like javascript or ja or g, etc.
            const language = hljs.getLanguage(lang) ? lang : "plaintext";
            return hljs.highlight(code, { language }).value;
        },
        pedantic: false,
        gfm: true,
        breaks: true, //add '\n' to value automaticly when hit the Enter button
        smartLists: true,
        smartypants: false,
        xhtml: false,
    });

    //shortcut handler
    window.addEventListener("keydown", (event) => {
        if (event.repeat) return;
        if (event.key === "Tab") {
            handleKeydown(event, "\t", 1);
        } else if (
            (event.key === "1" || event.key === "2" || event.key === "3") &&
            event.ctrlKey
        ) {
            handleKeydown(event, "#".repeat(parseInt(event.key)) + " ");
        } else if (event.key === "-" && event.ctrlKey) {
            handleKeydown(event, "---");
        } else if (event.key === "i" && event.ctrlKey) {
            handleKeydown(event, "![alt]()", 7);
        } else if (event.key === "c" && event.ctrlKey) {
            handleKeydown(event, "```\n```", 3);
        } else if (event.key === '"') {
            handleKeydown(event, '""', 1);
        } else if (event.key === "(") {
            handleKeydown(event, "()", 1);
        } else if (event.key === "{") {
            handleKeydown(event, "{}", 1);
        } else if (event.key === "[") {
            handleKeydown(event, "[]", 1);
        }
    });

    const handleKeydown = (
        /** @type {KeyboardEvent | any} */ event,
        /** @type {string | any[]} */ code,
        /** @type {number} */ length
    ) => {
        const codeLength = typeof length !== "undefined" ? length : code.length;
        event.preventDefault();
        const start = event.target.selectionStart;
        const end = event.target.selectionEnd;
        event.target.value =
            event.target.value.substring(0, start) +
            code +
            event.target.value.substring(end);
        event.target.selectionStart = event.target.selectionEnd =
            start + codeLength;
    };
</script>

<div
    class="bg-white dark:bg-slate-900 text-black dark:text-white grid grid-flow-row-dense xl:grid-cols-6 min-h-screen"
>
    {#if popup}
        <PopUp on:submit={deleteBlog} />
    {/if}
    <div class="xl:col-span-4 flex flex-col">
        <div class="grid p-4 gap-4">
            <div class="relative z-0 w-full group">
                <label
                    for="title"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >Title</label
                >
                <textarea
                    class="w-full p-4 markdown-body"
                    name="title"
                    placeholder="Title"
                    rows="1"
                    maxlength="150"
                    bind:value={title}
                    required
                />
            </div>

            <div class="relative z-0 w-full group">
                <label
                    for="description"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >Description</label
                >
                <textarea
                    class="w-full p-4 markdown-body"
                    name="description"
                    placeholder="Description"
                    rows="2"
                    maxlength="280"
                    bind:value={description}
                    required
                />
            </div>
        </div>
        <div class="ml-4">
            <label
                for="content"
                class="block text-sm font-medium text-gray-900 dark:text-white"
                >Content</label
            >
        </div>
        <div class="grid xl:grid-cols-2 gap-4 m-4 flex-grow xl:max-h-[700px]">
            <textarea
                class="markdown-body p-4 overflow-x-auto min-h-[500px]"
                name="content"
                bind:value={content}
                bind:this={editor}
                required
            />
            <div class="markdown-body p-4 overflow-x-auto" bind:this={reader}>
                {@html marked(content)}
            </div>
        </div>
    </div>
    <div class="p-4 space-y-4 w-full xl:col-span-2">
        <img src={image_url} alt="blog" />
        <div class="relative z-0 w-full group">
            <label
                for="image_url"
                class="block mb-2  text-sm font-medium text-gray-900 dark:text-white"
                >Image URL</label
            >
            <textarea
                class="w-full markdown-body p-4"
                name="image_url"
                rows="5"
                placeholder="Image URL"
                bind:value={image_url}
                required
            />
        </div>
        <div class="relative z-0 w-full group">
            <label
                for="keywords"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Keywords</label
            >
            <textarea
                class="w-full markdown-body p-4"
                rows="5"
                name="keywords"
                placeholder="Keywords"
                bind:value={keywords}
                required
            />
        </div>
        <div class="grid grid-flow-col">
            <button
                class="focus:outline-none text-white bg-yellow-400 hover:bg-yellow-500 focus:ring-4 focus:ring-yellow-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:focus:ring-yellow-900"
                on:click={() =>
                    id != undefined ? updateBlog(false) : createBlog(false)}
            >
                Save
            </button>
            <button
                class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
                on:click={() =>
                    id != undefined ? updateBlog(true) : createBlog(true)}
            >
                Publish
            </button>
            {#if id != undefined}
                <button
                    class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
                    on:click={() => isPopUpVisable.set(true)}
                >
                    Delete
                </button>
            {/if}
        </div>
    </div>
</div>

<style>
    textarea:focus {
        outline: none !important;
    }

    .markdown-body {
        border: 2px solid teal;
    }

    @media (prefers-color-scheme: light) {
        @import "highlight.js/styles/stackoverflow-dark.css";
    }

    @media (prefers-color-scheme: dark) {
        @import "highlight.js/styles/stackoverflow.css";
    }

    .markdown-body::-webkit-scrollbar {
        width: 0.4em;
    }

    .markdown-body::-webkit-scrollbar-track,
    textarea {
        box-shadow: inset 0 0 6px rgba(103, 109, 109, 0.7);
        resize: none;
    }

    .markdown-body::-webkit-scrollbar-thumb {
        background-color: teal;
        outline: 1px solid teal;
    }
</style>
