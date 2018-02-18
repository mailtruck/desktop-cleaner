package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func main() {
	desktop := filepath.Join(os.Getenv("HOME"), "Desktop")
	// TODO configurable dir name
	bak := filepath.Join(desktop, "bak")
	if _, err := os.Stat(bak); os.IsNotExist(err) {
		err = os.Mkdir(bak, 0777)
		if err != nil {
			log.Fatal(errors.Wrap(err, "making dir bak"))
		}
	}

	files, err := filepath.Glob(filepath.Join(desktop, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file == bak {
			continue
		}
		err := os.Rename(file, filepath.Join(bak, filepath.Base(file)))
		if err != nil {
			log.Println(errors.Wrap(err, "moving file").Error())
		}
	}
}
