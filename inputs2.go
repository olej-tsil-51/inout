package main

import "strings"
import "bufio"
import "io"
import "os"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			println()        // ввод ^D - EOF
			break
		}
		line = strings.TrimRight(line, " \n"); 
		if 0 == len(line) { // ввод Enter - пустая строка
			break
		}
		println(line)
	}
}
