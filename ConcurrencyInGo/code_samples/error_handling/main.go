package main

import (
	"errors"
	"fmt"
)

func main() {
	done := make(chan struct{})
	nums := []int{3, -6, 2, -2, 4}

	results := square(done, nums...)

	for r := range results {
		if r.Err != nil {
			fmt.Printf("input %d -> error: %v\n", r.Input, r.Err)
			continue
		}
		fmt.Printf("input %d -> value: %d\n", r.Input, r.Value)
	}

}

type Result struct {
	Input int
	Value int
	Err   error
}

// Pour chaque nombre :

// si le nombre est négatif, envoie une erreur
// sinon envoie son carré

// Exemples :

// 2 → 4
// 5 → 25
// -3 → erreur
// Contraintes
// le travail doit être fait dans une goroutine
// les résultats doivent passer par un channel
// le channel doit être fermé à la fin
// si done est fermé, la goroutine doit s’arrêter
// Ce que tu travailles
// envoyer résultat + erreur
// struct de communication
// cancellation
func square(done <-chan struct{}, nums ...int) <-chan Result {
	result := make(chan Result)

	go func() {
		defer close(result)

		for _, n := range nums {
			var r Result
			if n < 0 {
				r = Result{
					Input: n,
					Value: n,
					Err:   errors.New("couldn't square negative number"),
				}
				continue
			}
			r = Result{
				Input: n,
				Value: n * n,
				Err:   nil,
			}
			select {
			case <-done:
				return
			case result <- r:
			}
		}
	}()

	return result
}
