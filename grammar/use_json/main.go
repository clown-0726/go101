package main

import "encoding/json"

type Human struct {
	Name string
	Age  int
}

// valid 的字符串转成结构体
func unmarshal2Struct(humanStr string) Human {
	h := Human{}

	err := json.Unmarshal([]byte(humanStr), &h)
	if err != nil {
		println(err)
	}

	return h
}

// 结构体转成 json 字符串
func marshal2JsonString(h Human) string {
	h.Age = 30
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		return ""
	}
	return string(updatedBytes)
}

func main() {
	h := Human{
		Name: "Zhangsan",
		Age:  18,
	}
	res := marshal2JsonString(h)
	println(res)
}
