
export const apiUrl = (path: string): string => {
	if (path.startsWith("/")) {
		return settings.API_URL + path;
	} else {
		return settings.API_URL + "/" + path;
	}
}

interface Settings {
	API_URL: string
	DEBUG: boolean
}

export const settings: Settings = {
	API_URL: "https://epms-kw8f.onrender.com/api/v1",
	DEBUG: true,
}

export const debugLog = (...args: unknown[]): void => {
	if (settings.DEBUG) {
		console.log(...args);
	}
}
