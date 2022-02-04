package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	dur := ParseDuration("15h")

	t := AfterFunc(dur, ban())

	dur := ParseDuration("9h")

	t := AfterFunc(dur, unban())

	//	ban()

	//	unban()

}
