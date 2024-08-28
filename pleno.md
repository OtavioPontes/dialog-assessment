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
A api foi feita usando o biblioteca http padrão, devido a sua robustez e grande compatibilidade com outras bibliotecas e projetos.



![image](https://github.com/user-attachments/assets/25f2d266-0403-4b4a-9c07-1ac1f91459c5)

O projeto foi estruturado de forma bem simples e assertiva, visando a diminuição de boilerplates aliado a um certo nível de escalabilidade. 

```O fluxo de chamada é Router -> Controller -> Repository```

### Swagger
A documentação da api foi feita usando o Swagger e está disponível na seguinte url:
```
http://localhost:5000/swagger/index.html
```
