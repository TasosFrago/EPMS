export interface CookieOptions {
	maxAge?: number;
	expires?: Date;
	path?: string;
	secure?: boolean;
	sameSite?: 'Strict' | 'Lax' | 'None';
}

export const CookieManager = {
	/**
	 * Set a cookie with specified options.
	 * @param {string} name - The name of the cookie.
	 * @param {string} value - The value of the cookie.
	 * @param {Object} options - Additional options for the cookie.
	 * @param {number} [options.maxAge] - Cookie's max age in seconds.
	 * @param {Date} [options.expires] - Expiry date of the cookie.
	 * @param {string} [options.path] - Path for the cookie (default is `/`).
	 * @param {boolean} [options.secure] - Secure flag (default is false).
	 * @param {string} [options.sameSite] - SameSite policy (default is `Lax`).
	 */
	set: (name: string, value: string, options: CookieOptions = {}) => {
		let cookieString = `${encodeURIComponent(name)}=${encodeURIComponent(value)};`;

		if (options.maxAge) cookieString += ` Max-Age=${options.maxAge};`;
		if (options.expires) cookieString += ` Expires=${options.expires.toUTCString()};`;
		cookieString += ` Path=${options.path || '/'};`;
		if (options.secure) cookieString += ` Secure;`;
		cookieString += ` SameSite=${options.sameSite || 'Lax'};`;

		document.cookie = cookieString;
	},

	/**
	 * Get the value of a specific cookie.
	 * @param {string} name - The name of the cookie to retrieve.
	 * @returns {string|null} - The value of the cookie or `null` if not found.
	 */
	get: (name: string): string | null => {
		const cookies = document.cookie.split('; ').reduce((acc: Record<string, string>, cookie) => {
			const [key, value] = cookie.split('=');
			acc[decodeURIComponent(key)] = decodeURIComponent(value);
			return acc;
		}, {});

		return cookies[name] || null;
	},

	/**
	 * Delete a specific cookie by setting its expiry date to the past.
	 * @param name - The name of the cookie to delete.
	 * @param options - Additional optionsl for the cookie path and domain.
	 */
	delete: (name: string, options: Pick<CookieOptions, 'path'> = {}) => {
		CookieManager.set(name, '', {
			path: options.path || '/',
			expires: new Date(0),
		});
	},
}
