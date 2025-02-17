# Imagen base oficial de Go
FROM golang:1.21 AS builder

# Directorio donde trabajaremos dentro del contenedor
WORKDIR /app

# Copiamos el código fuente al contenedor
COPY . .

# Instalamos las dependencias
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/mygoapp

# Imagen final de la aplicación
FROM alpine as runtime

# Copiamos solo el ejecutable en la imagen final
WORKDIR /

COPY --from=builder /app/mygoapp /mygoapp
COPY --from=builder /app/public /public
COPY --from=builder /app/templates /templates

# Configuramos los permisos para que el ejecutable pueda ejecutarse
RUN chmod +x /mygoapp

# Ejecutamos el programa
CMD ["/mygoapp"]
