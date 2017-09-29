package main

import (
	"sync"
	"log"
)
var (
	wg sync.WaitGroup
)
func main(){
	//cantidad de gorutinas a correr
	wg.Add(2)
	log.Print("Iniciamos las gorutinas...")
	go imprimir("A")
	go imprimir("B")
	log.Println("Esperando q finalicen...")
	wg.Wait()
	log.Println("Fin...")
}
func imprimir(obj string){
	//una corutina ya echa
	//defer se ejecuta al finalizaar la funcion
	defer wg.Done()
	for i:=0;i<10;i++{
		log.Printf("%d:"+obj,i)
	}
}