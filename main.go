package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func writeAscendent(arr []string) {
	file, err := os.Create("ascendente.txt")

	if err != nil {
		fmt.Println("-------ERROR WHILE WRITING FILE-------")
		fmt.Println(err)
		return
	}
	defer file.Close()

	text := strings.Join(arr, "\n")
	bytes := []byte(text)
	file.Write(bytes)
	fmt.Printf("[LOG]	Wrote %d bytes to ascendente.txt\n", len(bytes))
}

func writeDescendent(arr []string) {
	file, err := os.Create("descendente.txt")

	if err != nil {
		fmt.Println("-------ERROR WHILE WRITING FILE-------")
		fmt.Println(err)
		return
	}
	defer file.Close()

	text := strings.Join(arr, "\n")
	bytes := []byte(text)
	file.Write(bytes)
	fmt.Printf("[LOG]	Wrote %d bytes to descendeente.txt\n", len(bytes))
}

func main() {
	var lines int
	var strsAsc, strsDesc []string
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Escribe el numero de lineas a generar: ")
	fmt.Scanln(&lines)

	for i := 0; i < lines; i++ {
		fmt.Printf("Escribe la linea: ")
		scanner.Scan()
		input = scanner.Text()
		strsAsc = append(strsAsc, input)
	}
	strsDesc = make([]string, len(strsAsc))
	copy(strsDesc, strsAsc)

	sort.Strings(strsAsc)
	sort.Sort(sort.Reverse(sort.StringSlice(strsDesc)))

	writeAscendent(strsAsc)
	writeDescendent(strsDesc)
}
