package main
 
import (
	//"encoding/csv"
	//"log"
	//"os"
	"fmt"
	"time"
)
/*
func Ordenar() int32 {
	
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}
	rows := []string{"Name", "City", "Language"}
	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	csvwriter := csv.NewWriter(file)
	_ = csvwriter.Write(rows)
	for e, row := range rows {
		fmt.Println(e)
		_ = csvwriter.Write(row)
	}
 
	csvwriter.Flush()
 
	file.Close()

	return 1
}
*/

func main() {
	X := [][]string{{"holi","holi"},{"holi","holi"},{"holi","holi"}}
	time.Sleep(6)
	fmt.Println(len(X))


}