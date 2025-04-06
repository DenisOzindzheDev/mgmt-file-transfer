FROM golang:1.24.1 as builder
WORKDIR /build
COPY . . 
RUN go mod tidy
RUN go build .
FROM debian:latest
# Декларируем ARG, который можно передать при сборке
WORKDIR /app
ARG IMAGE_TAG
# Фиксируем его в ENV переменной
ENV SERVER_VERSION=$IMAGE_TAG
COPY --from=builder /build/mgmt-file-transfer /app/
CMD [ "/app/mgmt-file-transfer" ]