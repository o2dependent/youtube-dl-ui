import { get, writable } from "svelte/store";

export const createErrorsStore = <
	T extends Record<string, (boolean | number | string) | null>,
>(
	init: T,
) => {
	const errors = writable<T>(init);

	const reset = () => errors.set(init);
	const set = (v: T) => errors.set(v);
	const setKey = (key: keyof T, v: T[keyof T]) =>
		errors.update((prev) => ({ ...prev, [key]: v }));
	const resetAndUpdate = (v: Partial<T>) => errors.set({ ...init, ...v });
	const update = (v: Partial<T>) =>
		errors.update((prev) => ({ ...prev, ...v }));

	return {
		set,
		reset,
		setKey,
		resetAndUpdate,
		update,
		$: errors,
	};
};
