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

<div class="flex flex-col items-center h-full">
	<div class="prose container mx-auto px-4 w-full py-6">
		<h1>Welcome to Lowky Youtube-dl UI</h1>
		<!-- FORM -->
		<form on:submit|preventDefault={submit} class="w-full flex gap-2">
			<div class="relative w-full">
				<span class="sr-only">Width</span>
				<span
					aria-hidden
					class="absolute left-5 top-4 text-xxs text-muted-foreground">URL</span
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
		<pre class="max-w-md">{JSON.stringify(info, null, 2)}</pre>
	{/if}
</div>
