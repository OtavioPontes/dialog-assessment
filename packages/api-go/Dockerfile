FROM golang:latest AS builder

WORKDIR /usr/src/app

RUN go mod init api-go && go mod tidy

COPY . .
RUN go build -o api main.go

# FROM scratch
# WORKDIR /usr/src/app
# COPY --from=builder /usr/src/app .

CMD [ "./api" ]
