package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("test travis ci")
	testRegexp()
}

func testRegexp() {
	comment := `/remove-lifecycle active
/lifecycle frozen
`
	lifecycleRe := regexp.MustCompile(`(?mi)^/(remove-)?lifecycle (active|frozen|stale|rotten)\s*$`)
	submatchs := lifecycleRe.FindAllStringSubmatch(comment, -1)
	fmt.Println(submatchs)
	for _, mat := range submatchs {
		fmt.Println(fmt.Sprintf("%s===%s", mat[1], mat[2]))
	}
	m := map[string]bool{}
	m["1"] = true
	m["2"] = true
	fmt.Println(m["3"], m["2"])
}
