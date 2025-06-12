import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  base: '/daily-mood-tracker-app/', // 👈 required for GitHub Pages
  plugins: [react()],
});
