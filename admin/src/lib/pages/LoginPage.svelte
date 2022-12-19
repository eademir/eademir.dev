<script>
    "user strict";
    import Logo from "../assets/logo.jpeg";
    import { result, statusCode } from "../../stores";
    import { navigate } from "svelte-navigator";

    var username = "";
    var password = "";
    var errorMessage;

    result.subscribe((value) => {
        errorMessage = value;
    });

    async function loginRequest() {
        await fetch("http://localhost:8000/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            credentials: "include",
            body: JSON.stringify({
                username,
                password,
            }),
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
                result.set(data.message);
                if (errorMessage === "success") {
                    navigate("/panel", { replace: true });
                }
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    }
</script>

<div>
    <section
        class="h-full gradient-form min-h-screen bg-gray-200 dark:bg-slate-900"
    >
        <div class="container py-12 px-6 h-full">
            <div
                class="flex justify-center items-center flex-wrap h-full g-6 text-gray-800"
            >
                <div class="xl:w-10/12">
                    <div
                        class="block bg-gray-50  dark:bg-black shadow-lg rounded-lg dark:text-white text-black"
                    >
                        <div class="lg:flex lg:flex-wrap g-0">
                            <div class="lg:w-6/12 px-4 md:px-0">
                                <div class="md:p-12 md:mx-6">
                                    <div class="text-center">
                                        <img
                                            class="mx-auto w-48"
                                            src={Logo}
                                            alt="logo"
                                        />
                                        <h4
                                            class="text-xl font-semibold mt-1 mb-12 pb-1"
                                        >
                                            Admin Panel Login Page
                                        </h4>
                                    </div>
                                    <form>
                                        <p class="mb-4">
                                            Please login to your account
                                        </p>
                                        <div class="mb-4">
                                            <input
                                                type="text"
                                                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                                placeholder="Username"
                                                bind:value={username}
                                            />
                                        </div>
                                        <div class="mb-4">
                                            <input
                                                type="password"
                                                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                                placeholder="Password"
                                                bind:value={password}
                                            />
                                        </div>
                                        {#if errorMessage != undefined && errorMessage !== "success"}
                                            <div
                                                class="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800"
                                                role="alert"
                                            >
                                                <span class="font-medium"
                                                    >Login Failed!</span
                                                >
                                                <span class="capitalize">
                                                    {errorMessage.message}!
                                                </span>
                                            </div>
                                        {/if}
                                        {#if errorMessage === "success"}
                                            <div
                                                class="p-4 mb-4 text-sm text-green-700 bg-green-100 rounded-lg dark:bg-green-200 dark:text-green-800 relative"
                                                role="alert"
                                            >
                                                <span
                                                    class="font-medium align-bottom"
                                                    >Logging in!</span
                                                >
                                                <div
                                                    class="absolute inline-block top-2 pl-3"
                                                >
                                                    <div
                                                        class="relative inline-block py-2"
                                                    >
                                                        <div
                                                            class="w-2 h-2 border-purple-200 border-2 rounded-full"
                                                        />
                                                        <div
                                                            class="w-2 h-2 border-purple-700 border-t-2 animate-spin rounded-full absolute left-0 top-2"
                                                        />
                                                    </div>
                                                    <div
                                                        class="relative inline-block py-1"
                                                    >
                                                        <div
                                                            class="w-4 h-4 border-purple-200 border-2 rounded-full"
                                                        />
                                                        <div
                                                            class="w-4 h-4 border-purple-700 border-t-2 animate-spin rounded-full absolute left-0 top-1"
                                                        />
                                                    </div>
                                                    <div
                                                        class="relative inline-block"
                                                    >
                                                        <div
                                                            class="w-8 h-8 border-purple-200 border-2 rounded-full"
                                                        />
                                                        <div
                                                            class="w-8 h-8 border-purple-700 border-t-2 animate-spin rounded-full absolute left-0 top-0"
                                                        />
                                                    </div>
                                                </div>
                                            </div>
                                        {/if}
                                        <div
                                            class="text-center pt-1 mb-12 pb-1"
                                        >
                                            <button
                                                class="inline-block px-6 py-2.5 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out w-full mb-3"
                                                type="button"
                                                data-mdb-ripple="true"
                                                data-mdb-ripple-color="light"
                                                on:click={loginRequest}
                                            >
                                                Log in
                                            </button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                            <div
                                class="side lg:w-6/12 lg:rounded-r-lg rounded-b-lg lg:rounded-bl-none content-center place-items-center grid"
                            >
                                <div
                                    class="text-white px-4 py-6 md:p-12 md:mx-6"
                                >
                                    <p class="text-sm ali">
                                        a-well-a don't you know about the bird?<br
                                        />
                                        well, everybody knows that the bird is the
                                        word <br />
                                        a-well, a bird, bird, bird, well-a bird is
                                        the word
                                    </p>

                                    <iframe
                                        src="https://www.youtube.com/embed/9Gc4QTqslN4"
                                        title="YouTube video player"
                                        frameborder="0"
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                        allowfullscreen
                                    />

                                    <h4 class="text-xl font-semibold mb-6">
                                        AND
                                    </h4>
                                    <p class="text-sm">
                                        bad boys bad boys whatcha gonna do<br />
                                        whatcha gonna do when they come for you
                                    </p>

                                    <iframe
                                        src="https://www.youtube.com/embed/Zd68AthoNIw"
                                        title="YouTube video player"
                                        frameborder="0"
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                        allowfullscreen
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</div>

<style>
    button,
    .side {
        background: linear-gradient(
            to right,
            #ee7724,
            #d8363a,
            #dd3675,
            #b44593
        );
    }

    iframe {
        padding-top: 10px;
        padding-bottom: 10px;
    }
</style>
