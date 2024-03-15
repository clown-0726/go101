package app

import "fmt"

type MockRetriever struct {
	Name string
}

func (r MockRetriever) Get(url string) string {
	return "MockRetriever..."
}

type RealRetriever struct {
	Name string
}

func (r *RealRetriever) Get(url string) string {
	return "RealRetriever..."
}

// 定义接口
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://google.com/")
}

func MainWhatIsInterface() {
	var r Retriever
	r = MockRetriever{Name: "I am a MockRetriever"}
	fmt.Println(download(r))
	fmt.Printf("%T, %v\n", r, r)

	// RealRetriever 中 Get 方式是指针接收者，因此要取其指针地址
	r = &RealRetriever{Name: "I am a RealRetriever"}
	fmt.Println(download(r))
	fmt.Printf("%T, %v\n", r, r)
}
