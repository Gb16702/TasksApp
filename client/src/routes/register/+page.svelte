<script lang="ts">
    import Input from "../../lib/components/auth/Input.svelte"
    import Eye from "$lib/components/auth/Eye.svelte";
    import Logo from "$lib/components/Logo/Logo.svelte";
    import Button from "$lib/components/auth/Button.svelte";
    import { enhance } from "$app/forms";
    import type { ActionData, Snapshot } from "./$types"
	import type { SubmitFunction } from "@sveltejs/kit";
    import * as Alert from "$lib/components/ui/alert";


    let formData: {email:string, password: string, passwordConfirm: string} = {
        email: "",
        password: "",
        passwordConfirm: ""
    }

    export const snapshot: Snapshot = {
        capture: () => formData,
        restore: data => formData = data
    }

    export let type: "email" | "password" | "text" = "password";

    const handleClick: () => void = () => type = type === "password" ? "text" : "password";

    let loading: boolean = false

    export let form: ActionData;

    let response = false
    let status: "success" | "error" | null = null
    let message: any;

    const handleSubmit: SubmitFunction = () => {
        loading = true
        return async ({result, update} : {result: any, update : any}) => {
            response = true;
            loading= false
            setTimeout(() => {
                response = false
            }, 3000);

            result.status == 200 ? status = "success" : status = "error"
            message = result.data.body ?? result.data
            await update()
        }
    }
</script>

{#if response}
    <Alert.Root
    variant="default"
    title={status == "success" ? "Succès" : "Erreur"}
    description={message}
    type="error" />
{/if}

<Logo className={`fill-[#fff] w-[38px] h-[38px] drop-shadow-[0_0_25px_white]`} />
<div class=" max-h-[450px] w-[440px] rounded-[8px]  py-4 px-5">
    <h1 class="text-zinc-100 font-bold text-[30px] text-center">Inscription à Tasks</h1>
    <form method="POST" class="py-5 mt-3" action="?/register" use:enhance={handleSubmit}>
        <div class="flex flex-col text-sm text-zinc-200">
            <label for="email" class="font-medium">Email</label>
            {#if form?.errors?.email}
                <Input name="email" placeholder={"alan.turing@example.com"} variant="error" bind:value={formData.email} />
            {:else}
                <Input name="email" bind:value="{formData.email}" placeholder={"alan.turing@exemple.com"}  />
            {/if}
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
            {#if  form?.errors?.password}
                <Input name=password placeholder={"•••••••"} type={type} variant="error" bind:value="{formData.password}" />
                {:else}
                <Input name=password type={type} placeholder={"•••••••"} bind:value="{formData.password}" />
            {/if}
                <Eye onClick="{handleClick}"  className="w-5 h-5 absolute right-[15px] top-[55%] -translate-y-1/2 resize-none" />
            </div>
        </div>
        {#if form?.errors?.password}
        <div class="mt-1 text-sm text-red-400">
            {form?.errors?.password}
        </div>
    {/if}
        <span class="h-[20px] block"></span>
        <div class="flex flex-col text-sm text-zinc-200">
            <div class="flex items-center justify-between">
                <label for="confirmPassword" class="font-medium">Confirmer le mot de passe</label>
            </div>
            <div class="relative">
                {#if form?.errors?.passwordConfirm}
                <Input  name="passwordConfirm" type={type} variant="error" bind:value="{formData.passwordConfirm}" />
                {:else}
                <Input name="passwordConfirm" placeholder={"•••••••"} type={type} bind:value="{formData.passwordConfirm}" />
                {/if}

                <Eye onClick="{handleClick}" className="w-5 h-5 absolute right-[15px] top-[55%] -translate-y-1/2" />
            </div>
        </div>
        {#if form?.errors?.passwordConfirm}
            <div class="mt-1 text-sm text-red-400">
                {form?.errors?.passwordConfirm}
            </div>
        {/if}

        <span class="h-[20px] block"></span>
        <Button {loading}>
            {#if !loading}
                Valider
            {/if}
        </Button>
    </form>
    <div class="flex items-center justify-center mt-1">
        <span class="text-sm text-zinc-300">Déjà membre ?</span>
        <a href="/login" class="ml-1 text-sm font-medium text-white">Connexion</a>
    </div>
    <span class="h-[20px] block"></span>
</div>