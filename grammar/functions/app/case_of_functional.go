package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ------------------------------------------------------------------------------------------

func MainInvokeAdder() {
	ad := adder()
	fmt.Println(ad(1))
	fmt.Println(ad(2))
	fmt.Println(ad(3))
	fmt.Println(ad(4))
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// ------------------------------------------------------------------------------------------

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func MainInvokeAdder2() {
	ad := adder2(0)
	res, ad := ad(1)
	res, ad = ad(2)
	res, ad = ad(3)
	res, ad = ad(4)
	fmt.Println(res)
}

// ------------------------------------------------------------------------------------------

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

// 这个函数接收一个 IO 流，通过 bufio.NewScanner 循环输出内容
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func MainFibo() {
	f := fibonacci()
	printFileContents(f)
}
