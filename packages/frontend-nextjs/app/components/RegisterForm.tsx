"use client";

import { FormEvent, useState } from "react";
import { Bounce, toast } from "react-toastify";
import { Login, Register } from "../api/route";
import { useRouter } from "next/navigation";

export function RegisterForm() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setIsLoading(true);
    setError(null);

    try {
      const formData = new FormData(event.currentTarget);
      const requestBody: RegisterRequestDTO = {
        Email: formData.get("email")?.toString() ?? "",
        Password: formData.get("password")?.toString() ?? "",
        Name: formData.get("name")?.toString() ?? "",
        Nick: formData.get("nick")?.toString() ?? "",
      };
      event.currentTarget.reset();
      await Register(requestBody).then(() => {
        toast.success("Cadastro deu certo!", {
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
      });
    } catch (error) {
      toast.error("Cadastro falhou, tente novamente", {
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
              <h2 className="text-2xl">Vamos iniciar seu</h2>
              <h2 className="text-purple-500 text-3xl">Cadastro</h2>
            </div>
          </div>
          <div className="w-1/2 max-sm:w-2/3 space-y-12 mt-16">
            <div className="space-y-4 w-full">
              <h4 className="tracking-widest">Entre com o email</h4>
              <input
                required
                type="email"
                id="email"
                name="email"
                className="bg-gray-100 w-full h-12 rounded-lg px-4"
              />
            </div>
          </div>
          <div className="w-1/2 max-sm:w-2/3 space-y-12 mt-16">
            <div className="space-y-4 w-full">
              <h4 className="tracking-widest">Entre com seu nome</h4>
              <input
                required
                type="text"
                id="name"
                name="name"
                className="bg-gray-100 w-full h-12 rounded-lg px-4"
              />
            </div>
          </div>
          <div className="w-1/2 max-sm:w-2/3 space-y-12 mt-16">
            <div className="space-y-4 w-full">
              <h4 className="tracking-widest">Entre com um nick</h4>
              <input
                required
                type="text"
                id="nick"
                name="nick"
                className="bg-gray-100 w-full h-12 rounded-lg px-4"
              />
            </div>
          </div>
          <div className="space-y-4 max-sm:w-2/3 w-1/2  mt-16">
            <h4 className="tracking-widest">Entre com a senha</h4>
            <input
              required
              type="password"
              id="password"
              name="password"
              className="bg-gray-100 w-full h-12 rounded-lg px-4"
            />
          </div>
          <button
            disabled={isLoading}
            type="submit"
            className={`mt-16 px-12 py-4 tracking-wider bg-purple-500 text-white text-xl  border-2 rounded-lg self-center  font-semibold`}
          >
            {isLoading ? "Enviando..." : "Cadastrar"}
          </button>
        </div>
      </form>
    </>
  );
}
