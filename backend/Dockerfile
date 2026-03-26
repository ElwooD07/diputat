FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ENV PORT=8080
ENV DATA_DIR=/data/samples

EXPOSE 8080

CMD ["go", "run", "./cmd/server"]
