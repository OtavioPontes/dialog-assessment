"use client";

import Image from "next/image";
import { useEffect, useState } from "react";
import { GetPosts, GetUser, Logout } from "../api/api";
import { IoPersonCircleSharp } from "react-icons/io5";
import { FaRegNewspaper, FaCirclePlus } from "react-icons/fa6";
import Link from "next/link";
import PostCard from "../components/PostCard";
import { InfinitySpin } from "react-loader-spinner";
import { BiLogOut } from "react-icons/bi";

export default function HomePage() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [user, setUser] = useState<User>();
  useEffect(() => {
    async function getPosts() {
      setIsLoading(true);
      const data = await GetPosts();
      setPosts(data);
      const user = await GetUser();
      setUser(user);
      setIsLoading(false);
    }
    getPosts();
  }, []);

  return (
    <div className="mx-40 max-sm:mx-12">
      <button className="fixed bottom-10 right-10 bg-purple-500 px-6 py-4 rounded-lg shadow-lg text-white text-xl">
        <Link href={"/post/"}>
          <div className="flex flex-row space-x-4 items-center">
            <FaCirclePlus />
            <p>Adicionar Post</p>
          </div>
        </Link>
      </button>
      <div className="absolute mt-4 items-center right-20 max-sm:relative max-sm:right-0">
        <div className="flex flex-row space-x-12 items-start max-sm:justify-between justify-end">
          <div className="flex flex-col items-center justify-center">
            <Link href={"/profile"}>
              <IoPersonCircleSharp size={50} />
            </Link>
            <h5 className="mt-2 text-sm">Meu Perfil</h5>
          </div>
          <div className="flex mt-2 flex-col items-center justify-center">
            <button onClick={() => Logout()}>
              <BiLogOut size={40} />
            </button>
          </div>
        </div>
      </div>

      <div className="flex justify-center">
        <Image
          height={250}
          width={250}
          src={"logo_postlogs.svg"}
          alt="Logo"
          priority
        ></Image>
      </div>
      <div className="mt-12 flex flex-row space-x-8 items-center max-sm:justify-center">
        <FaRegNewspaper size={25} />
        <h2 className="text-2xl tracking-wider">Posts da sua rede</h2>
      </div>
      <div className="mt-20 space-y-12">
        {isLoading ? (
          <div className="flex justify-center mt-20">
            <InfinitySpin width="300" color="#000000" />
          </div>
        ) : posts?.length > 0 ? (
          posts?.map((e) => <PostCard key={e.id} post={e} user={user!} />)
        ) : (
          <div className="flex flex-col justify-center mt-24 space-y-5 text-center">
            <h3 className="text-2xl">Ainda não temos posts 😥</h3>
            <h3 className="text-lg text-gray-400">Publique um agora mesmo</h3>
          </div>
        )}
      </div>
    </div>
  );
}
