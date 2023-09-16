<script lang="ts">
	import CircleSecondary from "../Spinners/CircleSecondary.svelte";

    export let loading: boolean | undefined = false;
    export let variant: "default" | "error" | "clear" = "default";
</script>

<button type="submit" aria-busy={loading} class={`w-full bg-[#EDEDED] h-[38px] rounded-[5px] px-2 sm:font-medium font-semibold hover:brightness-[.85] transition duration-200  flex justify-center items-center relative ${loading ? "brightness-[.85]" : "brightness-100"} ${variant == "error" && "bg-red-400 text-zinc-200"} ${variant == "clear" && "text-sm max-sm:h-[42px] h-[40px] font-bold "}`}>
  <slot />
    {#if loading}
        <div class="absolute top-0 left-0 flex items-center justify-center w-full h-full">
        <CircleSecondary />
        </div>
    {/if}
</button>
