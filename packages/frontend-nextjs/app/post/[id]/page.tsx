"use client";

import Image from "next/image";

import Link from "next/link";
import { MdOutlineKeyboardArrowLeft } from "react-icons/md";
import { PostForm } from "../../components/PostForm";
import { useEffect, useState } from "react";
import { GetPost, GetPosts } from "../../api/api";
import { InfinitySpin } from "react-loader-spinner";

export default function EditPostPage({ params }: { params: { id: string } }) {
  const [post, setPost] = useState<Post | undefined>();
  const [isLoading, setIsLoading] = useState(false);
  useEffect(() => {
    async function getPosts() {
      setIsLoading(true);
      const data = await GetPost(params.id);
      setPost(data);
      setIsLoading(false);
    }
    getPosts();
  }, [params.id]);
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
          src={"../logo_postlogs.svg"}
          alt="Logo"
          priority
        ></Image>

        {isLoading ? (
          <div className="flex justify-center mt-20">
            <InfinitySpin width="300" color="#000000" />
          </div>
        ) : (
          <PostForm post={post!} />
        )}
      </div>
    </>
  );
}
