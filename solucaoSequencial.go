package main

import (
	"fmt"
	"time"
)

var (
	phil = []string{"Aristoteles", "Platao", "Socrates", "Tales", "Nietzsche"}
	repeticoes = 1
	tempoDeSleep = time.Millisecond
)

func eat(name string){
	fmt.Println(name + " esta comendo");
}

func think(name string){
	fmt.Println(name + " esta pensando");
}

func main(){

	start := time.Now()
	for k := 0; k < repeticoes; k++ {
		for i := 0; i < len(phil); i++ {
			eat(phil[i])
			for j := 0; j < len(phil); j++ {
				if j != i {
					think(phil[j]);
				}
			}
			fmt.Println("<----------------------------------->")
			time.Sleep(tempoDeSleep)
		}	
	}
	execTime := time.Since(start)
	fmt.Printf("Tempo de execucao: %s", execTime)
}