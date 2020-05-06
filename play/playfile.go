package play

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func PlayFilePath() {
	p, err := filepath.Abs("sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("abs: ", p)
}

func ReadLineByLine() {
	f, err := os.Open("play/playfile.go")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("text: ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func ReadCsvLineByLine() {
	f, err := os.Open("dummy.csv")
	if err != nil {
		panic(err)
	}
	csvr := csv.NewReader(f)
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("record: ", record)
	}
}

func PlayWalk() {
	i := 0
	filepath.Walk("./play", func(path string, info os.FileInfo, err error) error {
		fmt.Println("path: ", path)
		i++
		if i == 3 {
			return errors.New("test error")
		}
		return nil
	})
}

func PlayFileMode() {
	f, err := os.Lstat("./play/playfile.go")
	if err != nil {
		panic(err)
	}
	m := f.Mode()
	fmt.Println("dir: ", m.IsDir())
	fmt.Println("permission: ", m.Perm())
	fmt.Println("regular: ", m.IsRegular())
}
