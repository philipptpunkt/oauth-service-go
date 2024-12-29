import Link from "next/link";
import DeleteAccountButton from "./DeleteAccountButton";
import { LogoutClientButton } from "./LogoutClientButton";

export default async function SettingsPage() {
  const data = await fetch(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/v1/client/profile`,
  );

  const profileData = await data.json();

  console.log(profileData);

  return (
    <div className="p-4">
      <h1>Settings</h1>
      <p className="mt-2 text-gray-600">
        Manage your account settings here. More options coming soon!
      </p>
      <div>
        <h2>Personal Data</h2>
        {/* List personal data of client here */}
        <div>
          <Link href="/settings/profile">Edit personal data</Link>
        </div>
      </div>
      <div className="mt-6">
        <LogoutClientButton />
        <div className="h-4" />
        <DeleteAccountButton />
      </div>
    </div>
  );
}
