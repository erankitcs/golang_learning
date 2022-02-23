package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wahts your name ?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
	fmt.Printf("We are using Go %v running in %v", runtime.Version(), runtime.GOOS)

}
