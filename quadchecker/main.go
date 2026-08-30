package main

import (
	"fmt"
	"os"
)

func check(s []byte, w, h int, a, b, c, d, x, y byte) bool {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			var q byte

			switch {
			case j == 0 && i == 0:
				q = a
			case j == 0 && i == w-1:
				q = b
			case j == h-1 && i == 0:
				q = c
			case j == h-1 && i == w-1:
				q = d
			case j == 0 || j == h-1:
				q = x
			case i == 0 || i == w-1:
				q = y
			default:
				q = ' '
			}

			if s[j*(w+1)+i] != q {
				return false
			}
		}
	}

	return true
}

func main() {
	s, err := os.ReadFile("/dev/stdin")
	if err != nil || len(s) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	
	w := 0
	for w < len(s) && s[w] != '\n' {
		w++
	}

	
	if w == 0 || w >= len(s) {
		fmt.Println("Not a quad function")
		return
	}

	h := 0
	pos := 0

	for pos < len(s) {
		end := pos

		for end < len(s) && s[end] != '\n' {
			end++
		}

		if end-pos != w {
			fmt.Println("Not a quad function")
			return
		}

		if end >= len(s) {
			fmt.Println("Not a quad function")
			return
		}

		h++
		pos = end + 1
	}

	if h == 0 {
		fmt.Println("Not a quad function")
		return
	}

	n := 0

	if check(s, w, h, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[quadA] [%d] [%d]", w, h)
		n++
	}

	if check(s, w, h, '/', '\\', '\\', '/', '*', '*') {
		if n > 0 {
			fmt.Print(" || ")
		}
		fmt.Printf("[quadB] [%d] [%d]", w, h)
		n++
	}

	if check(s, w, h, 'A', 'A', 'C', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		fmt.Printf("[quadC] [%d] [%d]", w, h)
		n++
	}

	if check(s, w, h, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		fmt.Printf("[quadD] [%d] [%d]", w, h)
		n++
	}

	if check(s, w, h, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		fmt.Printf("[quadE] [%d] [%d]", w, h)
		n++
	}

	if n == 0 {
		fmt.Print("Not a quad function")
	}

	fmt.Println()
}
