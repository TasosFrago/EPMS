import { type HttpError } from "@sveltejs/kit";


export interface CustomHttpError extends HttpError {
	shouldRedirect: {
		flag: boolean,
		path: string
	}
}

export class EmptyCookie implements CustomHttpError {
	status: number;
	body: App.Error;
	shouldRedirect: { flag: boolean, path: string }

	constructor(message: string, shouldRedirect: { flag: boolean, path: string }) {
		//super(message);
		this.status = 401
		this.body = {
			message: message
		}
		this.shouldRedirect = shouldRedirect
		//this.name = "Empty Cookie";

		Object.setPrototypeOf(this, EmptyCookie.prototype);
	}
}

export class UnauthorizedUserError implements CustomHttpError {
	public status: number;
	body: App.Error;
	shouldRedirect: { flag: boolean, path: string }

	constructor(message: string, shouldRedirect: { flag: boolean, path: string }) {
		//super(message);
		//this.name = "Unauthorized User";
		this.status = 401;
		this.body = {
			message: message
		}
		this.shouldRedirect = shouldRedirect

		Object.setPrototypeOf(this, UnauthorizedUserError.prototype);
	}
}

export class InternalServerError implements CustomHttpError {
	public status: number;
	body: App.Error
	shouldRedirect: { flag: boolean, path: string }

	constructor(message: string, shouldRedirect: { flag: boolean, path: string }) {
		//this.name = "Internal Server Error";
		this.status = 500;
		this.body = {
			message: message
		}
		this.shouldRedirect = shouldRedirect

		Object.setPrototypeOf(this, InternalServerError.prototype);
	}
}

export enum PopupStatus {
	INFO = '#edede9',
	WARNING = '#ffd60a',
	ERROR = '#e63946',
	SUCCESS = '#28a745'
}
