import type { Config } from "tailwindcss"
import sharedConfig from "@oauth-service-go/tailwind-config"

const config: Pick<Config, "presets" | "content"> = {
  content: ["./src/**/*.tsx"],
  presets: [sharedConfig],
}

export default config
