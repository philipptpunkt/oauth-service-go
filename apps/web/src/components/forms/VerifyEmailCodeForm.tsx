"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { verifyEmailCode } from "../../actions/verifyEmailCode";

const RegisterClientSchema = z.object({
  code: z
    .string()
    .trim()
    .length(6, { message: "Code must be exactly 6 digits" })
    .regex(/^\d+$/, { message: "Code must only contain numbers" }),
});

type Inputs = z.infer<typeof RegisterClientSchema>;

export function VerifyEmailCodeForm() {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<Inputs>({ resolver: zodResolver(RegisterClientSchema) });

  return (
    <form onSubmit={handleSubmit(verifyEmailCode)}>
      <input
        className="mt-2 w-full rounded border p-2"
        type="text"
        placeholder="Enter the code sent to your email"
        {...register("code")}
      />
      {errors.code && (
        <span className="text-sm font-light text-red-600">
          {errors.code.message}
        </span>
      )}
      <input
        type="submit"
        className="mt-4 w-full cursor-pointer rounded-md bg-black py-2 text-lg text-white"
        value={isSubmitting ? "loading..." : "Verify"}
      />
    </form>
  );
}
