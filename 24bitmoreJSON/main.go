package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello World!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React", 1000, "Abc.in", "xyz123", []string{"web-dev", "js"}},
		{"PHP", 1200, "PHP.in", "def123", []string{"web-dev", "PHP"}},
		{"Go", 1000, "Go.in", "pqr123", nil},
	}

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	/*
		MarshalIndent is used to indent the JSON
		(Interface to be encoded, prefix to add to every line, spacing)
	*/
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		{
                "coursename": "React",
                "Price": 1000,
                "website": "Abc.in",
                "tags": ["web-dev","js"]
		} 
	`)
	var Course course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonData, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("JSON Invalid")
	}

	var newData map[string]interface{}
	json.Unmarshal(jsonData, &newData)

	for k, v := range newData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
