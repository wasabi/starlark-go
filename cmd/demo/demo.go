package main

import (
	"log"

	"go.starlark.net/starlark"
)

const fibonacciStar = `
def fibonacci(n):
	res = list(range(n))
	for i in res[2:]:
		res[i] = res[i-2] + res[i-1]
	return res
`

func main() {
	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "my thread"}
	globals, err := starlark.ExecFile(thread, "fibonacci.star", fibonacciStar, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve a module global.
	fibonacci := globals["fibonacci"]
	_ = fibonacci
	// // Call Starlark function from Go.
	// v, err := starlark.Call(thread, fibonacci, starlark.Tuple{starlark.MakeInt(10)}, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("fibonacci(10) = %v\n", v) // fibonacci(10) = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
}
