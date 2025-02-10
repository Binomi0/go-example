# Imagen base oficial de Go
FROM golang:1.21 AS builder

# Directorio donde trabajaremos dentro del contenedor
WORKDIR /go/src/mygoapp

# Copiamos el código fuente al contenedor
COPY . .

# Instalamos las dependencias
RUN go mod tidy && go build -o mygoapp

# Imagen final de la aplicación
FROM scratch as runtime

# Copiamos solo el ejecutable en la imagen final
WORKDIR /

COPY --from=builder /go/src/mygoapp/mygoapp ./

# Configuramos los permisos para que el ejecutable pueda ejecutarse
RUN chmod +x mygoapp

# Ejecutamos el programa
CMD ["./mygoapp"]
