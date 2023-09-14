<script lang="ts">
    export let done: boolean;
    const toggleChecked: () => void = () => done = !done;
</script>

<button type="submit" class="relative flex items-center cursor-pointer">
    <div class="checkbox" on:click={toggleChecked}></div>
    {#if done}
    <div class="absolute -translate-x-1/2 -translate-y-1/2 pointer-events-none top-1/2 left-1/2 ">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="white" class="w-[14px] h-[14px]">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.75l6 6 9-13.5" />
          </svg>
    </div>
    {/if}
</button>

<style>
    .checkbox{
    appearance: none;
    width: 14px;
    height: 14px;
    outline: 1px solid white;
    border-radius: 2px;
}
</style>