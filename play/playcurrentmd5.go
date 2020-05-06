package play

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
)

func PlayConcurrentMD5() {
	m, err := sumFiles("./play")
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		fmt.Printf("file: %s, sum: %x\n", k, v)
	}
}

func sumFiles(path string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{}, 1)
	// 1. walk files
	paths, errc := walk(path, done)
	// 2. sum files
	c := digest(paths, done)
	// 3. sort
	m := make(map[string][md5.Size]byte)
	for r := range c {
		m[r.path] = r.sum
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

func digest(paths <-chan string, done chan struct{}) <-chan result {
	c := make(chan result, 5)
	numDigester := 5
	var wg sync.WaitGroup
	wg.Add(numDigester)
	for i := 0; i < numDigester; i++ {
		go func() {
			defer wg.Done()
			for path := range paths {
				data, err := ioutil.ReadFile(path)
				select {
				case <-done:
					return
				case c <- result{path, md5.Sum(data), err}:
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func walk(path string, done <-chan struct{}) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case <-done:
				return errors.New("walk stopped")
			case paths <- path:
			}
			return nil
		})
		close(paths)
		errc <- err
	}()
	return paths, errc
}

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}
