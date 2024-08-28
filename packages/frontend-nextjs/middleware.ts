import { jwtDecode } from "jwt-decode";
import { cookies } from "next/headers";
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export const protectedRoutes = ["/profile", "/home", "/post"];
export const authRoutes = ["/"];
export const publicRoutes = ["/", "/register"];

export default async function middleware(request: NextRequest) {
  const currentUser = cookies().get("session")?.value;

  const isInvalid =
    !currentUser || Date.now() > (jwtDecode(currentUser).exp ?? Date.now());

  if (authRoutes.includes(request.nextUrl.pathname) && !isInvalid) {
    return NextResponse.redirect(new URL("/home", request.url));
  }

  if (protectedRoutes.includes(request.nextUrl.pathname) && isInvalid) {
    request.cookies.delete("session");
    const response = NextResponse.redirect(new URL("/", request.url));
    response.cookies.delete("session");

    return response;
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/:path*", "/post/:path*"],
};
