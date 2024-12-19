"use client"

import { useForm } from "react-hook-form"
import { registerClientWithFormData } from "../../actions/registerClientWithFormData"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useActionState } from "react"
import { registerClient } from "../../actions/registerClient"

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

type Inputs = z.infer<typeof RegisterClientSchema>

export function RegisterClientForm() {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Inputs>({ resolver: zodResolver(RegisterClientSchema) })

  return (
    <form onSubmit={handleSubmit(registerClient)}>
      <input
        className="mt-4 px-4 py-2 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        type="email"
        placeholder="email"
        {...register("email")}
      />
      {errors.email && (
        <span className="text-red-600 text-sm font-light">
          {errors.email.message}
        </span>
      )}
      <input
        className="mt-4 px-4 py-2 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        type="password"
        placeholder="password"
        {...register("password")}
      />
      {errors.password && (
        <span className="text-red-600 text-sm font-light">
          {errors.password.message}
        </span>
      )}

      <input
        type="submit"
        className="bg-black text-white py-2 w-full mt-4 rounded-md cursor-pointer text-lg"
        value={isSubmitting ? "loading..." : "Register"}
      />
    </form>
  )
}

export function RegisterClient() {
  const [state, submitAction, isPending] = useActionState(
    registerClientWithFormData,
    undefined
  )

  const hasError = state && state.error

  return (
    <form action={submitAction}>
      <input
        className="mt-4 px-4 py-2 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        name="email"
        type="email"
        placeholder="email"
        disabled={isPending}
      />
      {hasError && state?.email ? (
        <p className="text-red-600 text-sm font-light">{state.email}</p>
      ) : null}

      <input
        className="mt-4 px-4 py-2 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        name="password"
        type="password"
        placeholder="password"
        disabled={isPending}
      />
      {hasError && state?.password ? (
        <p className="text-red-600 text-sm font-light">{state.password}</p>
      ) : null}

      <input
        type="submit"
        className="bg-black text-white py-2 w-full mt-4 rounded-md cursor-pointer text-lg"
        value={isPending ? "loading..." : "Register"}
        disabled={isPending}
      />
    </form>
  )
}
