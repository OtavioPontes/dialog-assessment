import Image from "next/image";

import { AuthForm } from "../components/AuthForm";

export default function AuthPage() {
  return (
    <>
      <div className="flex flex-col items-center center space-y-12">
        <Image
          height={300}
          width={300}
          src={"logo_postlogs.svg"}
          alt="Logo"
          priority
        ></Image>
        <AuthForm />
      </div>
    </>
  );
}
