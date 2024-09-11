<script lang="ts">
	import { Button } from "bits-ui";
	import { GetImportantInfo } from "wails/go/main/App";

	let url = "";
	let info: Awaited<ReturnType<typeof GetImportantInfo>> | null = null;

	const submit = async () => {
		const i = await GetImportantInfo(url);
		if (i) {
			info = i;
		}
	};
</script>

<div class="flex flex-col items-center h-full w-full">
	<div class="prose container mx-auto px-4 w-full py-6">
		<h1>Welcome to Lowky Youtube-dl UI</h1>
		<!-- FORM -->
		<form on:submit|preventDefault={submit} class="w-full flex gap-2">
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
			hover:bg-dark/95 active:scale-98 active:transition-all">Submit</Button.Root
			>
		</form>
	</div>
	<!-- INFO -->
	{#if info !== null}
		<div class="container mx-auto flex flex-col gap-2 w-full px-4">
			<div class="grid sm:grid-cols-2 gap-2">
				<img
					class="rounded-card-lg shadow-md md:col-span-1"
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
				<div class="md:col-span-1 flex flex-col gap-1">
					<h2 class="text-4xl">{info.title}</h2>
					<p>{info.author}</p>
					<p>{info.duration}</p>
					<p>{info.time}</p>
				</div>
			</div>
			<div class="flex flex-col gap-1">
				{#each info.qualityInfo as quality}
					<div
						class="rounded-card border border-muted bg-background-alt p-3 shadow-card"
					>
						{JSON.stringify(quality)}
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
