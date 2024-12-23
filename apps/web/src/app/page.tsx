// import { Card } from "@oauth-service-go/ui/Cards"
// import { redirect } from "next/navigation"
// import { cookies } from "next/headers"
// import { RegisterClientForm } from "../components/forms/RegisterClientForm"

import { CookieTest } from "../components/test/CookieTest";

export default function Page() {
  return (
    <main className="">
      <div className="flex h-screen items-center">
        <div className="w-1/2 p-24">
          <h1 className="">Welcome to Oauth2.0 Service</h1>
        </div>
        <div className="w-1/2">
          Some content
          <CookieTest />
        </div>
      </div>
      <div className="h-screen bg-gray-500">Content</div>
    </main>
  );
}
