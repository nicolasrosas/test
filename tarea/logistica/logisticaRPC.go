package main

import (
	"encoding/csv"
	////// Comunicación Logística Finanzas
	"encoding/json"
	"github.com/streadway/amqp"
	////////
	"fmt"
	"os"
	"log"
	"time"
	"strconv"
	"context"
	"tarea/proto"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

/*
type ClientRequest struct{

    idProducto string
    producto string
    valor string
    origen string
    destino string
    prioritario string
    
}
*/
type Paquete struct{

	IdPaquete string
	Seguimiento string
	Tipo string
	Valor int32
	Intentos int32
	Estado string
	Origen string
	Destino string
}

type PaqueteSeguimiento struct{

	idPaquete string
	estado string
	idCamion string
	idSeguimiento string
	intentos int32

}

var (

	REGISTRO_LOG string = "registro.csv"
	PUERTO string = ":4040"
	COM_CAMIONES = "localhost:3030"

	IDpaquete int32 = 0
	Seguimiento int32 = 13000

	COLA_RETAIL []Paquete
	COLA_PRIORITARIO []Paquete
	COLA_PYME []Paquete

	EstadoPaquete = map[string]PaqueteSeguimiento{}

	TIEMPO_DESPACHO int32	

	// variables globales de conexion rabbitmq
	mqAddress string = "amqp://user:pass@10.6.40.74:5672/"
	mqConn, err = amqp.Dial(mqAddress)
	mqChan, err2 = mqConn.Channel()
	mqQueue, err3 = mqChan.QueueDeclare(
		"finanzas", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

)

/*

Funcion comunicación logistica - finanzas

*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func EnviarFinanzas(paquete *proto.PaqueteEntregado){

	body, err := json.Marshal(paquete)
	failOnError(err, "JSON invalido")
	err = mqChan.Publish(
		"",     // exchange
		mqQueue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	
	log.Printf("//////// \n[x] Se ha enviado informacion a finanzas %s \n////////", body)
}

/////////////////////////



func ComCamiones_Asignar(paquete Paquete, camion string) string{

	var respuesta string

	conn, err := grpc.Dial(COM_CAMIONES, grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	cliente := proto.NewCamionesServiceClient(conn)

	if (camion == "C1"){
		response, err := cliente.AsignarC1(context.Background(), &proto.Package{IdPaquete:paquete.IdPaquete,
																				Seguimiento:paquete.Seguimiento,
																				Tipo:paquete.Tipo,
																				Valor:paquete.Valor,
																				Intentos:paquete.Intentos,
																				Estado:paquete.Estado,
																				Origen:paquete.Origen,
																				Destino:paquete.Destino})
		if err != nil{
			panic(err)
		}
		respuesta = response.Idcamion
	}else if (camion == "C2"){
		response, err := cliente.AsignarC2(context.Background(), &proto.Package{IdPaquete:paquete.IdPaquete,
																				Seguimiento:paquete.Seguimiento,
																				Tipo:paquete.Tipo,
																				Valor:paquete.Valor,
																				Intentos:paquete.Intentos,
																				Estado:paquete.Estado,
																				Origen:paquete.Origen,
																				Destino:paquete.Destino})
		if err != nil{
			panic(err)
		}
		respuesta = response.Idcamion
	}else{
		response, err := cliente.AsignarC3(context.Background(), &proto.Package{IdPaquete:paquete.IdPaquete,
																				Seguimiento:paquete.Seguimiento,
																				Tipo:paquete.Tipo,
																				Valor:paquete.Valor,
																				Intentos:paquete.Intentos,
																				Estado:paquete.Estado,
																				Origen:paquete.Origen,
																				Destino:paquete.Destino})
		if err != nil{
			panic(err)
		}
		respuesta = response.Idcamion
	}
	

	conn.Close()

	return respuesta



}
/*
func ComCamiones_Estado(consulta *proto.ConsultaCamion) []string{

	conn, err := grpc.Dial(COM_CAMIONES, grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	cliente := proto.NewCamionesServiceClient(conn)

	
	respuesta, err := cliente.ConsultaPaquete(context.Background(), consulta)
	if err != nil{
		panic(err)
	}
	

	conn.Close()

	return []string{respuesta.Estado, strconv.Itoa(int(respuesta.Intentos))}



}
*/
/*
func Actualizar_datos(codigo string){

	DATOS := ComCamiones_Estado(&proto.ConsultaCamion{IdSeguimiento:EstadoPaquete[codigo].idPaquete, IdCamion:EstadoPaquete[codigo].idCamion})
	INT, _ :=strconv.Atoi(DATOS[1])
	TEMP := EstadoPaquete[codigo]
	TEMP.estado = DATOS[0]
	TEMP.intentos = int32(INT)
	EstadoPaquete[codigo] = TEMP
	
}
*/
func (s *server) EstadoPedido(ctx context.Context, codigo *proto.TrackCode) (*proto.OrderStatus, error){

	//Actualizar_datos(codigo.GetCode())
	var estado string
	var intentos int32

	TEMP, ok := EstadoPaquete[codigo.GetCode()]

	if (ok){
		estado = TEMP.estado
		intentos = TEMP.intentos
		
	}else{
		estado = "Paquete inexistente"
		intentos = 0
	}

	return &proto.OrderStatus{Estado: estado, Intentos:intentos},nil



}


func registrar(archivo string, datos []string){
	
	file, err := os.OpenFile(archivo, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	csvwriter := csv.NewWriter(file)
	_ = csvwriter.Write(datos)
	
 	csvwriter.Flush()
 
	file.Close()
}

func (s *server) InformarEntrega(ctx context.Context, paquete *proto.PaqueteEntregado) (*proto.Empty, error){
	
	fmt.Println("/////////////")
	fmt.Printf("Camion %s llega a la central\n",paquete.GetIdCamion())
	fmt.Printf("informacion de la entrega :\n")
	fmt.Printf("Paquete %s fue %s con %d intentos\n",paquete.GetIdPaquete(), paquete.GetEstado(), paquete.GetIntentos())
	TEMP := EstadoPaquete[paquete.GetSeguimiento()]
	TEMP.intentos = paquete.GetIntentos()
	TEMP.estado = paquete.GetEstado()
	EstadoPaquete[paquete.GetSeguimiento()] = TEMP

	// Info para finanzas 
	
	EnviarFinanzas(paquete)

	return &proto.Empty{},nil

}

func GuardarSeguimento(paquete Paquete, camion string, Estado string){
	
	EstadoPaquete[paquete.Seguimiento] = PaqueteSeguimiento{idPaquete:paquete.IdPaquete, estado:Estado, idCamion:camion, idSeguimiento:paquete.Seguimiento, intentos:0}

}

func Asignar(COLA *[]Paquete, idcamion string) string{

	respuesta := ComCamiones_Asignar((*COLA)[0], idcamion)
	if (respuesta != "no asignado"){

		fmt.Println("/////////////")
		fmt.Printf("Paquete %s, convenio %s, asignado a camion %s\n",(*COLA)[0].IdPaquete, (*COLA)[0].Tipo, idcamion)
		GuardarSeguimento((*COLA)[0], idcamion,"En camino")
		dequeue(COLA)
				
	}
	return respuesta

}


func Despacho(){

	
	for {
		time.Sleep(time.Duration(TIEMPO_DESPACHO) * time.Second)
		for (len(COLA_RETAIL) > 0){

			R := Asignar(&COLA_RETAIL, "C1")

			if (R == "no asignado"){
				break
			}
		}
		for (len(COLA_RETAIL) > 0){

			R := Asignar(&COLA_RETAIL, "C2")

			if (R == "no asignado"){
				break
			}
		}
		for (len(COLA_PRIORITARIO) > 0){

			R := Asignar(&COLA_PRIORITARIO, "C1")
			
			if (R == "no asignado"){
				break
			}
		}
		for (len(COLA_PRIORITARIO) > 0){

			R := Asignar(&COLA_PRIORITARIO, "C2")
			
			if (R == "no asignado"){
				break
			}
		}
		for (len(COLA_PRIORITARIO) > 0){

			R := Asignar(&COLA_PRIORITARIO, "C3")
			
			if (R == "no asignado"){
				break
			}
		}
		for (len(COLA_PYME) > 0){

			R := Asignar(&COLA_PYME, "C3")
			
			if (R == "no asignado"){
				break
			}
		}
		

	}

	

}



func enqueue(cola *[]Paquete, pack Paquete){

	(*cola) = append((*cola), pack)
	fmt.Println("/////////////")
	fmt.Printf("Paquete %s en cola %s\n", pack.IdPaquete, pack.Tipo)

}

func dequeue(cola *[]Paquete){
	
	(*cola) = (*cola)[1:]

}


func CalcularIntentos(convenio string, valor int32) int32{

	var N_intentos int32

	if (convenio == "retail"){
		N_intentos = 3
	}else{
		N_intentos = int32(valor / 10)
		if (N_intentos >= 2){
			N_intentos = 2
		}else {
			N_intentos = 1
		}
	}

	return N_intentos

}


func (s *server) Ordenar(ctx context.Context, request *proto.ClientRequest) (*proto.TrackCode, error) {

	var CodigoSeguimiento string
	var tipo string

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("/////////////")
	fmt.Println("Orden recibida: ", request)

	if (request.GetConvenio() == "pyme"){
		if (request.GetPrioritario() == "1"){
			tipo = "Prioritario"
			CodigoSeguimiento = strconv.Itoa(int(Seguimiento))
		}else{
			tipo = "normal"
			CodigoSeguimiento = strconv.Itoa(int(Seguimiento))
		}
	}else{
		tipo = "retail"
		CodigoSeguimiento = "nil"
	}

	NINT := CalcularIntentos(tipo, request.GetValor())

	data := []string{timestamp, strconv.Itoa(int(IDpaquete)), tipo, request.GetProducto(), strconv.Itoa(int(request.GetValor())), request.GetOrigen(), request.GetDestino(), CodigoSeguimiento}
	pack := Paquete{IdPaquete:strconv.Itoa(int(IDpaquete)), Seguimiento: strconv.Itoa(int(Seguimiento)), Tipo:tipo, Valor:request.GetValor(), Intentos: NINT, Estado: "En Bodega", Origen:request.GetOrigen(), Destino:request.GetDestino()}
	
	if (tipo == "retail"){
		enqueue(&COLA_RETAIL, pack)
	}else if (tipo == "Prioritario"){
		enqueue(&COLA_PRIORITARIO, pack)
	}else{
		enqueue(&COLA_PYME, pack)
	}
	GuardarSeguimento(pack,"sin asignar","En Bodega")
	/*
	fmt.Println("COLA RETAIL: ", len(COLA_RETAIL))
	fmt.Println("COLA PRIORITARIO: ", len(COLA_PRIORITARIO))
	fmt.Println("COLA PYME: ", len(COLA_PYME))
	*/
	IDpaquete = IDpaquete + 1
	Seguimiento = Seguimiento + 1
	
	registrar(REGISTRO_LOG,data)
	
	return &proto.TrackCode{Code:CodigoSeguimiento},nil

}

func INICIO(){

	csvFile, err := os.Create(REGISTRO_LOG)
	if err != nil{
		log.Fatalf("failed creating file: %s", err)
	}

	csvFile.Close()
	TIEMPO_DESPACHO = 10
}

func main(){
	/*
	estructura := []string{"timestamp","id-paquete","tipo","nombre","valor","origen","destino","seguimiento"}

	file, err := os.Create("registro.csv")
	if err != nil{
		log.Fatal(err)
	}
	file.Close()
	registrar(estructura)

	dato:= ClientRequest{idProducto:"23423",producto:"123",valor:234,origen:"123",destino:"123",prioritario: true}
	response := Ordenar(dato)

	fmt.Println(response)
	*/
	INICIO()

	//  conexión rabbitmq

	
	// failOnError(err, "Conexion a RabbitMQ fallida")
	// defer mqConn.Close()

	
	// failOnError(err2, "Apertura de canal fallida")
	// defer mqChan.Close()

	//

	listener, err :=net.Listen("tcp",PUERTO)
	
	if err != nil{
		panic(err)
	}
	srv := grpc.NewServer()
	
	proto.RegisterLogisticaServiceServer(srv, &server{})
	reflection.Register(srv)
	go Despacho()
	if e := srv.Serve(listener); e!= nil{
		panic(e)
	}
	

}