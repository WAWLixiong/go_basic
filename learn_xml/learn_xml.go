package learn_xml

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func LearnXml() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "163.com",
	}

	// b, _ := xml.Marshal(p)
	// 有缩进格式
	b, _ := xml.MarshalIndent(p, " ", " ")
	fmt.Println(string(b))

	var m Person
	xml.Unmarshal(b, &m)
	fmt.Println(m)
}

func LearnXml2() {
	f, err := os.Open("a.xml")
	if err != nil {
		log.Panicln(err)
		return
	}
	encoder := xml.NewDecoder(f)
	var p Person
	err = encoder.Decode(&p)
	if err != nil {
		log.Panicln(err)
		return
	}
	fmt.Println(p)

}
