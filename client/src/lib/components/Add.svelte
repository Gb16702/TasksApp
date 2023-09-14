<script lang="ts">
    export let value: string = "";
    export let onKeyDown: (e: KeyboardEvent) => void = () => {};
</script>

<input name="name" bind:value={value} on:keydown={onKeyDown} type="text" placeholder="Entre une nouvelle tâche" class="w-full y-[6px] text-sm h-[40px] placeholder-zinc-400/[.8] mt-1 rounded-[5px] outline-1 bg-zinc-900/[.8] outline-zinc-600 px-2 transition-all duration-200 outline focus:outline-white text-white" />

