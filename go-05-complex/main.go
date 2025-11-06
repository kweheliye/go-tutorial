package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	if w, ok := words["maalin"]; ok {
		log.Printf("words[\"maalin\"] = %d\n", w)
	}

}
