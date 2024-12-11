FROM golang:1.23

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Instala las dependencias necesarias (mod y swag)
RUN go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest

# Genera la documentación Swagger
RUN swag init --output ./docs

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el contenedor
EXPOSE 8081

# Comando por defecto para ejecutar el servidor
CMD ["./main"]
