package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan string)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Go(func() {
		worker(done, jobs)
	})

	jobs <- "1"
	jobs <- "2"
	jobs <- "3"

	close(done)
	wg.Wait()
}

// Comportement :
// le worker boucle en permanence
// s’il reçoit un job, il affiche : job: <valeur>
// si done est fermé, il affiche worker stopped puis retourne
// Dans main :
// crée done et jobs
// lance le worker
// envoie 2 ou 3 jobs
// ferme done

// Ce que tu travailles
// squelette de for-select
// arrêt propre
// canal de cancellation
func worker(done <-chan struct{}, jobs <-chan string) {
	for {
		select {
		case <-done:
			fmt.Println("worker stopped")
			return
		case j, ok := <-jobs:
			if !ok {
				fmt.Println("jobs channel closed")
				return
			}
			fmt.Printf("job: %s\n", j)
		}
	}
}
