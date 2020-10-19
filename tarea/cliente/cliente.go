package main
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"io"
	"google.golang.org/grpc"
	"tarea/proto"
	"context"
	"math/rand"
	"time"
	"strconv"
	"bufio"

)
/*
type Pedido struct{

    idProducto string
    producto string
    valor int32
    origen string
    destino string
    prioritario bool
    
}*/

var (

	dirPymes string = "productos/pymes.csv"
	dirRetail string = "productos/retail.csv"
	
	PEDIDOSPYMES [][]string
	PEDIDOSRETAIL [][]string
	CODIGOS []string

	CON_LOGISTICA string = "localhost:4040"

	TIEMPO_PYMES int32 = 1
	TIEMPO_RETAIL int32 = 1

)


func Fill(dir string, tipo bool){

	
	csvfile, err := os.Open(dir)
	
	if err != nil {
		log.Fatalf("Couldn't open the csv file: %s", err)
	}

	csvreader := csv.NewReader(csvfile)

	for {

		record, err := csvreader.Read()
		if (err == io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}



		if (tipo){
			PEDIDOSPYMES = append(PEDIDOSPYMES, record)

		}else{
			PEDIDOSRETAIL = append(PEDIDOSRETAIL, record)
		}


	}

}


func RealizarPedido(tipo int) []string{
	var data []string

	if (tipo == 0){
		seleccion := rand.Intn(len(PEDIDOSPYMES))
		data = PEDIDOSPYMES[seleccion]

	}else{
		seleccion := rand.Intn(len(PEDIDOSRETAIL))
		data = PEDIDOSRETAIL[seleccion]
	}

	return data
}

func PreguntarEstado(codigo string){

	conn, err := grpc.Dial(CON_LOGISTICA, grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	client := proto.NewLogisticaServiceClient(conn)
	response, err := client.EstadoPedido(context.Background(), &proto.TrackCode{Code: codigo})
	if err != nil{
		panic(err)
	}
	fmt.Println("ID Paquete: ",codigo)
	fmt.Println("Estado del paquete: ", response.Estado)
	fmt.Println("Numero de intentos: ", response.Intentos)
}
func PYME(){

	conn, err := grpc.Dial(CON_LOGISTICA, grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	client := proto.NewLogisticaServiceClient(conn)

	for {
		dato := RealizarPedido(0)
		valor, _ := strconv.Atoi(dato[2])
		response, err := client.Ordenar(context.Background(), &proto.ClientRequest{IdProducto:dato[0],
																					Producto:dato[1], 
																					Valor:int32(valor), 
																					Origen:dato[3], 
																					Destino:dato[4], 
																					Prioritario:dato[5], 
																					Convenio:"pyme"})
		if err != nil{
			panic(err)
		}
		CODIGOS = append(CODIGOS,response.Code)
		time.Sleep(time.Duration(TIEMPO_PYMES)*time.Second)
		
	}

	

}
func RETAIL(){

	conn, err := grpc.Dial(CON_LOGISTICA, grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	client := proto.NewLogisticaServiceClient(conn)
	for {

		dato := RealizarPedido(1)
		valor, _ := strconv.Atoi(dato[2])
		_, err := client.Ordenar(context.Background(), &proto.ClientRequest{IdProducto:dato[0],
																					Producto:dato[1], 
																					Valor:int32(valor), 
																					Origen:dato[3], 
																					Destino:dato[4],
																					Convenio:"retail"})
		if err != nil{
			panic(err)
		}
		time.Sleep(time.Duration(TIEMPO_RETAIL)*time.Second)
	}
	
}

func INICIO(){


	Fill(dirPymes, true)
	Fill(dirRetail, false)
	fmt.Println("Tiempo entre envios de pymes [segundos]: ")
	fmt.Scan(&TIEMPO_PYMES)
	fmt.Println("Tiempo entre envios de retail [segundos]: ")
	fmt.Scan(&TIEMPO_RETAIL)

}


func main(){

	var input int

	
	INICIO()

	

	for {
		fmt.Println("Opciones: ")
		fmt.Println("1.-Crear cliente Pyme")
		fmt.Println("2.-Crear cliente Retail")
		fmt.Println("3.-Solicitar estado de pedido")
		fmt.Scan(&input)

		if (input == 1){
			go PYME()
		}else if (input == 2){
			go RETAIL()
		}else{
			fmt.Println("Opciones: ")
			fmt.Println("1.-Manual: ")
			fmt.Println("2.-Aleatorio: ")
			fmt.Scan(&input)
			if (input == 1){
				fmt.Println("Codigos disponibles (x para cancelar): ")
				fmt.Println(CODIGOS)
				reader := bufio.NewReader(os.Stdin)
				code, _ := reader.ReadString('\n')
				if (code != "x"){
					PreguntarEstado(code[:len(code)-1])
				}
				

			}else{
				if (len(CODIGOS)>0){
					seleccion := rand.Intn(len(CODIGOS))

					PreguntarEstado(CODIGOS[seleccion])
				}else{
					fmt.Println("No ser han realizado pedidos")
				}
				
			}
			
		}

	}

	
	
	
}

