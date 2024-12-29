import tailwindConfig from "@oauth-service-go/tailwind-config";
import type { Config } from "tailwindcss";
import resolveConfig from "tailwindcss/resolveConfig";

export function getTailwindTheme() {
  const { theme } = resolveConfig(tailwindConfig as Config);

  return theme;
}
