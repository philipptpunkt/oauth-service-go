"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { clearClientCookies } from "./clearClientCookies";

export async function verifyEmailCode(data: { code: string }) {
  const cookieStore = await cookies();
  const tempToken = cookieStore.get("temp_token");

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/v1/client/verify-email-code`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${tempToken?.value}`,
      },
      body: JSON.stringify(data),
    },
  );

  if (!response.ok) {
    const error = await response.text();
    console.log(error);
    throw new Error(error);
    // return { error: true, message: error };
  }

  const { token, refresh_token, profile_completed } = await response.json();

  await clearClientCookies(["temp_token"]);

  cookieStore.set("refresh_token", refresh_token, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "strict",
    maxAge: 7 * 24 * 60 * 60, // 7 days
  });

  cookieStore.set("token", token, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "strict",
    maxAge: 15 * 60, // 15 minutes
  });

  if (profile_completed) {
    redirect("/dashboard");
  }

  redirect("/register/complete-profile");
}
