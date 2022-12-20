FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o apoor-ssh ./cmd/server

ENV APP_HOST=0.0.0.0
ENV APP_PORT=22

CMD [ "./apoor-ssh" ]
