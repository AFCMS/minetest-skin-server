import type {Config} from "tailwindcss";

// noinspection JSUnusedGlobalSymbols
export default {
    content: ["./src/**/*.{js,jsx,ts,tsx}"],
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/forms"), require("@headlessui/tailwindcss")],
} satisfies Config;

