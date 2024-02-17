package main

import "fmt"

func main() {
	for {
		var line string
		n, _ := fmt.Scanf("%s", &line)
		if 0 == n { // ввод Enter - пустая строка
			break // ввод ^D - EOF
		}
		println(line)
	}
}
