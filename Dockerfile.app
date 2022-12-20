FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o apoor-ssh ./cmd/server
CMD [ "./apoor-ssh", "-host", "0.0.0.0", "-port", "22" ]
