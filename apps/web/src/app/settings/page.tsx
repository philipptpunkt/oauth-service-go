import DeleteAccountButton from "./DeleteAccountButton";
import { LogoutClientButton } from "./LogoutClientButton";

export default function SettingsPage() {
  return (
    <div className="p-4">
      <h1 className="text-xl font-bold">Settings</h1>
      <p className="mt-2 text-gray-600">
        Manage your account settings here. More options coming soon!
      </p>
      <div className="mt-6">
        <LogoutClientButton />
        <div className="h-4" />
        <DeleteAccountButton />
      </div>
    </div>
  );
}
