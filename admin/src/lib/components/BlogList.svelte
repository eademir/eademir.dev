<script>
    import { links, Link } from "svelte-navigator";
    import Loading from "./Loading.svelte";

    let page = 0;
    let limit = 10;
    let offset = limit * page;

    let url =
        "http://localhost:8000/api/v1/blogs?limit=" +
        limit +
        "&offset=" +
        offset;

    const getBlogs = async () =>
        await fetch(url, {
            method: "GET",
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                return data.blogs;
            })
            .catch((error) => {
                console.error("Error:", error);
            });

    let promise = getBlogs();
</script>

<div class="grid grid-cols-1 gap-4 text-black dark:text-white p-4" use:links>
    <div class="flow-root">
        <Link
            to="/blog/create-post"
            class="text-center w-[300px] float-right inline-block px-6 py-2.5 bg-green-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-green-600 hover:shadow-lg focus:bg-green-600 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-700 active:shadow-lg transition duration-150 ease-in-out"
        >
            Create A New Blog</Link
        >
    </div>
    {#await promise}
        <Loading />
    {:then data}
        {#each data as blog}
            <Link
                to="/blog/{blog.id}"
                class="flex flex-col items-center bg-white border rounded-lg shadow-md md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700"
            >
                <img
                    class="object-cover w-full rounded-t-lg h-96 md:h-[200px] md:w-[300px] md:rounded-none md:rounded-l-lg"
                    src={blog.image_url}
                    alt=""
                />
                <div class="flex flex-col justify-between p-4 leading-normal">
                    <h5
                        class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
                    >
                        {blog.title}
                    </h5>
                    <p
                        class="mb-3 font-normal text-gray-700 dark:text-gray-400"
                    >
                        {blog.description}
                    </p>
                </div>
            </Link>
        {/each}
    {/await}
</div>
