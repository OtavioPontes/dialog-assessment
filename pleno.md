# Desafio nível Pleno
Neste nível do desafio, foi escolhido a stack com api em **Golang**, frontend em **Nextjs** e banco de dados **Postgres**.

## Como rodar o projeto:
Foi usado o docker, em especial o docker-compose para gerar as imagens e os containeres do projeto. O arquivo docker-compose.yml se encontra na raiz do projeto.

![image](https://github.com/user-attachments/assets/437b2bdf-7be8-47c9-95f9-5f200950e9dd)

### Observações
É preciso adicionar um arquivo .env com as credenciais para os projetos da **api-go** e **frontend-nextjs**. Como foi disponibilizado um .example.env para cada projeto, basta copiá-lo, renomeando-o para .env

![image](https://github.com/user-attachments/assets/e2b20eee-89b9-48e6-8d2f-6ac73c6ce361)


Para rodar basta rodar:
```
docker-compose up -d --build
```

Se ocorrer tudo conforme o esperado, os seguintes containeres devem estar rodando:

![image](https://github.com/user-attachments/assets/2fd322f7-c125-49f7-9ecd-9800c3646c1f)


O build demora um pouco, muito em função do build do Nextjs. Contudo, após finalizado, o banco de dados rodará sozinho o script para criação das tabelas. 
Ademais, se mantido o que está no .example.env, a url do frontend-nextjs é:
```
http://localhost:3000
```


## API (GOLANG)
A api foi feita usando o biblioteca http padrão e mux, devido a sua robustez e grande compatibilidade com outras bibliotecas e projetos.



![image](https://github.com/user-attachments/assets/25f2d266-0403-4b4a-9c07-1ac1f91459c5)


O projeto foi estruturado de forma bem simples e assertiva, visando a diminuição de boilerplates aliado a um certo nível de escalabilidade. 

```O fluxo de chamada é Router -> Controller -> Repository```

### Models da aplicação
![image](https://github.com/user-attachments/assets/3df2d005-63f4-404b-a25c-f8dbdc3880bb)  ![image](https://github.com/user-attachments/assets/5fa44f4d-b982-46fa-95a2-b69ed6a6341a)

Em ambos os casos foi escolhido o uuid com identificador por ser um padrão consolidado da indústria. Ao criar uma dessas entidades o próprio banco de dados gera o id.

### Segurança e Hash
![image](https://github.com/user-attachments/assets/a5303783-cc16-4e20-930b-89710be6d6c2)

Neste caso, por praticidade, foi utilizada a biblioteca bcrypt para hashear e comparar as senhas do usuário, tanto no momento de criação do usuário quanto na autenticação.

### Repositories
![image](https://github.com/user-attachments/assets/00b1ac8b-c3bc-48f7-81d2-cc7f58141c6b)

Os repositórios são, em grande parte, a camada de comunicação da aplicação com o banco de dados. No exemplo da imagem acima, foi utilizado também o redis como cache temporário.

### Autenticação
![image](https://github.com/user-attachments/assets/b56c92c0-3db2-435f-9a7b-8d1109ae46d2)

Conforme requisitado, o fluxo de autenticação foi feito todo em cima do JWT. No caso acima, foi criado o token com as informações de expiração, em 6 horas, e o id do usuário.

### Testes unitários
![image](https://github.com/user-attachments/assets/756f3783-ada0-42dc-81a3-30ccbd79aada)

Foram criados testes unitários bem simples na camada de repository da aplicação, que de certa forma é o core da aplicação. Poderiam ser feitos outros testes na camada de controller e em funções auxiliares, como as de hash.


### Swagger
A documentação da api foi feita usando o Swagger e está disponível na seguinte url:
```
http://localhost:5000/swagger/index.html
```
