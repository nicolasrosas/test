package main

import (
	"log"

	"encoding/json"

	"github.com/streadway/amqp"
)

type PaqueteProcesado struct{

	IdPaquete string 
	Tipo string 
	Valor int32 
	Intentos int32
    Timestamp string
	IdCamion string
	Estado string
	Seguimiento string

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://user:pass@10.6.40.74:5672/")
	failOnError(err, "Conexion a RabbitMQ fallida")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Apertura de canal fallida")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"finanzas", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Declaracion de cola fallida")

	paquete := PaqueteProcesado{
		IdPaquete: "1", 
		Tipo:"normal", 
		Valor: 2000, 
		Intentos: 3, 
		Timestamp: "0", 
		IdCamion: "12", 
		Estado:"Recibido", 
		Seguimiento:"132321",
	}

	body, err := json.Marshal(paquete)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Fallo al publicar un mensaje")
}