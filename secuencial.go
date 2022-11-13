package main

//importar las librerias necesarias
import (
	"fmt"
)

// funcion main
func main() {
	//creando dos procesos para la funcion main
	name("Nat")
	name("Felix")

}

// funcion process para imprimir n veces cada nombre
func name(item string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(item, i)
	}
}
