"use client";
import ToastProvider from "./components/ToastProvider";

export default function Template({ children }: { children: React.ReactNode }) {
  return <ToastProvider>{children}</ToastProvider>;
}
