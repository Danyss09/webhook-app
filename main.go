package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "webhook-app/docs" // Importa la documentación generada

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Webhook API
// @version 1.0
// @description Ejemplo simple de un servidor de webhook con documentación Swagger.
// @host localhost:8081
// @BasePath /

// webhookHandler gestiona las solicitudes a /webhook
// @Summary Recibir datos en el webhook
// @Description Endpoint para recibir datos enviados mediante POST.
// @Tags Webhook
// @Accept json
// @Produce plain
// @Param data body string true "Datos del Webhook"
// @Success 200 {string} string "Webhook recibido exitosamente"
// @Failure 405 {string} string "Método no permitido"
// @Router /webhook [post]
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Datos recibidos en el webhook: %s\n", body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook recibido exitosamente"))
}

// defaultHandler gestiona rutas no definidas
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Ruta no encontrada", http.StatusNotFound)
}

func main() {
	// Configurar rutas
	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/", defaultHandler)

	// Configurar Swagger en /swagger/
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Iniciar servidor
	port := ":8081"
	fmt.Printf("Servidor de webhook en ejecución en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
