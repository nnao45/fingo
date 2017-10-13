package fingo

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

func para(root, word string, d os.FileInfo, buff *[]byte, wg *sync.WaitGroup) {
	*buff = append(*buff, dirwalk(word, filepath.Join(root, d.Name()))...)
	defer wg.Done()
}

func FindFile(root, word string) string {
	dir, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	buff := make([]byte, 0, 1000)

	for _, d := range dir {
		wg.Add(1)
		go para(root, word, d, &buff, wg)
	}
	wg.Wait()
	return string(buff)
}

func dirwalk(word, dir string) []byte {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	paths := make([]byte, 0, 200)

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(word, filepath.Join(dir, file.Name()))...)
			continue
		}
		if strings.Contains(file.Name(), word) {
			path := filepath.Join(dir, file.Name()) + "\n"
			bp := *(*[]byte)(unsafe.Pointer(&path))
			paths = append(paths, bp...)
		}
	}
	return paths
}
