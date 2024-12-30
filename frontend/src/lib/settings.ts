
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
	API_URL: "http://localhost:8080/api/v1",
	DEBUG: true,
}

export const debugLog = (text: string): void => {
	console.log(text);
}
