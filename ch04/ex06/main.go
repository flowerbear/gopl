// ch04/ex06 は、隣接している Unicode スペースを 1 つの ASCII スペースに圧縮します。
package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "1  +  1     =  2"
	fmt.Println(string(removeDupSpace([]byte(s)))) // "1 + 1 = 2"
}

// removeDupSpace は、隣接している Unicode スペースを 1 つの ASCII スペースに圧縮します。
func removeDupSpace(b []byte) []byte {
	var buf bytes.Buffer
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(string(b[i:]))
		if unicode.IsSpace(r) {
			nextrune, _ := utf8.DecodeRuneInString(string(b[i+size:]))
			if !unicode.IsSpace(nextrune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		i += size
	}
	return buf.Bytes()
}
