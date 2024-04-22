package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   apple,banana,grape,jackfruit,mango     "
	splt_str := strings.Split(str, ",")  // splitting
	count_str := strings.Count(str, "s") // count the number of time a ch or set of ch is present in there
	trimSpace_str := strings.TrimSpace(str)
	fmt.Println(splt_str, "\n", count_str, "\n", trimSpace_str)
}
