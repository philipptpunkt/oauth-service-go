import resolveConfig from "tailwindcss/resolveConfig"
import type { Config } from "tailwindcss"
import tailwindConfig from "@oauth-service-go/tailwind-config"

export function getTailwindTheme() {
  const { theme } = resolveConfig(tailwindConfig as Config)

  return theme
}
