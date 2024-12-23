"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

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

  redirect("/dashboard");
}
