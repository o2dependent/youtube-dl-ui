<script lang="ts">
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import App from "./App.svelte";
	import InstallFfmpeg from "./InstallFFMPEG.svelte";
	import { CheckFFMPEG } from "wails/go/main/App";

	let ffmpegInstalled = false;
	let loading = true;

	onMount(async () => {
		try {
			ffmpegInstalled = await CheckFFMPEG();
		} catch (error) {}
		loading = false;
	});

	const recheckFFMPEG = async () => {
		try {
			ffmpegInstalled = await CheckFFMPEG();
		} catch (error) {}
	};
</script>

{#if loading}
	<p
		in:fade={{ delay: 250, duration: 500 }}
		class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 animate-pulse text-2xl text-center w-full"
	>
		Checking FFMPEG
	</p>
{:else if ffmpegInstalled}
	<App />
{:else}
	<InstallFfmpeg {recheckFFMPEG} />
{/if}
