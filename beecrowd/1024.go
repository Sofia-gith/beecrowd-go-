package beecrowd

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Criptografia() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	reader.ReadString('\n')

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')


		for len(line) > 0 && (line[len(line)-1] == '\n' || line[len(line)-1] == '\r') {
			line = line[:len(line)-1]
		}

		runes := []rune(line)


		for j, c := range runes {
			if unicode.IsLetter(c) {
				runes[j] = c + 3
			}
		}


		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}

	
		mid := len(runes) / 2
		for j := mid; j < len(runes); j++ {
			runes[j] = runes[j] - 1
		}

		fmt.Fprintln(writer, string(runes))
	}
}