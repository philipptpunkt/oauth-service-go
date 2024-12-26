"use server";

import { cookies } from "next/headers";

export async function clearClientCookies(cockieNames: string[]) {
  const cookieStore = await cookies();

  for (const name of cockieNames) {
    cookieStore.delete(name);
  }
}
