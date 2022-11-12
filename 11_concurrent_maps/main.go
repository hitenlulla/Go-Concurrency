/*
Consider routines that are performing operations on a shared Map object.
Simultaneous operations on a single map is not allowed.
To solve this, use a regular Map with mutex or use sync.Map
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	regularMap := make(map[int]interface{})
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		regularMap[0] = i
	// 	}()
	// }

	syncMap := sync.Map{}

	// Put
	regularMap[0] = 0
	syncMap.Store(0, 0)

	// get
	regularValue, regularOk := regularMap[0]
	fmt.Println(regularValue, regularOk)

	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)

	// delete
	regularMap[1] = nil

	syncMap.Delete(1)

	mu := sync.Mutex{}
	// Load or Store
	mu.Lock()
	regularValue, regularOk = regularMap[1]
	if regularOk {
		regularMap[1] = 1
		regularValue = regularMap[1]
	}
	mu.Unlock()
	fmt.Println(regularValue)

	syncVal, loaded := syncMap.LoadOrStore(1, 1)
	fmt.Println(syncVal, loaded)

	// Load and Delete
	mu.Lock()
	regularValue = regularMap[1]
	mu.Unlock()
	fmt.Println(regularValue)

	syncVal, loaded = syncMap.LoadAndDelete(1)
	fmt.Println(syncVal, loaded)

	// Iteration
	for key, val := range regularMap {
		fmt.Print(key, val, "|")
	}
	fmt.Println()

	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, value, "|")
		return true
	})
}
