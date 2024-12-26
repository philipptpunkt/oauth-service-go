"use client";

import { useState } from "react";
import { deleteUser } from "../../actions/deleteClient";

export default function DeleteAccountButton() {
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    setIsDialogOpen(false);
    setIsDeleting(true);
    try {
      await deleteUser();
    } catch (error) {
      console.error("Failed to delete user:", error);
      alert("Something went wrong. Please try again.");
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <div>
      <button
        className="rounded bg-red-500 px-4 py-2 text-white"
        onClick={() => setIsDialogOpen(true)}
        disabled={isDeleting}
      >
        {isDeleting ? "Deleting..." : "Delete Account"}
      </button>

      {isDialogOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="rounded bg-white p-4 shadow-md">
            <h2 className="text-lg font-bold">Are you sure?</h2>
            <p>This action cannot be undone.</p>
            <div className="mt-4 flex justify-end">
              <button
                className="mr-2 rounded bg-gray-300 px-4 py-2"
                onClick={() => setIsDialogOpen(false)}
              >
                Cancel
              </button>
              <button
                className="rounded bg-red-500 px-4 py-2 text-white"
                onClick={handleDelete}
              >
                Confirm
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
