import { writable } from 'svelte/store';
import { PopupStatus } from './types';

export interface PopupStoreT {
	show: boolean;
	msg: string;
	status: PopupStatus;
}

export const popupStore = writable<PopupStoreT>({
	show: false,
	msg: '',
	status: PopupStatus.INFO
});
