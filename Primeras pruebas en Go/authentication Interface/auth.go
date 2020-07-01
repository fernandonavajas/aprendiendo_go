package main

import (
	"fmt"
)

type User struct {
	nombre, apellido string
}

// administrador
type Cliente interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5

}
func (this Admin) Nombre() string {
	return this.nombre
}

//Editor
type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3

}
func (this Editor) Nombre() string {
	return this.nombre
}

func main() {
	admin := Admin{"Navajas"}
	editor := Editor{"Copano"}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
}

func auth(cliente Cliente) string {
	permisos := cliente.Permisos()
	if permisos > 4 {
		return cliente.Nombre() + " tiene permisos de administrador"
	} else if permisos == 3 {
		return cliente.Nombre() + " tiene permisos de editor"
	}
	return ""

}
