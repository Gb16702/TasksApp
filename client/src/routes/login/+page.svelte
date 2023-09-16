<script lang="ts">
    import Input from "../../lib/components/auth/Input.svelte"
    import Eye from "$lib/components/auth/Eye.svelte";
    import Logo from "$lib/components/Logo/Logo.svelte";
    import Button from "$lib/components/auth/Button.svelte";
	import CircleSecondary from "$lib/components/Spinners/CircleSecondary.svelte";
    import { enhance } from "$app/forms";
	import type { SubmitFunction } from "@sveltejs/kit";
	import type { ActionData, Snapshot } from "./$types";

    export let type: "email" | "password" | "text" = "password";
    export let form: ActionData;

    let loading: boolean = false

    let formData: {email: string, password: string} = {
        email: "",
        password: "",
    }

    export const snapshot: Snapshot = {
        capture: () => formData,
        restore: data => formData = data
    }

    const handleClick: () => void = () => type = type === "password" ? "text" : "password";

    const handleSubmit:SubmitFunction = () => {
        loading = true
        return async ({update}) => {
            await update({
                reset: false,
            })
        }
    }
</script>

<Logo className={`fill-[#fff] w-[38px] h-[38px] drop-shadow-[0_0_25px_white]`} />
<div class=" max-h-[450px] w-[440px] rounded-[8px]  py-4 px-5">
    <h1 class="text-zinc-100 font-bold text-[30px] text-center">Connexion à Tasks</h1>
    <form method="POST" class="py-5 mt-3" action="?/login" use:enhance={handleSubmit}>
        <div class="flex flex-col text-sm text-zinc-200">
            <label for="email" class="font-medium">Email</label>
            <!-- <input type="text" class="w-full mt-1 rounded-[5px] outline-1 bg-zinc-900/[.8] outline-zinc-700 h-[39px] px-2 transition-all duration-200 outline focus:outline-white" /> -->
            <Input name=email  bind:value="{formData.email}" type="email" placeholder={"alan.turing@exemple.com"} />
        </div>
        {#if form?.errors?.email}
            <div class="mt-1 text-sm text-red-400">
                {form?.errors?.email}
            </div>
        {/if}
        <span class="h-[20px] block"></span>
        <div class="flex flex-col text-sm text-zinc-200">
            <div class="flex items-center justify-between">
                <label for="email" class="font-medium">Mot de passe</label>
            </div>
            <div class="relative">
                <Input name={type} bind:value="{formData.password}" type={type} placeholder={"•••••••"} />
                <Eye onClick="{handleClick}" className="w-5 h-5 absolute right-[15px] top-[55%] -translate-y-1/2" />
            </div>
        </div>
        {#if form?.errors?.password}
            <div class="mt-1 text-sm text-red-400">
                {form?.errors?.password}
            </div>
        {/if}
        <span class="h-[20px] block"></span>
        <Button {loading}>
            {#if !loading}
                Valider
            {:else}
                <CircleSecondary />
            {/if}
        </Button>
    </form>
    <div class="flex items-center justify-center mt-1">
        <span class="text-sm text-zinc-300">Pas encore de compte ?</span>
        <a href="/register" class="ml-1 text-sm font-medium text-white">Inscription</a>
    </div>

</div>