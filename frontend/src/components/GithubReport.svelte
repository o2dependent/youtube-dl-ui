<script lang="ts">
	import GithubIcon from "@/icons/GithubIcon.svelte";
	import { BrowserOpenURL } from "wailsjs/runtime/runtime";
</script>

<button
	class="text-foreground"
	on:click={() =>
		BrowserOpenURL("https://github.com/o2dependent/youtube-dl-ui/issues")}
>
	<GithubIcon class="size-8" />
</button>
