<script lang="ts">
  import Logo from "./Logo/Logo.svelte";
  import clickOutside from "../../utils/clickOutside";

  let open:boolean = false
  const handleClick: () => boolean = () => open = true
  const handleClickOutside: () => boolean = () => open = false
</script>

{#if open}
    <div class="absolute top-0 left-0 w-full h-full bg-black/[.3] backdrop-blur-sm z-10"></div>
        <div class="max-w-[700px] w-full max-sm:max-w-[500px] max-sm:w-[95%] h-[200px] absolute top-[15%] left-1/2 -translate-x-1/2 flex flex-col rounded-[5px] border borer-zinc-200 z-20 bg-black" use:clickOutside on:clickOutside={handleClickOutside}>
            <div class="h-[10%] p-[10px]">
                <div class="flex justify-center">
                    <Logo className="w-[48px] h-[45px] fill-white -translate-y-[34px] relative bg-black z-10 p-[11px] blur-[6px]" />
                    <Logo className="w-[48px] h-[45px] fill-white -translate-y-[34px]  bg-black z-10 p-[11px] absolute" />
                </div>
            </div>
            <div class="h-[90%] p-[12px] pt-[6px] text-white text-sm flex flex-col justify-between">
                <div class="flex flex-col font-medium cursor-pointer gap-y-3">
                    <a href="https://github.com/Gb16702?tab=repositories" target="_blank">
                        Github
                    </a>
                    <button class="w-fit" on:click={() => open=false}>
                        Fermer
                    </button>
                    <form action="?/logout" method="POST">
                        <button >
                            Se déconnecter
                        </button>
                    </form>
                </div>
                <div>
                    <small class=" text-zinc-300 text-[11px]">
                        Projet
                        <span class="font-semibold text-zinc-200">
                            personnel</span>, à usage <span class="font-semibold text-zinc-200">non commercial</span>
                    </small>
                </div>
            </div>
        </div>
{/if}
<div class={`absolute top-[23px] right-[20px] p-2 rounded-[5px] list-none text-white flex items-center justify-between cursor-pointer gap-x-[6px] transition duration-200 hover:bg-[#131315]`} on:click={handleClick}>
    <li class="text-sm font-semibold">
        Menu
    </li>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class={`relative w-[14px] h-[14px] top-[2px] stroke-white transition-transform duration-100 ${open && " rotate-180"}`}>
        <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
      </svg>
</div>
