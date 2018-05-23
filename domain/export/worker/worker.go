package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/ricardolonga/workshop-go/domain"
)

type worker struct{}

func New() *worker {
	return &worker{}
}

func (w *worker) Do(export *domain.Export, wg *sync.WaitGroup) {
	fmt.Printf("escrevendo no arquivo: %s\n", export.Filename)
	time.Sleep(5 * time.Second)
	wg.Done()
}
