"use server";

import { cookies } from "next/headers";

export async function createTestCookie() {
  const cookieStore = await cookies();
  cookieStore.set("test_cookie", "test", {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV !== "development",
    maxAge: 36000,
  });

  console.log("Create Cookie: ", cookieStore.get("test_cookie"));

  return "Cookie created";
}
