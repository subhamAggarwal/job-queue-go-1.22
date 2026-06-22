import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
// Preview-proxy awareness
const PREVIEW_BASE = process.env.PREVIEW_BASE_5173 || '/';
const API_PREFIX_IN_CLIENT_URL = `${PREVIEW_BASE.replace(/\/$/, '')}/api`;
const API_PROXY_KEY = `^${API_PREFIX_IN_CLIENT_URL}(/|$)`;

export default defineConfig({
  plugins: [react()],
  base: PREVIEW_BASE,
  server: {
    port: 5173,
    host: '0.0.0.0',
    strictPort: true,
    proxy: {
      [API_PROXY_KEY]: {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(new RegExp(API_PROXY_KEY), '/api/')
      }
    }
  }
})
