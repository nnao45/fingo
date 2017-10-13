# fingo
find cmd is writtend in golang, moving parallel processes.
## usage
```bash
package main

import (
	"fmt"
	"github.com/nao4arale/fingo"
	"os"
	"runtime"
)

func main() {
  /* os.Args[1]...Dirctory, os.Args[2]...To find Words. */
	fmt.Printf("%s", fingo.FindFile(os.Args[1], os.Args[2]))
}
```
