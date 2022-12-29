<script>
    import { onMount } from "svelte";
    import { result } from "../../stores";
    let dat;

    result.subscribe((value) => (dat = value));

    onMount(
        async () =>
            await fetch("/api/v1/dashboard", {
                method: "GET",
                credentials: "include",
            })
                .then((response) => response.json())
                .then((data) => result.set(data))
                .catch((error) => {
                    console.error("Error:", error);
                })
    );
</script>

<h1 class="font-bold">{JSON.stringify(dat)}</h1>
