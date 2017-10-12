package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

func main() {
	firstDir(os.Args[1])
}

func para(root string, d os.FileInfo, buff *[]byte, wg *sync.WaitGroup) {
	*buff = append(*buff, dirwalk(filepath.Join(root, d.Name()))...)
	wg.Done()
}

func firstDir(root string) {
	dir, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	buff := make([]byte, 0, 1000)

	for _, d := range dir {
		wg.Add(1)
		go para(root, d, &buff, wg)
	}
	wg.Wait()
	fmt.Print(string(buff))
}

func dirwalk(dir string) []byte {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	paths := make([]byte, 0, 200)
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		if strings.Contains(file.Name(), os.Args[2]) {
			path := filepath.Join(dir, file.Name()) + "\n"
			bp := *(*[]byte)(unsafe.Pointer(&path))
			paths = append(paths, bp...)
		}
	}
	return paths
}
