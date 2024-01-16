FROM golang:latest

WORKDIR /server
COPY . .

RUN go install github.com/cosmtrek/air@latest 

CMD ["air", "-c", ".air.toml"]