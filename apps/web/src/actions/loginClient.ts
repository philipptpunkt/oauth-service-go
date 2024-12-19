"use server"

import { redirect } from "next/navigation"
// import { cookies } from "next/headers"

export async function loginClient({
  email,
  password,
}: {
  email: string
  password: string
}) {
  console.log("DATA:", email, password)

  const response = await fetch(
    process.env.NEXT_PUBLIC_BACKEND_URL + "/api/v1/client/login",
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    }
  )

  if (!response.ok) {
    const error = await response.text()
    return { error: true, message: error }
  }

  redirect("/")
}
