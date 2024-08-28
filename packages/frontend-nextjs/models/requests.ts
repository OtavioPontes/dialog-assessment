interface AuthRequestDTO {
  Email: string;
  Password: string;
}

interface RegisterRequestDTO {
  Email: string;
  Name: string;
  Nick: string;
  Password: string;
}

interface RegisterPostDTO {
  title: string;
  content: string;
}
