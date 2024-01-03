package pkgs

import "fmt"

func LoopFunc() {
	for i := 0; i < 3; i++ {
		go fmt.Println("Seq: ", i)
	}
}
