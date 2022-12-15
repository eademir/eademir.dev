<script>
    import Loading from "../components/loading/Loading.svelte";
    import { onMount, onDestroy } from "svelte";
    import { currentPath } from "../../stores";

    import { links, Link } from "svelte-navigator";

    let page = 0;
    let limit = 10;
    let offset = limit * page;

    async function getBlogs() {
        let url =
            "http://localhost:8000/blogs?limit=" + limit + "&offset=" + offset;

        let res = await fetch(url, { method: "GET" });

        if (res.ok) {
            return res.json();
        } else {
            throw new Error();
        }
    }

    let promise = getBlogs();

    onMount(() => ($currentPath = window.location.pathname));
    onDestroy(() => ($currentPath = null));
</script>

<div class="grid grid-cols-1 gap-4 text-black dark:text-white p-4" use:links>
    {#await promise}
        <Loading />
    {:then data}
        {#each data.blogs as blog}
            <Link
                to="/blog/{blog.id}"
                class="flex flex-col m-auto items-center w-full 2xl:w-2/3 bg-white border rounded-lg shadow-md md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700"
            >
                <img
                    class="object-cover w-full rounded-t-lg h-96 md:h-[200px] md:w-[300px] md:rounded-none md:rounded-l-lg"
                    src={blog.image_url}
                    alt=""
                />
                <div class="flex flex-col justify-between p-4 leading-normal">
                    <h5
                        class="mb-2 text-2xl font-bold font-serif tracking-tight text-gray-900 dark:text-white"
                    >
                        {blog.title}
                    </h5>
                    <p
                        class="mb-3 font-normal text-gray-700 dark:text-gray-400 font-sans"
                    >
                        {blog.description}
                    </p>
                </div>
            </Link>
        {/each}
        <ul class="flex flex-row p-4 m-auto">
            {#each { length: data.totalBlogs / limit } as _, i}
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                    on:click={() => {
                        page = i;
                        promise = getBlogs();
                    }}
                >
                    <li>
                        {i + 1}
                    </li></a
                >
            {/each}
        </ul>
    {:catch}
        {Error}
    {/await}
</div>

<style>
    a {
        cursor: pointer;
    }
    li {
        border: solid 1px #ccc;
        padding: 5px 8px;
        margin: 4px;
        font-family: Georgia, "Times New Roman", Times, serif;
    }
</style>
