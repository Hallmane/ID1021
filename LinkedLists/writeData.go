package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func writeDatShit(fileName string, arraySize int, time int) {
	//fo, err := os.Create(fileName + ".dat")
	fo, err := os.OpenFile(fileName+".dat", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(fo)
	_, err = fmt.Fprintf(w, "%d %d\n", arraySize, time)
	w.Flush()
}
