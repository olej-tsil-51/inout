package main

import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // ввод ^D - EOF
		line := scanner.Text()
		if 0 == len(line) {
			break // ввод Enter - пустая строка
		}
		println(line)
	}
}
