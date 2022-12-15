import { writable, derived } from 'svelte/store';

export const currentPath = writable()
export const apiData = writable({});

export const aboutResponse = derived(apiData, ($apiData) => {
    console.log($apiData)
    if ($apiData["about"]) {
        return $apiData["about"]
    }

    return {}
});