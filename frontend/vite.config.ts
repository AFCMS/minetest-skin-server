import {defineConfig} from "vite"
import react from "@vitejs/plugin-react"

// https://vitejs.dev/config/
// noinspection JSUnusedGlobalSymbols
export default defineConfig({
    build: {
        manifest: true,
        modulePreload: {
            polyfill: false
        }
    },
    server: {
        host: "0.0.0.0",
        port: 5173,
        strictPort: true,
    },
    plugins: [react()],
    assetsInclude: ["**/*.gltf"],
})
