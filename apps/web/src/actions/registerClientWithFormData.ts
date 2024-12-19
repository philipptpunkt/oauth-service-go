"use server"

import { redirect } from "next/navigation"
import { z } from "zod"

export async function registerClientWithFormData(
  _: unknown,
  formData: FormData
) {
  const RegisterClientSchema = z.object({
    email: z
      .string()
      .trim()
      .nonempty({ message: "Email is required" })
      .email({ message: "Email format not valid" }),
    password: z
      .string()
      .trim()
      .nonempty({ message: "Password is required" })
      .min(8, { message: "Password must contain at least 8 characters" }),
  })

  const dataCheckResult = RegisterClientSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  })

  if (!dataCheckResult.success) {
    const errors = dataCheckResult.error.format()

    const emailError = errors.email?._errors[0]
    const passwordError = errors.password?._errors[0]

    return { error: true, email: emailError, password: passwordError }
  }

  // const response = await fetch(
  //   process.env.NEXT_PUBLIC_BACKEND_URL + "/api/v1/client/register",
  //   {
  //     method: "POST",
  //     headers: { "Content-Type": "application/json" },
  //     body: JSON.stringify(data),
  //   }
  // )

  // if (!response.ok) {
  //   const error = await response.text()
  //   throw new Error(error || "Registration failed")
  // }
  redirect("/")
}
