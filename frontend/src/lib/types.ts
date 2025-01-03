export class EmptyCookie extends Error {

	constructor(message: string) {
		super(message);
		this.name = "Empty Cookie";

		Object.setPrototypeOf(this, EmptyCookie.prototype);
	}
}

export class UnauthorizedUserError extends Error {
	public code: number;

	constructor(message: string, code: number) {
		super(message);
		this.name = "Unauthorized User";
		this.code = code;

		Object.setPrototypeOf(this, UnauthorizedUserError.prototype);
	}
}

export class InternalServerError extends Error {
	public code: number;

	constructor(message: string, code: number) {
		super(message);
		this.name = "Internal Server Error";
		this.code = code;

		Object.setPrototypeOf(this, InternalServerError.prototype);
	}
}

export enum PopupStatus {
	INFO = '#edede9',
	WARNING = '#ffd60a',
	ERROR = '#e63946',
	SUCCESS = '#28a745'
}
