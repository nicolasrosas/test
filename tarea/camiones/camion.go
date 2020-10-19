package main
import (
	"encoding/csv"
	"fmt"
	"math/rand"
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

	PUERTO = ":3030"
	COM_LOGISTICA string = "localhost:4040"

	EstadoC1 int32 = 0//Camion retail 1 ; 0 = vacio ; 1 = esperando segundo paquete ; 2 = camion lleno/Camion en camino
	EstadoC2 int32 = 0//Camion retail 2
	EstadoC3 int32 = 0//Camion normal

	CargaC1 [2]Paquete
	CargaC2 [2]Paquete
	CargaC3 [2]Paquete

	TIEMPO_ESPERA int32
	TIEMPO_ENTREGA int32

	REGISTRO_C1 string = "Camion_Retail1.csv"
	REGISTRO_C2 string = "Camion_Retail2.csv"
	REGISTRO_C3 string = "Camion_Normal.csv"

)



func ComLogistica(paquete PaqueteProcesado){

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	cliente := proto.NewLogisticaServiceClient(conn)

	cliente.InformarEntrega(context.Background(), &proto.PaqueteEntregado{
																		IdPaquete: paquete.IdPaquete, 
																		Tipo: paquete.Tipo, 
																		Valor: paquete.Valor, 
																		Intentos: paquete.Intentos,
																		Timestamp: paquete.Timestamp,
																		IdCamion: paquete.IdCamion,
																		Estado: paquete.Estado,
																		Seguimiento: paquete.Seguimiento})

	conn.Close()



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

func EntregarPedido_2P(paquete *[2]PaqueteProcesado,NINT []int32, orden int32){

	var Exito int
	IntentosTotales := NINT[0]+NINT[1]
	Flag1 := 0
	Flag2 := 0

	for ((*paquete)[orden].Intentos + (*paquete)[1-orden].Intentos < IntentosTotales){

		time.Sleep(time.Duration(TIEMPO_ENTREGA) * time.Second)
		Exito = rand.Intn(9)
		if (Exito <= 7 && Flag1 == 0 && (*paquete)[orden].Intentos < NINT[orden]){
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			(*paquete)[orden].Intentos = (*paquete)[orden].Intentos +1
			(*paquete)[orden].Timestamp = timestamp
			Flag1 = 1
		}else if ((*paquete)[orden].Intentos < NINT[orden]){
			(*paquete)[orden].Intentos = (*paquete)[orden].Intentos +1
		}
		time.Sleep(time.Duration(TIEMPO_ENTREGA) * time.Second)
		Exito = rand.Intn(9)
		if (Exito <= 7 && Flag2 == 0  && (*paquete)[1-orden].Intentos < NINT[1-orden]){
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			(*paquete)[1-orden].Intentos = (*paquete)[1-orden].Intentos +1
			(*paquete)[1-orden].Timestamp = timestamp
			Flag2 = 1
		}else if ((*paquete)[1-orden].Intentos < NINT[1-orden]){
			(*paquete)[1-orden].Intentos = (*paquete)[1-orden].Intentos +1
		}

		if (Flag1 + Flag2  == 2){
			
			break

		}

		
	}
	if (Flag1 == 0){
		(*paquete)[orden].Estado = "No Recibido"
	}else{
		(*paquete)[orden].Estado = "Recibido"
	}
	if (Flag2 == 0){
		(*paquete)[1-orden].Estado = "No Recibido"
	}else{
		(*paquete)[1-orden].Estado = "Recibido"
	}

}
func EntregarPedido_1P(paquete *PaqueteProcesado, IntentosTotales int32, orden int32){

	Flag := 0

	for ((*paquete).Intentos < IntentosTotales){
		time.Sleep(time.Duration(TIEMPO_ENTREGA) * time.Second)
		Exito := rand.Intn(9)
		if (Exito <= 7 && Flag == 0){
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			(*paquete).Intentos = (*paquete).Intentos +1
			(*paquete).Timestamp = timestamp
			Flag = 1
		}else{
			(*paquete).Intentos = (*paquete).Intentos +1
		}

		if (Flag == 1){
			
			break

		}

		
	}
	if (Flag == 0){
		(*paquete).Estado = "No Recibido"

	}else{
		(*paquete).Estado = "Recibido"
	}

}
func En_Transito_2P(idcamion string, REGISTRO string, Carga [2]Paquete){

	var EstadoPaquete [2]PaqueteProcesado
	
	EstadoPaquete[0] = PaqueteProcesado{IdPaquete: Carga[0].IdPaquete, Tipo:Carga[0].Tipo, Valor: Carga[0].Valor, Intentos: 0, Timestamp: "0", IdCamion: idcamion, Estado:"En camino", Seguimiento:Carga[0].Seguimiento}
	EstadoPaquete[1] = PaqueteProcesado{IdPaquete: Carga[1].IdPaquete, Tipo:Carga[1].Tipo, Valor: Carga[1].Valor, Intentos: 0, Timestamp: "0", IdCamion: idcamion, Estado:"En camino", Seguimiento:Carga[1].Seguimiento}
	
	if (Carga[1].Valor < Carga[0].Valor){

		EntregarPedido_2P(&EstadoPaquete, []int32{Carga[0].Intentos, Carga[1].Intentos}, 0)

	}else{

		EntregarPedido_2P(&EstadoPaquete, []int32{Carga[1].Intentos, Carga[0].Intentos}, 1)
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println(EstadoPaquete[0].Intentos)
	fmt.Println(EstadoPaquete[1].Intentos)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	DATA1 := []string{EstadoPaquete[0].IdPaquete, EstadoPaquete[0].Tipo, strconv.Itoa(int(EstadoPaquete[0].Valor)), Carga[0].Origen, Carga[0].Destino, strconv.Itoa(int(EstadoPaquete[0].Intentos)), EstadoPaquete[0].Timestamp}
	DATA2 := []string{EstadoPaquete[1].IdPaquete, EstadoPaquete[1].Tipo, strconv.Itoa(int(EstadoPaquete[1].Valor)), Carga[1].Origen, Carga[1].Destino, strconv.Itoa(int(EstadoPaquete[1].Intentos)), EstadoPaquete[1].Timestamp}
	Registrar(REGISTRO, DATA1)
	Registrar(REGISTRO, DATA2)

	ComLogistica(EstadoPaquete[0])
	ComLogistica(EstadoPaquete[1])
	if (idcamion == "C1"){
		EstadoC1 = 0
	}else if (idcamion == "C2"){
		EstadoC2 = 0
	}else{
		EstadoC3 = 0
	}

}
func En_Transito_1P(idcamion string, REGISTRO string, Carga [2]Paquete){

	var EstadoPaquete PaqueteProcesado
	
	EstadoPaquete = PaqueteProcesado{IdPaquete: Carga[0].IdPaquete, Tipo:Carga[0].Tipo, Valor: Carga[0].Valor, Intentos: 0, Timestamp: "0", IdCamion: idcamion, Estado: "En camino", Seguimiento:Carga[0].Seguimiento}
	
	EntregarPedido_1P(&EstadoPaquete, Carga[0].Intentos, 0)


	DATA := []string{EstadoPaquete.IdPaquete, EstadoPaquete.Tipo, strconv.Itoa(int(EstadoPaquete.Valor)), Carga[0].Origen, Carga[0].Destino, strconv.Itoa(int(EstadoPaquete.Intentos)), EstadoPaquete.Timestamp}
	Registrar(REGISTRO, DATA)

	ComLogistica(EstadoPaquete)
	if (idcamion == "C1"){
		EstadoC1 = 0
	}else if (idcamion == "C2"){
		EstadoC2 = 0
	}else{
		EstadoC3 = 0
	}

}



func EsperaPaqueteC1(){

	time.Sleep(time.Duration(TIEMPO_ESPERA) * time.Second)
	if (EstadoC1 == 1){
		EstadoC1 = 2
		fmt.Println("///////////////")
		fmt.Println("C1 sale de la central con 1 paquete")
		En_Transito_1P("C1", REGISTRO_C1, CargaC1)
		
	}


}
func EsperaPaqueteC2(){

	time.Sleep(time.Duration(TIEMPO_ESPERA) * time.Second)
	if (EstadoC2 == 1){
		EstadoC2 = 2
		fmt.Println("///////////////")
		fmt.Println("C2 sale de la central con 1 paquete")
		En_Transito_1P("C2", REGISTRO_C2, CargaC2)
		
	}


}
func EsperaPaqueteC3(){

	time.Sleep(time.Duration(TIEMPO_ESPERA) * time.Second)
	if (EstadoC3 == 1){
		EstadoC3 = 2
		fmt.Println("///////////////")
		fmt.Println("C3 sale de la central con 1 paquete")
		En_Transito_1P("C3", REGISTRO_C3, CargaC3)
		
	}


}
/*
func EstadoPaquete(Codigo proto.Trackcode){

	return 

	 
}
func (s *server) ConsultaPaquete(ctx context.Context, Consulta *proto.ConsultaCamion) (*proto.OrderStatus, error){

	var ESTADO string 
	var INTENTOS int32

	if (Consulta.GetIdCamion() == "C1"){

		if (CargaC1[0].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC1[0].Estado 
			INTENTOS = CargaC1[0].Intentos 
		}else if (CargaC1[1].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC1[1].Estado 
			INTENTOS = CargaC1[1].Intentos 
		}else{
			ESTADO = "Paquete no encontrado"
			INTENTOS = 0
		}

	}else if (Consulta.GetIdCamion() == "C2"){
		if (CargaC2[0].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC2[0].Estado 
			INTENTOS = CargaC2[0].Intentos 
		}else if (CargaC2[1].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC2[1].Estado 
			INTENTOS = CargaC2[1].Intentos 
		}else{
			ESTADO = "Paquete no encontrado"
			INTENTOS = 0
		}
	}else{

		if (CargaC3[0].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC3[0].Estado 
			INTENTOS = CargaC3[0].Intentos 
		}else if (CargaC3[1].Seguimiento == Consulta.GetIdSeguimiento()){
			ESTADO = CargaC3[1].Estado 
			INTENTOS = CargaC3[1].Intentos 
		}else{
			ESTADO = "Paquete no encontrado"
			INTENTOS = 0
		}
	}

	return &proto.OrderStatus{Estado:ESTADO, Intentos:INTENTOS},nil

}
*/

func (s *server) AsignarC1(ctx context.Context, paquete *proto.Package) (*proto.RespuestaCamion, error){

	idcamion := "C1"
	fmt.Println("///////////////")
	fmt.Println("Paquete recibido C1: ",paquete)
	if (EstadoC1 == 0){

		EstadoC1 = 1
		CargaC1[0] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}
		go EsperaPaqueteC1()
		fmt.Println("///////////////")
		fmt.Println("C1 en espera de un segundo paquete")

	}else if (EstadoC1 == 1){

		EstadoC1 = 2
		CargaC1[1] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}

		go En_Transito_2P("C1", REGISTRO_C1, CargaC1)
		fmt.Println("///////////////")
		fmt.Println("C1 sale de la central con 2 paquetes")

	}else{
		idcamion = "no asignado"
		fmt.Println("///////////////")
		fmt.Println("C1 no esta disponible")
	}

	return &proto.RespuestaCamion{Idcamion:idcamion},nil

}
func (s *server) AsignarC2(ctx context.Context, paquete *proto.Package) (*proto.RespuestaCamion, error){

	idcamion := "C2"
	fmt.Println("///////////////")
	fmt.Println("Paquete recibido C2: ",paquete)
	if (EstadoC2 == 0){

		EstadoC2 = 1
		CargaC2[0] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}
		go EsperaPaqueteC2()
		fmt.Println("///////////////")
		fmt.Println("C2 en espera de un segundo paquete")

	}else if (EstadoC2 == 1){

		EstadoC2 = 2
		CargaC2[1] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}

		go En_Transito_2P("C2", REGISTRO_C2, CargaC2)
		fmt.Println("///////////////")
		fmt.Println("C2 sale de la central con 2 paquetes")

	}else{
		idcamion = "no asignado"
		fmt.Println("///////////////")
		fmt.Println("C2 no esta disponible")
	}

	return &proto.RespuestaCamion{Idcamion:idcamion},nil

}
func (s *server) AsignarC3(ctx context.Context, paquete *proto.Package) (*proto.RespuestaCamion, error){

	idcamion := "C3"
	fmt.Println("///////////////")
	fmt.Println("Paquete recibido C3: ",paquete)
	if (EstadoC3 == 0){

		EstadoC3 = 1
		CargaC3[0] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}
		go EsperaPaqueteC1()
		fmt.Println("///////////////")
		fmt.Println("C3 en espera de un segundo paquete")
	}else if (EstadoC3 == 1){

		EstadoC3 = 2
		CargaC3[1] = Paquete{		IdPaquete:paquete.GetIdPaquete(),
									Seguimiento:paquete.GetSeguimiento(),
									Tipo:paquete.GetTipo(),
									Valor:paquete.GetValor(),
									Intentos:paquete.GetIntentos(),
									Estado:paquete.GetEstado(),
									Origen:paquete.GetOrigen(),
									Destino:paquete.GetDestino()}

		go En_Transito_2P("C3", REGISTRO_C3, CargaC3)
		fmt.Println("///////////////")
		fmt.Println("C3 sale de la central con 2 paquetes")

	}else{
		idcamion = "no asignado"
		fmt.Println("///////////////")
		fmt.Println("C3 no esta disponible")
	}

	return &proto.RespuestaCamion{Idcamion:idcamion},nil

}

func INICIO(){

	csvFile, err := os.Create(REGISTRO_C1)
	if err != nil{
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()

	csvFile, err = os.Create(REGISTRO_C2)
	if err != nil{
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()

	csvFile, err = os.Create(REGISTRO_C3)
	if err != nil{
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()

	TIEMPO_ENTREGA = 15

	fmt.Println("Tiempo de espera de segundo paquete [segundos]: ")
	fmt.Scan(&TIEMPO_ESPERA)
	


}

func main(){

	INICIO()

	listener, err :=net.Listen("tcp",":3030")
	if err != nil{
		panic(err)
	}
	srv := grpc.NewServer()

	proto.RegisterCamionesServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e!= nil{
		panic(e)
	}


}