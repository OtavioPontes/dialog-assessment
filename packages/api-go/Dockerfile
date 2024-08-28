FROM golang:alpine

WORKDIR /api

COPY . .  

RUN go mod tidy

EXPOSE 5000

CMD ["go","run","main.go"]