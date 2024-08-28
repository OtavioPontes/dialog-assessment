"use client";

import { FormEvent, useEffect, useState } from "react";
import { Bounce, toast } from "react-toastify";
import { CreatePost, UpdatePost } from "../api/route";
import { useRouter } from "next/navigation";

export function PostForm({ post }: { post: Post | null }) {
  useEffect(() => {
    setIsLoading(true);
    if (post != null) {
      setTitle(post?.title);
      setContent(post?.content);
    }

    setIsLoading(false);
  }, [post]);

  const router = useRouter();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setIsLoading(true);
    setError(null);

    try {
      const formData = new FormData(event.currentTarget);
      const requestBody: RegisterPostDTO = {
        content: formData.get("content")?.toString() ?? "",
        title: formData.get("title")?.toString() ?? "",
      };
      event.currentTarget.reset();
      if (post != null) {
        await UpdatePost(requestBody, post.id);
      } else {
        await CreatePost(requestBody);
      }

      toast.success("Post criado com sucesso!", {
        onClose: () => {
          router.back();
        },
        position: "top-right",
        autoClose: 1000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        transition: Bounce,
      });
    } catch (error) {
      toast.error("Criação do Post falhou, tente novamente", {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        transition: Bounce,
      });

      console.error(error);
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <>
      <form onSubmit={onSubmit} className="w-full">
        <div className=" bg-white max-sm:p-4  p-12 rounded-lg flex flex-col items-center mt-4 shadow-lg">
          <div className="space-y-4 w-full text-center">
            <div className="space-y-2 tracking-widest">
              <h2 className="text-2xl">
                {post != null ? "Editar" : "Adicionar"}
              </h2>
              <h2 className="text-purple-500 text-3xl">Post</h2>
            </div>
          </div>
          <div className="w-1/2 max-sm:w-2/3 space-y-12 mt-16">
            <div className="space-y-4 w-full">
              <h4 className="tracking-widest">
                Entre com o título da postagem
              </h4>
              <input
                required
                type="text"
                id="title"
                name="title"
                className="bg-gray-100 w-full h-12 rounded-lg px-4"
                value={title}
                onChange={(event) => setTitle(event.target.value)}
              />
            </div>
          </div>
          <div className="w-1/2 max-sm:w-2/3 space-y-12 mt-16">
            <div className="space-y-4 w-full">
              <h4 className="tracking-widest">
                Entre com o conteúdo da postagem
              </h4>
              <input
                required
                type="text"
                id="content"
                name="content"
                className="bg-gray-100 w-full h-12 rounded-lg px-4"
                value={content}
                onChange={(event) => setContent(event.target.value)}
              />
            </div>
          </div>

          <button
            disabled={isLoading}
            type="submit"
            className={`mt-16 px-12 py-4 tracking-wider bg-purple-500 text-white text-xl  border-2 rounded-lg self-center  font-semibold`}
          >
            {isLoading ? "Enviando..." : "Enviar"}
          </button>
        </div>
      </form>
    </>
  );
}
