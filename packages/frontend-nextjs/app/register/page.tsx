import Image from "next/image";

import { RegisterForm } from "../components/RegisterForm";
import Link from "next/link";
import { MdOutlineKeyboardArrowLeft } from "react-icons/md";

export default function AuthPage() {
  return (
    <>
      <div className="absolute top-20 left-20 max-sm:invisible">
        <Link href={"/"}>
          <MdOutlineKeyboardArrowLeft size={50} />
        </Link>
      </div>
      <div className="flex flex-col items-center center space-y-12">
        <Image
          height={300}
          width={300}
          src={"logo_postlogs.svg"}
          alt="Logo"
          priority
        ></Image>
        <RegisterForm />
      </div>
    </>
  );
}
