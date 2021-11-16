package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/smstong/hello-go/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, Go"))
	fmt.Println(cmp.Diff("Hello world", "Hello go"))
}
