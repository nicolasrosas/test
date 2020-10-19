package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"fmt"

	"encoding/json"
	"encoding/csv"

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


var (
	output_file = "finanzas.csv"
	connection = "amqp://user:pass@10.6.40.74:5672/"
	total = 0.0
	perdida = 0.0
	ganancia = 0.0
)


func ActualizarInfo(paquete PaqueteProcesado) ([]string){

	var ganancia_paquete = 0.0
	var perdida_paquete = 0.0
	// Ganancias por paquetes entregados o no entregados
	if paquete.Estado == "Recibido"{

		ganancia_paquete += float64(paquete.Valor)  

	}else if paquete.Estado == "No Recibido"{		//Paquete no recibido

		if paquete.Tipo == "prioritario" {
			ganancia_paquete += float64(0.3) * float64(paquete.Valor)
		} else if paquete.Tipo == "retail" && paquete.Intentos == 3 {
			ganancia_paquete += float64(paquete.Valor) 
		}else{
			ganancia_paquete += 0
		}
	}

	// Perdidas por reintento de envio

	perdida_paquete = float64(10) * float64(paquete.Intentos - 1)

	// Total de cada paquete
	total_paquete := ganancia_paquete - perdida_paquete
	
	total += ganancia_paquete - perdida_paquete
	ganancia += ganancia_paquete
	perdida += perdida_paquete

	data := []string{paquete.IdPaquete, paquete.Tipo, paquete.Estado, strconv.Itoa(int(paquete.Intentos)), fmt.Sprintf("%f",ganancia_paquete), fmt.Sprintf("%f",perdida_paquete), fmt.Sprintf("%f",total_paquete)}

	Registrar(output_file, data)

	return data
}

func Registrar(archivo string, datos []string){
	
	file, err := os.OpenFile(archivo, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	csvwriter := csv.NewWriter(file)
	_ = csvwriter.Write(datos)
	
 	csvwriter.Flush()
 
	file.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IniciarPrograma() {
    f, err := os.OpenFile("finanzas.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
	}
	f.Close()
}


func CerrarPrograma() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("\nCierre de operaciones de Finanzas\n ==============================\n Total = %f \n Ganancias = %f \n Perdidas = %f \n", total, ganancia, perdida)
		os.Exit(0)
	}()
}

func main() {
	IniciarPrograma()

	conn, err := amqp.Dial(connection)
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var m PaqueteProcesado

	CerrarPrograma()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			err := json.Unmarshal(d.Body, &m)
			failOnError(err, "JSON invalido")
			ActualizarInfo(m)
			// log.Printf("%s", data)
			// log.Printf("Received a message: \n Total:%f \t Ganancia:%f \t  Perdida:%f", total, ganancia, perdida)
		}


	}()

	log.Printf(" [*] Esperando Mensajes. Para salir presiones CTRL+C")
	<-forever

}