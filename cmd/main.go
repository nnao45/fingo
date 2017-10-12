package main

import (
	"fmt"
	"github.com/nao4arale/fingo"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("cpu: %d\n", runtime.NumCPU())
	fingo.FindFile(os.Args[1], os.Args[2])
}
