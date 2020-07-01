package main

import (
	"fmt"
)

type User struct {
	nombre, apellido string
}

func main() {
	var fernando User
	felipe := new(User)
	felipe.nombre = "Felipe"
	felipe.apellido = "Leon"
	fmt.Println(fernando)
	fmt.Println(felipe)
	felipe.set_name("matias") //Cambiar valor por puntero
	fmt.Println(User{nombre: "Fernando", apellido: "navajas"})

	fmt.Println(felipe.nombre_completo())
}

func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
}
func (this *User) set_name(n string) {
	this.nombre = n
}
