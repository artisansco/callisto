package server

import (
	"fmt"
	"net/http"
	"sync"
)

func BackgroundTask(r *http.Request, fn func() error) {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		defer func() {
			pv := recover()
			if pv != nil {
				reportServerError(r, fmt.Errorf("%v", pv))
			}
		}()

		err := fn()
		if err != nil {
			reportServerError(r, err)
		}
	}()
}
