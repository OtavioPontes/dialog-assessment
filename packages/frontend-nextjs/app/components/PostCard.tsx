import { FaCalendar, FaPen } from "react-icons/fa";
import { IoPersonCircleSharp } from "react-icons/io5";
import { FaHeart } from "react-icons/fa6";
import { LikePost } from "../api/route";
import Link from "next/link";

export default function PostCard({ post, user }: { post: Post; user: User }) {
  return (
    <div className=" bg-white px-12 py-8 rounded-lg flex flex-col mt-4 shadow-md  max-sm:justify-center max-sm:items-center">
      <div className="flex-row flex max-sm:flex-col items-center max-sm:justify-center justify-between max-sm:mb-2 mb-8 ">
        <h2 className="text-2xl">{post.title}</h2>
        <div className="max-sm:invisible">
          {post.authorId == user.id ? (
            <Link href={`/post/${post.id}`}>
              <FaPen size={20} />
            </Link>
          ) : null}
        </div>
      </div>

      <h3 className="text-lg mb-4 max-sm:mb-8">{post.content}</h3>
      <div className="flex max-sm:flex-col flex-row justify-end max-sm:space-x-0 max-sm:items-start max-sm:place-self-center space-x-12 max-sm:space-y-2 max-sm:mb-4">
        <div className="flex flex-row space-x-4 ml-4 max-sm:ml-0 items-center">
          <FaCalendar size={12} />
          <span className="text-sm text-end">
            {Intl.DateTimeFormat("pt-BR").format(post.createdat)}
          </span>
        </div>
        <div className="flex flex-row space-x-4 ml-4 max-sm:ml-0 items-center">
          <IoPersonCircleSharp size={15} />
          <span className="text-sm text-cyan-600 text-end">
            {post.authorNick}
          </span>
        </div>
      </div>
      <div className="mt-4 flex flex-row space-x-4">
        <button onClick={() => LikePost(post.id)}>
          <FaHeart className="text-red-600" />
        </button>
        <p>{post.likes}</p>
      </div>
      <div className="max-sm:visible invisible self-end">
        {post.authorId == user.id ? (
          <Link href={`/post/${post.id}`}>
            <FaPen size={20} />
          </Link>
        ) : null}
      </div>
    </div>
  );
}
