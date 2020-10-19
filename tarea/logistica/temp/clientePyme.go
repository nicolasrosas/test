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
			PEDIDOSRETAIL = append(PEDIDOSPYMES, record)
		}


	}

}


func RealizarPedido() []string{
	var data []string

	tipo := rand.Intn(1)//0:pyme, 1:retail

	if (tipo == 0){
		seleccion := rand.Intn(len(PEDIDOSPYMES)-1)
		data = PEDIDOSPYMES[seleccion]

	}else{
		seleccion := rand.Intn(len(PEDIDOSRETAIL)-1)
		data = PEDIDOSRETAIL[seleccion]
	}

	return data
}



func main(){

	Fill(dirPymes, true)
	Fill(dirRetail, false)
	
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	client := proto.NewVentaClient(conn)


	for {

		dato := RealizarPedido()
		
		response, err := client.Ordenar(context.Background(), &proto.ClientRequest{IdProducto:dato[0], Producto:dato[1], Valor:dato[2], Origen:dato[3], Destino:dato[4], Prioritario:dato[5]})
		if err != nil{
			panic(err)
		}
		fmt.Println(response)
		time.Sleep(4* time.Second)
	}

	
	
	
}

