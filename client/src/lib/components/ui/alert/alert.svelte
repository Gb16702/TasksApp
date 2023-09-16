<script lang="ts">
    import { cn } from "$lib/utils";
    import type { HTMLAttributes } from "svelte/elements";
    import { alertVariants, type Variant } from ".";
    import { fly, slide } from "svelte/transition";
	import Error from "$lib/components/ui/alert/Error.svelte";
	import Success from "$lib/components/ui/alert/Success.svelte";

    type $$Props = HTMLAttributes<HTMLDivElement> & {
        variant?: Variant;
        title: string;
        description: string;
		type?: "success" | "error" | null;
    };

    let className: $$Props["class"] = undefined;
    export let variant: $$Props["variant"] = "default";
    export let title: $$Props["title"];
    export let description: $$Props["description"];
	export let type: $$Props["type"] = null;

	export { className as class };
</script>

<div
    class={cn(alertVariants({ variant }), className, "fixed bottom-4 right-4 max-w-[35%] w-[400px] max-sm:right-1/2 max-sm:translate-x-1/2  max-sm:max-w-[95%] max-sm:w-full bg-black outline outline-1 outline-zinc-600")}
    {...$$restProps}
    role="alert"
    out:fly={{
        delay: 300,
        duration: 300,
        x: 100,
    }}
>
    <div class="flex flex-row items-start justify-start gap-x-2">
		<div class="relative top-[2px]">
			{#if type == "success"}
			<Success className="w-[23px] h-[23px] stroke-white" />
			{:else if type == "error"}
			<Error className="w-[23px] h-[23px] stroke-white" />
			{/if}
		</div>
		<div>
			<h2 class="text-base font-semibold text-zinc-100">{title}</h2>
			<h4 class="text-sm leading-relaxed text-zinc-300">
				{description}
			</h4>
		</div>
    </div>
</div>