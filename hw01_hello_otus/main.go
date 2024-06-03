package main

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	// Place your code here.
}
