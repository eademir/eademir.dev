import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import sveltePreprocess from 'svelte-preprocess';
import mkcert from 'vite-plugin-mkcert'

export default defineConfig({
  server: {
    https: true
  },
  plugins: [
    svelte({
      preprocess: [sveltePreprocess({ postcss: true })],
    }),
    mkcert()
  ],
})