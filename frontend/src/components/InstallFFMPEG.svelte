<script lang="ts">
	import { Button } from "bits-ui";
	import { InstallFFmpeg } from "wails/go/main/App";
	import Downloading from "./Downloading.svelte";
	import { BrowserOpenURL } from "wails/runtime/runtime";

	export let recheckFFMPEG: () => Promise<void>;

	let downloading = false;
	let rechecking = false;

	const install = async () => {
		await InstallFFmpeg();
		await recheckFFMPEG();
	};
</script>

<div class="flex flex-col items-center h-full w-full">
	<div class="prose prose-invert container mx-auto px-4 w-full py-6">
		<h1 class="w-full text-center">FFMPEG needs to be installed</h1>
		<p>
			If this process does not work please visit <button
				class="text-blue-500"
				on:click={() => BrowserOpenURL("https://www.ffmpeg.org/")}
			>
				ffmpeg.org
			</button> and recheck FFMPEG.
		</p>

		<Button.Root
			disabled={downloading}
			on:click={recheckFFMPEG}
			type="button"
			class="mx-auto inline-flex h-12 items-center justify-center rounded-input bg-dark
			px-[21px] text-[15px] font-semibold text-background shadow-mini
			hover:bg-dark/95 active:scale-98 active:transition-all disabled:opacity-50 disabled:cursor-not-allowed"
		>
			Recheck FFMPEG
		</Button.Root>
		<Button.Root
			disabled={downloading}
			on:click={install}
			type="button"
			class="mx-auto inline-flex h-12 items-center justify-center rounded-input bg-dark
			px-[21px] text-[15px] font-semibold text-background shadow-mini
			hover:bg-dark/95 active:scale-98 active:transition-all disabled:opacity-50 disabled:cursor-not-allowed"
		>
			Download
		</Button.Root>
	</div>
</div>
<Downloading
	open={downloading || rechecking}
	text={rechecking ? "Rechecking FFMPEG" : "Downloading"}
/>
