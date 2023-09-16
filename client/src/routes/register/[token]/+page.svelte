<script lang="ts">
    import Logo from "$lib/components/Logo/Logo.svelte";
    import { goto } from "$app/navigation";
	import { onMount } from "svelte";

    let countdown = 4;
    onMount(() => {
        const countdownInterval = setInterval(() => {
            countdown--;
            if (countdown === 0) {
                clearInterval(countdownInterval);
                goto("/");
            }
        }, 1000)
    })
</script>

    <div class="w-[500px] h-[180px] bg-zinc-900/[.8] outline outline-1 outline-zinc-600 rounded-lg px-2  max-sm:rounded-none max-sm:h-[220px] max-sm:w-[95%]">
        <div class="flex flex-col items-center h-full py-3">
            <div class="flex items-center justify-start w-full h-[35%] px-3 max-sm:h-fit">
                <div class="flex flex-row items-center justify-center gap-x-3 max-sm:flex-col max-sm:items-start max-sm:gap-y-3">
                    <div class="flex outline outline-[1px] outline-zinc-500 w-[45px] h-[45px]  bg-black  rounded-[8px] items-center justify-center">
                        <Logo className="stroke-white stroke-[2px] fill-transparent w-[70%] h-[70%] fill" />
                    </div>
                    <div class="relative flex flex-col items-start justify-start bottom-[2px]">
                        <h1 class="text-zinc-100 font-bold text-[18px] text-center">Adresse valide</h1>
                        <p class="text-sm text-zinc-400">Votre adresse mail est vérifiée avec succès !</p>
                    </div>
                </div>
            </div>
            <div class="flex items-end justify-center flex-1 w-full px-3">
                <button class="w-full outline outline-[1px] outline-zinc-600 rounded-[7px] text-sm font-medium text-zinc-300 h-[40px] bg-black">
                    Redirection dans {countdown}
                </button>
            </div>
        </div>
    </div>
