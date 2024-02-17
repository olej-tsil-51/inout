package main

import "os"
import "io"

func main() {
	line := make([]byte, 160)
	for {
		n, err := os.Stdin.Read(line)
		if err == io.EOF { // ввод ^D - EOF
			break
		}
		s := string(line[:n-1])
		if 0 == len(s) { // ввод Enter - пустая строка
			break
		}
		println(s)
	}
}
