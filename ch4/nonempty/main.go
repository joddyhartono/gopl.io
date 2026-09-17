package main

import "fmt"

/* pointer masih mengarah ke underlying array-nya masih sama, oleh karena itu array-nya dapat berubah */
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/* out mengarah ke underlying array strings, dengan len = 0 */
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // one, three
	fmt.Printf("%q\n", data) // one three three
}