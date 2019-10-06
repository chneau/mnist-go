package mnist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var baseURL = "https://raw.githubusercontent.com/cazala/mnist/master/src/digits/%d.json"

type Mnist map[int][][]float32

func (m Mnist) Get(n int) [][]float32 {
	return m[n]
}

func (m Mnist) Initiliaze() error {
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	errors := []error{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			url := fmt.Sprintf(baseURL, i)
			resp, err := http.Get(url)
			mtx.Lock()
			defer mtx.Unlock()
			defer wg.Done()
			if err != nil {
				errors = append(errors, err)
				return
			}
			x := &struct {
				Data []float32 `json:"data"`
			}{}
			err = json.NewDecoder(resp.Body).Decode(x)
			if err != nil {
				errors = append(errors, err)
				return
			}
			result := [][]float32{}
			for i := 0; i < len(x.Data); i++ {
				x.Data[i] *= 255
			}
			for i := 0; i < len(x.Data); i += 28 * 28 {
				result = append(result, x.Data[i:i+28*28])
			}
			m[i] = result
		}()
	}
	wg.Wait()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func New() Mnist {
	m := Mnist{}
	return m
}
