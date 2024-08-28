FROM node:21-alpine

WORKDIR /app

COPY package.json ./

RUN npm install 

COPY . .

RUN npm install next && \ 
    npm run build

CMD ["npm", "run", "start"]