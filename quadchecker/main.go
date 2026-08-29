package main

import (
	"fmt"
	"os"
)

func check(s []byte, w, h int, a, b, c, d, x, y byte) bool {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			var q byte

			if j == 0 && i == 0 {
				q = a
			} else if j == 0 && i == w-1 {
				q = b
			} else if j == h-1 && i == 0 {
				q = c
			} else if j == h-1 && i == w-1 {
				q = d
			} else if j == 0 || j == h-1 {
				q = x
			} else if i == 0 || i == w-1 {
				q = y
			} else {
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
	s, _ := os.ReadFile("/dev/stdin")

	if len(s) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	w := 0
	for w < len(s) && s[w] != '\n' {
		w++
	}

	h := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			h++
		}
	}

	if w == 0 || h == 0 || len(s) != h*(w+1) {
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