package main

import "fmt"

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func main() {
	//скрипт не запускается
	// получаю ошибку: golang.org/toolchain@v0.0.1-go1.22.windows-amd64: Get "https://proxy.golang.org/golang.org/toolchain/@v/v0.0.1-go1.22.windows-amd64.zip": proxyconnect tcp: dial tcp [::1]:80: connectex: No connection could be made because the target machine actively refused it.
	//go.mod не менял (не понял суть)
	fmt.Printf(Reverse("Hello, OTUS!"))
}
