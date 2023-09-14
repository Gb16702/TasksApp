<div class="absolute top-[60%] left-1/2 -translate-x-1/2 -translate-y-1/2 flex gap-x-[6px] items-center justify-center">
    <span class="inline-block w-[5px] h-[5px] bg-black rounded-full"></span>
    <span class="inline-block w-[5px] h-[5px] bg-black rounded-full"></span>
    <span class="inline-block w-[5px] h-[5px] bg-black rounded-full"></span>
</div>

<style>
    span {
        animation: moveup .6s linear infinite alternate
    }

    @keyframes moveup {
        100% {
            transform: translateY(-8px)
        }
    }

    span:nth-child(2) {
        animation-delay: .25s
    }

    span:nth-child(3) {
        animation-delay: .5s
    }

</style>