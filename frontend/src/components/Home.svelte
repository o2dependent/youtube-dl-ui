<script lang="ts">
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import App from "./App.svelte";
	import { checkFFMPEG, ffmpegInstalled } from "@/stores/ffmpegInstalled";

	let loading = true;
	let installationError = false;

	onMount(async () => {
		await checkFFMPEG();
		loading = false;
	});
</script>

{#if loading}
	<p
		in:fade={{ delay: 250, duration: 500 }}
		class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 animate-pulse text-2xl text-center w-full"
	>
		Checking FFMPEG
	</p>
{:else if $ffmpegInstalled}
	<App />
{:else}
	<div>
		<p>Something went wrong!</p>
		<p>Contact the developer for help!</p>
		<a href="mailto:131eolsen@gmail.com">131eolsen@gmail.com</a>
	</div>
{/if}
