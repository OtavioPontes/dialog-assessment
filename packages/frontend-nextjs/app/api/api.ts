"use server";

import { jwtDecode } from "jwt-decode";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function Register(dto: RegisterRequestDTO) {
  const res = await fetch(`${process.env.API_URL}/api/users`, {
    method: "POST",

    body: JSON.stringify({
      email: dto.Email,
      password: dto.Password,
      name: dto.Name,
      nick: dto.Nick,
    }),

    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }
}

export async function CreatePost(dto: RegisterPostDTO) {
  const token = cookies().get("session")!.value;

  const res = await fetch(`${process.env.API_URL}/api/posts`, {
    method: "POST",

    body: JSON.stringify({
      title: dto.title,
      content: dto.content,
    }),

    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }
}
export async function UpdatePost(dto: RegisterPostDTO, id: string) {
  const token = cookies().get("session")!.value;

  const res = await fetch(`${process.env.API_URL}/api/posts/${id}`, {
    method: "PUT",

    body: JSON.stringify({
      title: dto.title,
      content: dto.content,
    }),

    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }
}

export async function UpdateUser(dto: RegisterRequestDTO) {
  const token = cookies().get("session")!.value;
  const jwt: any = jwtDecode(token)!;

  const res = await fetch(`${process.env.API_URL}/api/users/${jwt["userId"]}`, {
    method: "PUT",

    body: JSON.stringify({
      name: dto.Name,
      nick: dto.Nick,
      email: dto.Email,
    }),

    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }
}

export async function Logout() {
  cookies().delete("session");
  redirect(`/`);
}

export async function Login(dto: AuthRequestDTO) {
  const res = await fetch(`${process.env.API_URL}/api/login`, {
    method: "POST",

    body: JSON.stringify({
      email: dto.Email,
      password: dto.Password,
    }),

    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }

  const json = await res.json();

  cookies().set("session", json["token"]);
  redirect(`/home`);
}

export async function GetUser(): Promise<User> {
  const token = cookies().get("session")!.value;
  const jwt: any = jwtDecode(token)!;

  const res = await fetch(`${process.env.API_URL}/api/users/${jwt["userId"]}`, {
    method: "GET",

    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
      "Access-Control-Allow-Origin": "*",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }

  return res.json();
}

export async function DeleteUser() {
  const token = cookies().get("session")!.value;
  const jwt: any = jwtDecode(token)!;

  const res = await fetch(`${process.env.API_URL}/api/users/${jwt["userId"]}`, {
    method: "DELETE",

    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
      "Access-Control-Allow-Origin": "*",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }

  cookies().delete("session");
  redirect(`/`);
}

export async function GetPost(id: string): Promise<Post> {
  const token = cookies().get("session")!.value;

  const res = await fetch(`${process.env.API_URL}/api/posts/${id}`, {
    method: "GET",

    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
      "Access-Control-Allow-Origin": "*",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }

  return res.json();
}

export async function GetPosts(): Promise<Post[]> {
  const token = cookies().get("session")!.value;

  const res = await fetch(`${process.env.API_URL}/api/posts`, {
    method: "GET",

    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
      "Access-Control-Allow-Origin": "*",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }

  return res.json();
}

export async function LikePost(id: string) {
  const token = cookies().get("session")!.value;
  const res = await fetch(`${process.env.API_URL}/api/posts/${id}/like`, {
    method: "POST",

    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    console.log(res);
    throw new Error(await res.text());
  }
}
