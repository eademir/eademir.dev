import { writable } from 'svelte/store';


export const result = writable()
export const statusCode = writable()

export const activePage = writable(0)
export const blogURL = writable("")

export const isVisible = writable(false)
export const isPopUpVisable = writable(false)
export const showToast = writable(false)
export const message = writable("")