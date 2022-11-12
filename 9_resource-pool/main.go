package main

import (
	"fmt"
	"sync"
)

/*
Every go-routine requires some memory to perform operations.
Suppose there are 1Million go routines that do the same task, if each routine requires 1 unit of memory, 1M units of memory will be used to perform a same task.
This memory can be saved effeciently by using a memory pool (Resource pool).
i.e. every go routine will ask for some memory from a memory pool, and when the task is completed, it will put the memory back into the pool.
If there is no empty memory in the pool, only then a new unit of memory will be created.

This can be done using sync.Pool
*/
func main() {
	var memoryUnits int
	// Initialising a memory pool
	memoryPool := &sync.Pool{
		// When a new memory is requested, increase the number of units and make some memory and return it.
		New: func() interface{} {
			memoryUnits++
			memory := make([]byte, 1024)
			return &memory
		},
	}

	const numberOfRoutines = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numberOfRoutines)
	for i := 0; i < numberOfRoutines; i++ {
		// Run a routine
		go func() {
			// Only getting memory without returning to the resource pool
			// memoryPool.Get()
			// fmt.Sprintf("Aquired Memory succesfully")
			// -> 1009679 memory units are created in memory pool

			// Getting memory from memory pool and returing it back
			mem := memoryPool.Get().(*[]byte)
			// Perform some task
			fmt.Sprintf("Aquired Memory succesfully")
			// Give back the memory
			memoryPool.Put(mem)
			// -> 1661 memory units are created in memory pool
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(memoryUnits, "memory units are created in memory pool")
}
