package main2
import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"time"
	"strconv"
)


type ClientRequest struct{

    idProducto string
    producto string
    valor int32
    origen string
    destino string
    prioritario bool
    
}

type Paquete struct{

	IDPaquete uint32
	seguimiento uint32
	tipo uint32
	valor uint32
	intentos uint32
	estado string

}

var (
	IDpaquete int32 = 0
	Seguimiento int32 = 13000
)

func registrar(datos []string){
	
	file, err := os.OpenFile("registro.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	csvwriter := csv.NewWriter(file)
	_ = csvwriter.Write(datos)
	
 	csvwriter.Flush()
 
	file.Close()
}

func Ordenar(request ClientRequest) int32 {

	var tipo string

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if (request.origen == "pyme"){
		if (request.prioritario){
			tipo = "Prioritario"
		}else{
			tipo = "normal"
		}
	}else{
		tipo = "retail"
	}
	
	data := []string{timestamp, strconv.Itoa(int(IDpaquete)), tipo, request.producto, strconv.Itoa(int(request.valor)), request.origen, request.destino, strconv.Itoa(int(Seguimiento))}
	fmt.Println(data)
	IDpaquete = IDpaquete + 1
	Seguimiento = Seguimiento + 1
	
	registrar(data)
	
	return Seguimiento - 1

}


func main(){
	
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

}