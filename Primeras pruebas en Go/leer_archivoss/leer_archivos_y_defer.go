package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	primeraForma()
	ejecucionsegundaForma := segundaForma()
	fmt.Println("Se ejecuto ejecucionsegundaForma: ", ejecucionsegundaForma)
	fmt.Println("Fin")
}
func primeraForma() {
	file_data, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")

	}
	fmt.Println(string(file_data))

}

func segundaForma() bool {
	archivo, error := os.Open("./archiavo.txt")
	defer func() {
		archivo.Close()
		fmt.Println("defer")
		r := recover() //no levantes el panico
		fmt.Println("r: ", r)
	}()
	if error != nil {
		fmt.Println("hubo un eror")
		panic(error) //levantar error
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(linea)
		fmt.Println(i)

	}
	return true
}
