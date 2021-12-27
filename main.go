package main

import "log"

func main() {

	if 1 == 0 {
		log.Println("ok")
	} else {
		log.Println("Hola mundo como estas")
	}

	if 1 == 1 {
		log.Println("Hola Hermoso llegaras lejos")
	} else {
		log.Println("Espero si llegues lejos man")
	}

	for x := 0; x < 10; x++ {
		log.Println("Valor: " + string(x))
	}

	animal := "Lobo"
	switch animal {
	case "Lobo":
		log.Println("Es un Lobo")
		break
	case "Perro":
		log.Println("Es un Perro")
		break

	}
}
