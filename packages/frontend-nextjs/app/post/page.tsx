"use client";

import Image from "next/image";

import Link from "next/link";
import { MdOutlineKeyboardArrowLeft } from "react-icons/md";

import { PostForm } from "../components/PostForm";

export default function PostPage() {
  return (
    <>
      <div className="absolute top-20 left-20 max-sm:invisible">
        <Link href={"/home"}>
          <MdOutlineKeyboardArrowLeft size={50} />
        </Link>
      </div>
      <div className="flex flex-col items-center center space-y-12">
        <Image
          height={200}
          width={200}
          src={"logo_postlogs.svg"}
          alt="Logo"
          priority
        ></Image>

        <PostForm post={null} />
      </div>
    </>
  );
}
