package main

//importar los paquetes time necesario
//para controlar los tiempos
import (
	"fmt"
	"time"
)

// go palabra reservada para las goroutine
func main() {
	go name("Guido")
	go name("Juan")
	fmt.Scanln()

}

// time para controlar los tiempos
func name(item string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second / 3)
		fmt.Println(item, i)
	}
}
