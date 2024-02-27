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

func (r RealRetriever) Get(url string) string {
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

	r = RealRetriever{Name: "I am a RealRetriever"}
	fmt.Println(download(r))
}
