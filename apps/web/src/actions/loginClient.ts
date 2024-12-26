"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function loginClient({
  email,
  password,
}: {
  email: string;
  password: string;
}) {
  const response = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/api/v1/client/login",
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    },
  );

  if (!response.ok) {
    const error = await response.text();
    return { error: true, message: error };
  }

  const { token, refresh_token } = await response.json();

  const cookieStore = await cookies();

  cookieStore.set("token", token, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "strict",
    maxAge: 15 * 60, // 15 minutes
  });

  cookieStore.set("refresh_token", refresh_token, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "strict",
    maxAge: 7 * 24 * 60 * 60, // 7 days
  });

  redirect("/dashboard");
}
