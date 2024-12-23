"use client";

import Link from "next/link";

export default function ErrorPage() {
  return (
    <div>
      <h1>Ohh Nooo!!!!</h1>
      <h2>Something went wrong</h2>
      <div className="flex">
        <Link href="/register">Try Again</Link>
        <Link href="/">Back to Home</Link>
      </div>
    </div>
  );
}
