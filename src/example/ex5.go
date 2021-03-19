spackage main

import (
	"fmt"
)

func main() {
	// for 문
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	// 조건식만 쓰는 for 문
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)
	fmt.Println()

	// for range 문
	names := []string{"빌 게이츠", "스티브 잡스", "제프 베조스"}
	for idx, name := range names {
		fmt.Println(idx, name)
	}
	fmt.Println()

	// break, continue, goto
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}

	fmt.Println(a)
	if a == 11 {
		goto END
	}
	fmt.Println(a)

END:
	fmt.Println("End\n")
}
