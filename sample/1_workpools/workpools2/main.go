package main

import (
	"fmt"
	"github.com/wazsmwazsm/mortar"
	"sync"
)

func main() {
	pool, err := mortar.NewPool(10)
	if err != nil {
		panic(err)
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		task := &mortar.Task{
			Handler: func(v ...interface{}) {
				wg.Done()
				fmt.Println(v)
			},
			Params: []interface{}{i, i * 2, "hello"},
		}

		pool.Put(task)
	}

	wg.Wait()

	pool.Close()

	//err = pool.Put(&mortar.Task{
	//	Handler: func(v ...interface{}) {
	//
	//	},
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
}
