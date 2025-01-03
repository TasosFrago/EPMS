
type Maybe<T> = T | null;

export type Result<R, E> = [Maybe<R>, Maybe<E>];

export enum ErrorHandler {
	REDIRECT = 1,
	INTERNAL_SERVER_ERROR,
}
