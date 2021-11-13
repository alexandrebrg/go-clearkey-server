FROM golang:1.17-alpine3.14

ENV ENV "production"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o ./clearkey-server

WORKDIR /dist
RUN cp /app/clearkey-server .

EXPOSE 8080

CMD [ "/dist/clearkey-server" ]