"use client";

import { useState } from "react";
import { testCookie } from "../../actions/cookieTest";
import { createTestCookie } from "../../actions/createTestCookie";

export function CookieTest() {
  const [status, setStatus] = useState<string>("");
  // const { 0: cookieState, 1: createCookie } = useActionState(
  //   createTestCookie,
  //   undefined,
  // );
  // const { 0: state, 1: sendCookie } = useActionState(testCookie, undefined);

  return (
    <div>
      <h4>Cookie Test</h4>
      <button
        className="rounded-md bg-blue-500 px-8 py-4"
        onClick={() => createTestCookie()}
      >
        Create Cookie
      </button>
      {/* <p>{cookieState?.message}</p> */}

      <button
        className="rounded-md bg-amber-500 px-8 py-4"
        onClick={async () => {
          const response = await testCookie();
          setStatus(response.message);
        }}
      >
        Check Cookie
      </button>
      <p>{status}</p>
      <button
        className="rounded-md bg-amber-500 px-8 py-4"
        onClick={async () => {
          const response = await fetch(
            `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/v1/cookie-test`,
            {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              credentials: "include",
            },
          );
          const message = await response.text();
          console.log("Request message: ", message);
          setStatus(message);
        }}
      >
        Check Cookie Directly
      </button>
    </div>
  );
}
