"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useActionState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { registerClient } from "../../actions/registerClient";
import { registerClientWithFormData } from "../../actions/registerClientWithFormData";

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
});

type Inputs = z.infer<typeof RegisterClientSchema>;

export function RegisterClientForm() {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Inputs>({ resolver: zodResolver(RegisterClientSchema) });

  console.log("Client Form");

  return (
    <form onSubmit={handleSubmit(registerClient)}>
      <input
        className="mt-4 block w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        type="email"
        placeholder="email"
        {...register("email")}
      />
      {errors.email && (
        <span className="text-sm font-light text-red-600">
          {errors.email.message}
        </span>
      )}
      <input
        className="mt-4 block w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        type="password"
        placeholder="password"
        {...register("password")}
      />
      {errors.password && (
        <span className="text-sm font-light text-red-600">
          {errors.password.message}
        </span>
      )}

      <input
        type="submit"
        className="mt-4 w-full cursor-pointer rounded-md bg-black py-2 text-lg text-white"
        value={isSubmitting ? "loading..." : "Register"}
      />
    </form>
  );
}

export function RegisterClient() {
  const [state, submitAction, isPending] = useActionState(
    registerClientWithFormData,
    undefined,
  );

  const hasError = state && state.error;

  return (
    <form action={submitAction}>
      <input
        className="mt-4 block w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        name="email"
        type="email"
        placeholder="email"
        disabled={isPending}
      />
      {hasError && state?.email ? (
        <p className="text-sm font-light text-red-600">{state.email}</p>
      ) : null}

      <input
        className="mt-4 block w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
        name="password"
        type="password"
        placeholder="password"
        disabled={isPending}
      />
      {hasError && state?.password ? (
        <p className="text-sm font-light text-red-600">{state.password}</p>
      ) : null}

      <input
        type="submit"
        className="mt-4 w-full cursor-pointer rounded-md bg-black py-2 text-lg text-white"
        value={isPending ? "loading..." : "Register"}
        disabled={isPending}
      />
    </form>
  );
}
