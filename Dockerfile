# Шаг 1: Установка зависимостей
FROM golang:1.18-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Сборка проекта, добавление внешних шлюзов, инициализация ФЛК
FROM golang:1.18-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/main

# Step 3: Запуск ядра валидатора, запуск внешних шлюзов
FROM scratch
EXPOSE 9011
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
CMD ["mkdir --parents ~/.mongodb && wget 'https://storage.yandexcloud.net/cloud-certs/CA.pem' --output-document ~/.mongodb/root.crt && chmod 0644 ~/.mongodb/root.crt"]
CMD ["/app"]
