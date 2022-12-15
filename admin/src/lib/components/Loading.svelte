<script>
    import LoadingAnimation from "../assets/loading-animation.json";
    import "@lottiefiles/lottie-player";
</script>

<div
    class="bg-[rgba(0,0,0,0.4)] h-screen w-full fixed top-0 right-0 left-0 items-center"
>
    <div class="grid h-screen place-items-center">
        <lottie-player
            src={LoadingAnimation}
            style="width:400px; height: 400px;"
            autoplay={true}
            loop={true}
            speed={1.5}
        />
    </div>
</div>
