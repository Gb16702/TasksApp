<script lang="ts">
    import {onMount} from "svelte";
    import LoadingPage from "$lib/components/Spinners/LoadingPage.svelte";
    import Logo from "$lib/components/Logo/Logo.svelte";
    import Add from "$lib/components/Add.svelte";
    import { enhance, applyAction } from "$app/forms";
    import type { PageData, SubmitFunction } from "./$types";
    import Delete from "$lib/components/Delete.svelte";
    import Complete from "$lib/components/Complete.svelte";

    let loading: boolean = true

    let value: string = "";

    export let data:PageData;

    const addTodo = (input : any) => {
        console.log(input);

        return async ({update} : {update : any}) => {
            await update()
        }
    }

    console.log(data);


</script>


<div class="min-w-[500px] rounded-[6px] px-3">
        <form class="w-full" method="POST" action="?/addTodo" use:enhance={addTodo}>
            <Add bind:value={value} />
        </form>

        {#if data && data.tasks.length != 0}
        <div class="flex items-center flex-col gap-y-2 mt-4 bg-[#131315] outline-1 outline outline-zinc-600 rounded-[5px] py-2 justify-center">
            {#each data.tasks as t_}
                <div class={`h-[40px] flex items-center justify-between px-3 white w-[98%] rounded-lg ${t_.done && "bg-zinc-950"}`}>
                    <div class="flex flex-row items-center justify-center gap-x-5">
                        <form action="?/toggleStatus" method="POST" use:enhance={() => {
                            return async ({result, update}) => {
                                console.log(result);
                                await update()
                            }
                        }}>
                            <input type="hidden" name="todo" value={t_.id}>
                            <Complete done={t_.done} />
                        </form>
                        <form action="?/editName" method="POST" use:enhance={() => {
                            return async ({result, update}) => {
                                console.log(result);
                                await update()
                            }
                        }}>
                            <input name="id" type="hidden" value={t_.id}>
                            <input name="name" class={`text-sm font-medium text-zinc-200 bg-transparent outline-none ${t_.done && "line-through"}` } value={t_.name} />
                        </form>
                    </div>
                    <form action="?/deleteTodo" method="POST" class="flex items-center justify-center p-[4px] rounded-[4px] transition-colors duration-200 hover:bg-zinc-600/[.4] cursor-pointer" use:enhance={() => {
                        return async ({result, update}) => {
                            console.log(result);
                            await update()
                        }
                    }}>
                        <input type="hidden" name="todo" value={t_.id}>
                        <Delete />
                    </form>
                </div>
            {/each}
        </div>
    {/if}
    </div>
