FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o object_storage ./cmd/api

RUN chmod +x /app/object_storage

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/object_storage ./
COPY store/cover/knight.jpeg ./store/cover/knight.jpeg

EXPOSE 80

CMD ["./object_storage"]
