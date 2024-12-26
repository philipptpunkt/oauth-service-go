"use client";

import { logoutClient } from "../../actions/logoutClient";

export function LogoutClientButton() {
  const handleLogout = async () => {
    try {
      await logoutClient();
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };

  return (
    <button
      onClick={handleLogout}
      className="rounded-md border border-red-500 px-4 py-2 text-red-500"
    >
      Logout
    </button>
  );
}
