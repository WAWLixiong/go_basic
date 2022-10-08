package learn_json

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func LearnJson() {
	p := Person{
		Name:  "tom",
		Age:   18,
		Email: "163.com",
	}

	b, _ := json.Marshal(p)
	fmt.Println(string(b))

	var person = Person{}
	err := json.Unmarshal(b, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
}

func LearnJson2() {

	b := []byte(`{"name":"tom","age":18,"email":"163.com"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Println(p)
}

func LearnJson3() {
	// 解析嵌json
	b := []byte(`{"name":"tom","age":18,"email":"163.com", "parents":["big tom", "kite"]}`)
	var p interface{}
	json.Unmarshal(b, &p)
	fmt.Println(p)
}

func LearnJson4() {
	// Encoder, Decoder
	// io 流可扩展到 http websocket等场景
	f, _ := os.Open("a.json")
	defer f.Close()

	decoder := json.NewDecoder(f)
	encoder := json.NewEncoder(os.Stdout)

	for { // 看实际情况是否需要 for循环
		var v map[string]interface{}
		if err := decoder.Decode(&v); err != nil {
			// fmt.Println(err) // 输出EOF
			return
		}
		fmt.Println(v)

		if err := encoder.Encode(&v); err != nil {
			fmt.Println(err)
			return
		}
	}

}
