"use server";

import { cookies } from "next/headers";

export async function testCookie() {
  const cookieStore = await cookies();
  const cookieTemp = cookieStore.get("test_cookie");

  console.log("Cookie server action: ", cookieTemp);

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/v1/cookie-test`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    },
  );

  const message = await response.text();
  console.log("Request message: ", message);
  return { message };
}
