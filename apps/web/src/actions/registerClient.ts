"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function registerClient({
  email,
  password,
}: {
  email: string;
  password: string;
}) {
  const response = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/api/v1/client/register",
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

  const { token } = await response.json();

  const cookieStore = await cookies();

  cookieStore.set("temp_token", token, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production", // Use secure cookies in production
    maxAge: 3600, // 1 hour
  });

  redirect("/register/verify");
}
