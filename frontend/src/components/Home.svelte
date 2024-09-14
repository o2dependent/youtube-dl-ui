<script lang="ts">
	import { Button, type Selected } from "bits-ui";
	import { GetImportantInfo } from "wails/go/main/App";
	import InputSelect from "./InputSelect.svelte";
	import { capitalize } from "@/utils/capitalize";

	/**
	 * Data to be used for init info, download, and verification pre-download
	 */
	let url = "";
	let info: Awaited<ReturnType<typeof GetImportantInfo>> | null = null;

	/**
	 * Combobox inputs and handlers
	 */
	let fileExtInput = "";
	let fileExt: Selected<string>[] = [];

	let qualityInput = "";
	let quality: Selected<string>[] = [];

	let audioQualityInput = "";
	let audioQuality: Selected<string>[] = [];

	/**
	 * Functions
	 */
	const findVideoInfo = async () => {
		const i = await GetImportantInfo(url);
		if (i) {
			info = i;
			const newFileExt: string[] = [];
			const newQuality: string[] = [];
			const newAudioQuality: string[] = [];

			for (let i = 0; i < info.qualityInfo.length; i++) {
				const el = info.qualityInfo[i];
				if (!newAudioQuality.includes(el.audioQuality)) {
					newAudioQuality.push(el.audioQuality);
				}
				const ext = el.mimeType.split(";")[0].split("/")[1];
				if (!newFileExt.includes(ext)) {
					newFileExt.push(ext);
				}
				if (!newQuality.includes(el.quality)) {
					newQuality.push(el.quality);
				}
			}
			(audioQuality = newAudioQuality.map((v) => ({
				value: v,
				label: v ? v.replace("AUDIO_QUALITY_", "").toLowerCase() : "none",
			}))),
				(fileExt = newFileExt.map((v) => ({ value: v, label: v })));
			quality = newQuality.map((v) => ({ value: v, label: v }));
		} else {
			info = null;
			audioQuality = [];
			fileExt = [];
			quality = [];
		}
	};

	const formatQuality = (q: string): string => {
		if (q.includes("hd")) return q;
		else if (q === "large") return "480p";
		else if (q === "medium") return "360p";
		else if (q === "small") return "240p";
		else if (q === "tiny") return "144p";
		return q;
	};
</script>

<div class="flex flex-col items-center h-full w-full">
	<div class="prose prose-invert container mx-auto px-4 w-full py-6">
		<h1 class="w-full text-center">Lowky youtube-dl UI</h1>
		<!-- FORM -->
		<form on:submit|preventDefault={findVideoInfo} class="w-full flex gap-2">
			<div class="relative w-full">
				<span class="sr-only">Width</span>
				<span
					aria-hidden
					class="absolute left-3 top-4 text-xxs text-muted-foreground">URL</span
				>
				<input
					class="h-input w-full rounded-10px border border-border-input bg-background pl-10 pr-2 text-sm text-foreground"
					bind:value={url}
				/>
			</div>
			<Button.Root
				type="submit"
				class="inline-flex h-12 items-center justify-center rounded-input bg-dark
			px-[21px] text-[15px] font-semibold text-background shadow-mini
			hover:bg-dark/95 active:scale-98 active:transition-all"
			>
				Submit
			</Button.Root>
		</form>
	</div>
	<!-- INFO -->
	{#if info !== null}
		<div class="max-w-prose mx-auto flex flex-col gap-2 w-full px-4">
			<div class="grid sm:grid-cols-2 gap-2">
				<img
					class="rounded-card border border-muted shadow-md md:col-span-1"
					alt={info.title}
					srcset={info.thumbnails?.map((t) => `${t.URL} ${t.Width}w`).join(",")}
					sizes={info?.thumbnails
						?.map((t, i) =>
							i === (info?.thumbnails?.length ?? 1) - 1
								? `${t.Width}px`
								: `(max-width: ${t.Width * 2}px) ${t.Width}px`,
						)
						.join(",")}
				/>
				<div class="md:col-span-1 flex flex-col gap-1 pt-2">
					<h2 class="text-2xl sm:text-3xl">{info.title}</h2>
					<p class="opacity-50">{info.author}</p>
					<p class="opacity-50">{info.duration}</p>
				</div>
			</div>
			<form
				class="flex flex-col gap-2"
				on:submit|preventDefault={() => {
					console.log(audioQualityInput, qualityInput, fileExtInput);
				}}
			>
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-2">
					<InputSelect
						bind:items={audioQuality}
						bind:inputValue={audioQualityInput}
						name="audioQuality"
						placeholder="Audio Quality"
					/>
					<InputSelect
						bind:items={quality}
						bind:inputValue={qualityInput}
						name="quality"
						placeholder="Video Quality"
					/>
					<InputSelect
						bind:items={fileExt}
						bind:inputValue={fileExtInput}
						name="fileExt"
						placeholder="File Extension"
					/>
				</div>

				<Button.Root
					type="submit"
					class="inline-flex h-12 items-center justify-center rounded-input bg-dark
			px-[21px] text-[15px] font-semibold text-background shadow-mini
			hover:bg-dark/95 active:scale-98 active:transition-all"
				>
					Download
				</Button.Root>
			</form>
			<!-- <div class="flex flex-col gap-1">
				{#each info.qualityInfo as quality, i}
					<div
						class="flex gap-2 rounded-card border border-muted bg-background-alt p-3 shadow-card"
					>
						<div class="flex-grow items-center flex gap-1">
							<p class="flex-grow">
								{`${formatQuality(quality.quality)}.${quality.mimeType.split(";")[0].split("/")[1]}`}
							</p>
							{#if quality.mimeType.split(";")[0].split("/")[0] !== "video"}
								<VideoOff class="h-5 w-5 opacity-50" />
							{:else}
								<div class="w-5 h-5" />
							{/if}
							{#if !quality.audioQuality}
								<VolumeMuted class="w-5 h-5 opacity-50" />
							{:else}
								<div class="w-5 h-5" />
							{/if}
						</div>
						<Button.Root
							class="inline-flex h-8 items-center justify-center rounded-input bg-dark
						px-[14px] text-[12px] font-semibold text-background shadow-mini
						hover:bg-dark/95 active:scale-98 active:transition-all">Download</Button.Root
						>
					</div>
				{/each}
			</div> -->
		</div>
	{/if}
</div>
