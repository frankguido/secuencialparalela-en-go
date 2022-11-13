package main

//importar los paquetes time necesario
//para controlar los tiempos
import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

// go palabra reservada para las goroutine
func main() {
	start := time.Now()
	wg.Add(2)
	go name("Guido")
	go name("Juan")
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("tiempo transcurrido", elapsed)
	fmt.Scanln()

}

// time para controlar los tiempos
func name(item string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second / 3)
		fmt.Println(item, i)
	}
	wg.Done()
}

