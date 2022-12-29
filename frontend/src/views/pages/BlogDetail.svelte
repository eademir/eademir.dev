<script>
    import SvelteMarkdown from "svelte-markdown";
    import { marked } from "marked";
    import hljs from "highlight.js";
    import "highlight.js/styles/stackoverflow-dark.css";
    import Loading from "../components/loading/Loading.svelte";

    export let id;
    const monthNames = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    ];

    const url = "/api/v1/blogs/" + id;

    async function getBlog() {
        let res = await fetch(url);

        if (res.ok) {
            return res.json();
        } else {
            throw new Error();
        }
    }

    let promise = getBlog();

    //marked options
    marked.setOptions({
        renderer: new marked.Renderer(),
        highlight: (code, lang) => {
            //detect whether language is contained by the library
            //it crashes immediately if tries to highlight something like javascrip or ja or g, etc.
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
</script>

<svelte:head>
    {#await promise then blog}
        <title>{blog.title} - EADEMIR</title>
        <meta name="description" content={blog.description} />
        <meta name="keywords" content={blog.keywords} />
        <meta name="author" content={blog.username} />

        <script>
            var disqus_config = function () {
                this.page.url = window.location.href; // Replace PAGE_URL with your page's canonical URL variable
                this.page.identifier = blog.id; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
            };

            (function () {
                // DON'T EDIT BELOW THIS LINE
                var d = document,
                    s = d.createElement("script");
                s.src = "https://eademir-dev.disqus.com/embed.js";
                // @ts-ignore
                s.setAttribute("data-timestamp", +new Date());
                (d.head || d.body).appendChild(s);
            })();
        </script>
        <noscript
            >Please enable JavaScript to view the <a
                href="https://disqus.com/?ref_noscript"
                >comments powered by Disqus.</a
            ></noscript
        >
    {/await}
</svelte:head>

<div
    class="flex flex-col m-auto w-full xl:w-2/3 2xl:w-1/2 p-4 gap-4 text-slate-900 dark:text-gray-200"
>
    {#await promise}
        <Loading />
    {:then data}
        <div class="relative">
            <h1 class="md:text-center w-full font-bold font-sans">
                {data.title}
            </h1>
        </div>
        <div class="relative overflow-hidden w-full min-h-[300px]">
            <div
                class="grid grid-row absolute top-0 right-0 z-10 text-gray-100 bg-slate-700 p-2 rounded-l"
            >
                <small
                    >Created:
                    {new Date(Date.parse(data.createdAt)).getDate()}
                    {monthNames[
                        new Date(Date.parse(data.createdAt)).getMonth()
                    ]}
                    {new Date(Date.parse(data.createdAt)).getFullYear()}</small
                >
                {#if data.updatedAt != undefined}
                    <small
                        >Edited:
                        {new Date(Date.parse(data.updatedAt)).getDate()}
                        {monthNames[
                            new Date(Date.parse(data.updatedAt)).getMonth()
                        ]}
                        {new Date(
                            Date.parse(data.updatedAt)
                        ).getFullYear()}</small
                    >
                {/if}
            </div>
            <img
                class="absolute m-auto top-[-9999px] left-[-9999px] bottom-[-9999px] right-[-9999px]"
                src={data.imageURL}
                alt={data.title}
                width="100%"
            />
        </div>
        <div class="grid gap-6 font-sans">
            <SvelteMarkdown source={marked(data.content)} />
        </div>
    {/await}
    <div id="disqus_thread" class="md:p-10 mt-10" />
</div>

<style>
    :global(h2) {
        font-size: 25px;
        font-weight: bold;
        font-family: Georgia, "Times New Roman", Times, serif;
    }
    @media (prefers-color-scheme: light) {
        :global(pre) {
            background-color: #ddd;
            padding: 20px;
            max-width: 100% !important;
            overflow-x: auto;
        }
        :global(code) {
            background-color: #ddd;
        }
    }
    @media (prefers-color-scheme: dark) {
        :global(pre) {
            padding: 20px;
            background-color: #333;
            max-width: 100% !important;
            overflow-x: auto;
        }
        :global(code) {
            background-color: #333;
        }
    }
</style>
