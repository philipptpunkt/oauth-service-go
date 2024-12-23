import { VerifyEmailCodeForm } from "../../../components/forms/VerifyEmailCodeForm";

export default function EmailVerificationPage() {
  return (
    <div className="mx-auto mt-8 max-w-[500px]">
      <h1 className="text-2xl font-bold">Verify Your Email</h1>
      <VerifyEmailCodeForm />
    </div>
  );
}
