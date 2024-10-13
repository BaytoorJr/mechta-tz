package service

import (
	"mechta-tz.github.com/src/domain"
	"sync"
)

func worker(jobs <-chan []domain.TwoNumber, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for objects := range jobs {
		for _, obj := range objects {
			sum += obj.A + obj.B
		}
	}
	results <- sum
}
