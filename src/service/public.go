package service

import (
	"encoding/json"
	"mechta-tz.github.com/src/config"
	"mechta-tz.github.com/src/domain"
	"sync"
)

func CalculateJson(data []byte) (*int, error) {
	var numbers []domain.TwoNumber
	if err := json.Unmarshal(data, &numbers); err != nil {
		return nil, err
	}

	chunkSize := (len(numbers) + config.MainConfig.WorkersCount - 1) / config.MainConfig.WorkersCount

	jobs := make(chan []domain.TwoNumber, config.MainConfig.WorkersCount)
	results := make(chan int, config.MainConfig.WorkersCount)

	var wg sync.WaitGroup

	for w := 0; w < config.MainConfig.WorkersCount; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		jobs <- numbers[i:end]
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for sum := range results {
		totalSum += sum
	}

	return &totalSum, nil
}
