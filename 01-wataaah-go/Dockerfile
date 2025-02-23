# Usa una imagen base ligera de Go
FROM golang:1.23-alpine AS builder

# Configuración del entorno
WORKDIR /app

# Copia los archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila la aplicación en modo release
RUN go build -o main .

# Usa una imagen final más pequeña para ejecutar la aplicación
FROM alpine:latest
WORKDIR /root/

# Copia el binario desde la fase de compilación
COPY --from=builder /app/main .

# Expone el puerto del servidor Gin
EXPOSE 8080

# Comando de ejecución
CMD ["./main"]