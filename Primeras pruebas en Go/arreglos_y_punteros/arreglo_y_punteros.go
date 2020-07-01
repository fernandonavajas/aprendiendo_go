package main

import (
	"fmt"
	"strconv"
)

func main() {
	var z string
	x := 23
	name := "coco"
	z = "10"
	z_int, err := strconv.Atoi(z)
	fmt.Println("Hello, playground", x, name, z_int)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	println("")
	anonymous_field()
	println("")
	punteros()
	println("")
	arreglos()
}

func arreglos() {
	arreglo := []int{1, 2, 3, 14}
	fmt.Println(arreglo[3])

	var matriz [][]int
	if matriz == nil {
		fmt.Println("vacio")
	}

	arreglo_2 := [3]int{1, 2, 3}
	slice := arreglo_2[1:3]
	fmt.Println(slice)
}

func punteros() {
	var puntero *int
	entero := 5
	puntero = &entero
	fmt.Println(puntero)
	fmt.Println(*puntero)

	//Cambiar el valor, pero no la direccion
	*puntero = 6
	fmt.Println(puntero)
	fmt.Println(*puntero)
}

func anonymous_field() {
	type Human struct {
		name string
	}
	type Dummy struct {
		name string
	}
	type Tutor struct {
		Human
		Dummy
	}
	tutor := Tutor{Human{"Fernando"}, Dummy{"Damian"}}
	fmt.Println(tutor.Dummy.name)

}
