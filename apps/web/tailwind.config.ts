// tailwind config is required for editor support

import type { Config } from "tailwindcss"
import sharedConfig from "@oauth-service-go/tailwind-config"

const config: Pick<Config, "content" | "presets"> = {
  content: ["./src/**/*.tsx"],
  presets: [sharedConfig],
}

export default config
