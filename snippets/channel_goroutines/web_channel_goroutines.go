// Goroutines and channels

// Ejemplo de goroutines y channels. Varias cosas

// - Las instrucciones go son por defecto *fire and forget*. Para que esto no sea así, necesitamos los canales.
// - Los canales (chan) son como un grupo de tasks o hilos. Hacen wait o await cuando se lee del canal.
// - Podemos especificar la devolución de los canales.
// - Con <-channel nos ponemos a la espera del siguiente mensaje

package main

import (
	"fmt"
	"net/http"
	"time"
)

func testServer(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		// fmt.Println(url, " !!! KO :(")
		channel <- url + " !!! KO :("
	} else {
		// fmt.Println(url, " OK :)")
		channel <- url + " OK :)"
	}
}

func main() {
	initTime := time.Now()

	channel := make(chan string) // Creamos los canales
	urls := []string{
		"http://oregoo.com/",
		"http://www.udemy.com/",
		"http://www.youtube.com/",
		"http://www.facebook.com/",
		"http://www.google.com/",
	}

	for _, url := range urls {
		go testServer(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-channel)
	}

	totalTime := time.Since(initTime)
	fmt.Println("Total time: ", totalTime)
}
