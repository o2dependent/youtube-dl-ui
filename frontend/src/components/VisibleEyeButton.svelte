<script lang="ts">
	import EyeClosed from "@/icons/EyeClosed.svelte";
	import EyeOpen from "@/icons/EyeOpen.svelte";
	import { Button } from "bits-ui";

	export let show: boolean;
</script>

<Button.Root class={$$props.class} on:click={() => (show = !show)}>
	{#if show}
		<EyeOpen />
	{:else}
		<EyeClosed />
	{/if}
</Button.Root>
