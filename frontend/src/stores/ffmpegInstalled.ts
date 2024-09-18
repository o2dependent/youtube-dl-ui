import { writable } from "svelte/store";
import { CheckFFMPEG } from "wails/go/main/App";

export const ffmpegInstalled = writable<boolean>(false);
export const ffmpegInstallationError = writable<boolean>(false);

export const checkFFMPEG = async () => {
	try {
		const isFfmpegInstalled = await CheckFFMPEG();
		ffmpegInstalled.set(isFfmpegInstalled);
		ffmpegInstallationError.set(false);
	} catch (error) {
		ffmpegInstallationError.set(true);
	}
};
