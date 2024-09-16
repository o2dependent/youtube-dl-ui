<script lang="ts">
	import { flyAndScale } from "@/utils/transitions";
	import { AlertDialog } from "bits-ui";
	import { fade } from "svelte/transition";

	export let open: boolean;
	export let message: string;
	export let onOpenChange: () => void;
</script>

<AlertDialog.Root {open} {onOpenChange}>
	<AlertDialog.Portal>
		<AlertDialog.Overlay
			transition={fade}
			transitionConfig={{ duration: 150 }}
			class="fixed inset-0 z-50 bg-black/80"
		/>
		<AlertDialog.Content
			transition={flyAndScale}
			class="fixed left-[50%] top-[50%] z-50 grid w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] gap-4 rounded-card-lg border bg-background p-7 shadow-popover outline-none sm:max-w-lg md:w-full"
		>
			<div class="flex flex-col gap-4 pb-6">
				<AlertDialog.Title class="text-lg font-semibold tracking-tight"
					>Errors</AlertDialog.Title
				>
				<AlertDialog.Description class="text-sm">
					<pre class="text-red-500">{message}</pre>
				</AlertDialog.Description>
			</div>
			<div class="flex w-full items-center justify-center gap-2">
				<AlertDialog.Action
					class="inline-flex h-input w-full items-center justify-center rounded-input bg-dark text-[15px] font-semibold text-background shadow-mini transition-all hover:bg-dark/95 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-dark focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
					>Close</AlertDialog.Action
				>
			</div>
		</AlertDialog.Content>
	</AlertDialog.Portal>
</AlertDialog.Root>
