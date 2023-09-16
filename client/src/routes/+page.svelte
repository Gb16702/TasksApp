<script lang="ts">
    import Add from "$lib/components/Add.svelte";
    import { enhance } from "$app/forms";
    import type { PageData} from "./$types";
    import Delete from "$lib/components/Delete.svelte";
    import Complete from "$lib/components/Complete.svelte";
    import Button from "$lib/components/auth/Button.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import { fly, slide } from "svelte/transition";
    import * as Alert from "$lib/components/ui/alert";

    let value: string = "";

    export let data:PageData;

    let response = false
    let status: "success" | "error" | null = null
    let message: any;

    const addTodo = () => {
        return async ({result, update} : {result: any, update : any}) => {
            response = true;
            setTimeout(() => {
                response = false
            }, 3000);

            result.status == 200 ? status = "success" : status = "error"
            message = result.data.body ?? result.data
            await update()
        }
    }

    const deleteTodo = () => {
        return async ({result, update}: {result: any, update: any}) => {
            response = true;
            setTimeout(() => {
                response = false
            }, 2500);

            result.status == 200 ? status = "success" : status = "error"
            message = result.data.body ?? result.data
            await update()
        }
    }

    const toggleStatus = () => {
        return async ({result, update}: {result: any, update: any}) => {
            response = true;
            setTimeout(() => {
                response = false
            }, 2500);

            result.status == 200 ? status = "success" : status = "error"
            message = result.data.body ?? result.data
            await update()
        }
    }

    const clearTodo = () => {
        return async ({result, update}: {result: any, update: any}) => {
            response = true;
            setTimeout(() => {
                response = false
            }, 2500);

            result.status == 200 ? status = "success" : status = "error"
            message = result.data.body ?? result.data
            await update()
        }
    }
</script>

<Menu />

{#if response}
    <Alert.Root
    variant="default"
    title={status == "success" ? "Succès" : "Erreur"}
    description={message}
    type="error" />
{/if}

<div class="min-w-[500px] rounded-[6px] px-3 max-sm:w-[100%] max-sm:min-w-max">
        <form class="w-full" method="POST" action="?/addTodo" use:enhance={addTodo}>
            <Add bind:value={value} />
        </form>

        {#if data && data.tasks.length != 0}
        <div class="flex flex-col items-center justify-center py-2 mt-3 gap-y-1 rounded-[5px] bg-[#131313] outline outline-1 outline-zinc-600" in:fly={{ y: 20 }} out:slide>
            {#each data.tasks as t_, index}
                <div in:fly={{ y: 20 }} out:slide class={`h-[44px] flex items-center justify-between px-3 white w-[98%] rounded-lg transition duration-200 hover:bg-zinc-900 ${t_.done && "opacity-50"}`}>
                    <div class="flex flex-row items-center justify-center gap-x-5">
                        <form action="?/toggleStatus" method="POST" use:enhance={toggleStatus}>
                            <input type="hidden" name="todo" value={t_.id}>
                            <Complete done={t_.done} />
                        </form>
                        <form action="?/editName" method="POST" use:enhance={() => {
                            return async ({result, update}) => {
                                await update()
                            }
                        }}>
                            <input name="id" type="hidden" value={t_.id}>
                            <input autocomplete="off" name="name" class={`text-sm font-medium text-zinc-200 bg-transparent outline-none ${t_.done && "line-through"}` } value={t_.name} />
                        </form>
                    </div>
                    <form action="?/deleteTodo" method="POST" class="flex items-center justify-center p-[4px] rounded-[4px] transition-colors duration-200 hover:bg-zinc-600/[.4] cursor-pointer"  use:enhance={deleteTodo}>
                        <input type="hidden" name="todo" value={t_.id}>
                        <Delete />
                    </form>
                </div>
            {/each}
        </div>
        {#if data.tasks.length >= 2}
        <form action="?/clear" method="POST" class="mt-3 w-[160px] max-sm:w-full" use:enhance={clearTodo}>
            <input name="userId" type="hidden" value={data.id}>
            <Button variant="clear">
                Nettoyer
            </Button>
        </form>
        {/if}
    {/if}
</div>