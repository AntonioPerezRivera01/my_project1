package main

import "log"

func main() {

	if 1 == 0 {
		log.Println("ok")
	} else {
		log.Println("Hola mundo como estas")
	}

	if 1==1{
		log.Println("Hola Hermoso llegaras lejos")
	} else{
		log.Println("Espero si llegues lejos man")
	}

	//Integracion del ciclo for
	for i := 0; i < 10; i++ {
		log.Println("*")
	}
}
