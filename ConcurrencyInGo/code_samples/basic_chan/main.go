package main

import "fmt"

func main() {
	numbers := generator()

	for n := range numbers {
		fmt.Printf("number is: %d\n", n)
	}
}

// Comportement attendu :

// la fonction crée un channel
// lance une goroutine
// envoie les entiers de 1 à 5
// ferme le channel à la fin
// retourne un channel receive-only
// Puis dans main, parcours les valeurs et affiche-les
func generator() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	return ch
}
