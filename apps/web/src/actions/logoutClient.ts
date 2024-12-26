"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { clearClientCookies } from "./clearClientCookies";

export async function logoutClient() {
  const cookieStore = await cookies();
  const token = cookieStore.get("token");

  if (!token) {
    throw new Error("Not authenticated");
  }

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/v1/client/logout`,
    {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    },
  );

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || "Failed to logout");
  }

  await clearClientCookies(["token", "refresh_token"]);

  redirect("/");
}
