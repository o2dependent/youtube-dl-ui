<script lang="ts">
	import { flyAndScale } from "@/utils/transitions";

	import { Dialog } from "bits-ui";
	import { fade } from "svelte/transition";

	export let open: boolean;
	export let text: string = "Downloading";
</script>

<Dialog.Root {open} closeOnOutsideClick={false} closeOnEscape={false}>
	<Dialog.Portal>
		<Dialog.Overlay
			transition={fade}
			transitionConfig={{ duration: 150 }}
			class="fixed inset-0 z-50 bg-black/80"
		/>
		<Dialog.Content
			transition={flyAndScale}
			class="fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] rounded-card-lg border bg-background p-5 shadow-popover outline-none sm:max-w-[490px] md:w-full h-full max-h-[260px] sm:max-h-[360px] flex flex-col gap-2 justify-center items-center"
		>
			<div class="loader"></div>
			<p>{text}</p>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>

<style>
	.loader {
		width: 65px;
		aspect-ratio: 1;
		position: relative;
	}
	.loader:before,
	.loader:after {
		content: "";
		position: absolute;
		border-radius: 50px;
		box-shadow: 0 0 0 3px inset #fff;
		animation: l5 2.5s infinite;
	}
	.loader:after {
		animation-delay: -1.25s;
		border-radius: 0;
	}
	@keyframes l5 {
		0% {
			inset: 0 35px 35px 0;
		}
		12.5% {
			inset: 0 35px 0 0;
		}
		25% {
			inset: 35px 35px 0 0;
		}
		37.5% {
			inset: 35px 0 0 0;
		}
		50% {
			inset: 35px 0 0 35px;
		}
		62.5% {
			inset: 0 0 0 35px;
		}
		75% {
			inset: 0 0 35px 35px;
		}
		87.5% {
			inset: 0 0 35px 0;
		}
		100% {
			inset: 0 35px 35px 0;
		}
	}
</style>
