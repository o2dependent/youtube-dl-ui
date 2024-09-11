import { writable } from "svelte/store";

export const crauth = writable<null | { email: string; password: string }>(
	null,
);
