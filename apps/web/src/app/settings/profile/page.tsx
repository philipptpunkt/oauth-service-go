import { CompleteProfileForm } from "../../../components/forms/CompleteProfileForm";

export default function ProfilePage() {
  return (
    <div className="mx-auto mt-8 max-w-[500px]">
      <h1 className="text-2xl font-bold">Complete Your Profile</h1>
      <CompleteProfileForm />
    </div>
  );
}
