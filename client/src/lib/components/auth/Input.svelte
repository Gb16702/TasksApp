<script lang="ts">
    export let type: InputProps["type"] = "text";
    export let name:InputProps["type"];
    export let required:boolean = false;
    export let variant: "default" | "error" = "default";
    export let value = "";
    export let placeholder: string | null = null;
</script>

{#if type === "text"}
    <input bind:value="{value}" type="text" name={name} {required} autocomplete="off" placeholder={placeholder} class={`w-full placeholder-zinc-400/[.8] mt-1 rounded-[5px] outline-1 bg-zinc-900/[.8] outline-zinc-600 px-2 transition-all duration-200 outline focus:outline-white h-[38px] ${variant == "error" && "outline-red-400"}`} />
{:else if type === "email"}
    <input bind:value="{value}" type="email" name={name} {required} autocomplete="off" placeholder={placeholder} class={`w-full placeholder-zinc-400/[.8] mt-1 rounded-[5px] outline-1 bg-zinc-900/[.8] outline-zinc-600 px-2 transition-all duration-200 outline focus:outline-white h-[38px] ${variant == "error" && "outline-red-400"}`} />
{:else if type === "password"}
    <input bind:value="{value}" type="password" name={name} {required} placeholder={placeholder} autocomplete="off" class={`w-full placeholder-zinc-400/[.8] mt-1 rounded-[5px] outline-1 bg-zinc-900/[.8] outline-zinc-600 px-2 transition-all duration-200 outline focus:outline-white h-[38px] ${variant == "error" && "outline-red-400"}`} />
{/if}