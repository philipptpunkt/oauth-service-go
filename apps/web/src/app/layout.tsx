import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { NavigationBar } from "../components/navigation/NavigationBar";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "OAuth2.0 Service",
  description: "Experimental Service for OAuth2.0",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <NavigationBar />
        {children}
      </body>
    </html>
  );
}
