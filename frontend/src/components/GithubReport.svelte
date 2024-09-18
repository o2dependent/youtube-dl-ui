<script lang="ts">
	import GithubIcon from "@/icons/GithubIcon.svelte";
	import { BrowserOpenURL } from "wailsjs/runtime/runtime";
</script>

<button
	class="text-foreground flex gap-1 items-center opacity-50 hover:opacity-100 transition-opacity duration-300"
	on:click={() =>
		BrowserOpenURL("https://github.com/o2dependent/youtube-dl-ui/issues")}
>
	<GithubIcon class="size-6" />
	<p class="m-0 text-sm">Report issues</p>
</button>
