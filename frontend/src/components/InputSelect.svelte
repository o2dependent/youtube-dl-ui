<script lang="ts">
	import Check from "@/icons/Check.svelte";
	import Expand from "@/icons/Expand.svelte";
	import { flyAndScale } from "@/utils/transitions";
	import { Label, Select, type Selected } from "bits-ui";

	export let items: Selected<string>[];
	export let inputValue: string;
	export let name: string;
	export let placeholder: string;
	export let disabled: boolean = false;

	let selected: Selected<string> = { value: "", label: "" };
	$: inputValue = selected?.value ?? "";
</script>

<Select.Root {disabled} {items} bind:selected>
	<Select.Trigger
		class="inline-flex h-input w-full items-center rounded-9px border border-border-input bg-background px-[11px] text-sm transition-colors placeholder:text-foreground-alt/50  focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-offset-2 focus:ring-offset-background"
		aria-label="Select a item"
	>
		<Select.Value class="text-sm" {placeholder} />
		<Expand class="ml-auto size-6 text-muted-foreground" />
	</Select.Trigger>
	<Select.Content
		class="w-full rounded-xl border border-muted bg-background px-1 py-3 shadow-popover outline-none"
		transition={flyAndScale}
		sideOffset={8}
	>
		{#each items as item}
			<Select.Item
				class="flex h-10 w-full select-none items-center rounded-button py-3 pl-5 pr-1.5 text-sm outline-none transition-all duration-75 data-[highlighted]:bg-muted"
				value={item.value}
				label={item.label}
			>
				{item.label}
				<Select.ItemIndicator class="ml-auto" asChild={false}>
					<Check class="size-6" />
				</Select.ItemIndicator>
			</Select.Item>
		{/each}
	</Select.Content>
	<Select.Input {name} />
</Select.Root>
