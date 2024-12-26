import { jwtVerify } from "jose";
import { NextResponse } from "next/server";

export async function middleware(req: Request) {
  const secret = process.env.JWT_SECRET;
  const loginUrl = new URL("/login", req.url);

  if (!secret) {
    return NextResponse.json(
      { error: "JWT_SECRET not configured" },
      { status: 500 },
    );
  }

  const cookieHeader = req.headers.get("cookie");
  const token = cookieHeader
    ?.split("; ")
    ?.find((cookie) => cookie.startsWith("token="))
    ?.split("=")[1];

  if (!token) {
    return NextResponse.redirect(loginUrl);
  }

  try {
    const { payload } = await jwtVerify(
      token,
      new TextEncoder().encode(secret),
    );
    req.headers.set("x-client-id", payload.clientID as string);
    return NextResponse.next();
  } catch (err) {
    return NextResponse.redirect(loginUrl);
  }
}

export const config = {
  matcher: ["/dashboard/:path*", "/settings/:path*"],
};
