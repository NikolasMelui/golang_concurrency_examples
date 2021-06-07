package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Result struct {
	Sdk  map[string]string
	Env  map[string]string
	Lang map[string]string
}

const (
	FILE_PATH = "/path/to/the/file.json"
)

func main() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	result := &Result{
		Sdk:  make(map[string]string),
		Env:  make(map[string]string),
		Lang: make(map[string]string),
	}

	byteValue, _ := ioutil.ReadAll(file)
	data := &[]interface{}{}
	json.Unmarshal([]byte(byteValue), &data)

	result.parseArray(data)

	fmt.Println("Sdk - ", result.Sdk)
	fmt.Println("Env - ", result.Env)
	fmt.Println("Lang - ", result.Lang)

}

func (r *Result) parseMap(element *map[string]interface{}) {
	if (*element)["children"] != nil {
		cel := (*element)["children"].([]interface{})
		r.parseArray(&cel)
	}
	if (*element)["content"] != nil {
		cel := (*element)["content"].([]interface{})
		r.parseContent(&cel)
	}
}

func (r *Result) parseArray(element *[]interface{}) {
	for _, item := range *element {
		cel := item.(map[string]interface{})
		r.parseMap(&cel)
	}
}

func (r *Result) parseContent(element *[]interface{}) {
	for _, elem := range *element {
		celem := elem.(map[string]interface{})
		if celem["kind"] == "content_source" {
			r.Env[celem["env"].(string)] = ""
			for _, exam := range celem["examples"].([]interface{}) {
				cexam := exam.(map[string]interface{})
				if cexam["title"] == nil {
					r.Lang[cexam["lang"].(string)] = ""
				} else {
					r.Sdk[cexam["title"].(string)] = ""
					for _, sdkExam := range cexam["examples"].([]interface{}) {
						r.Lang[sdkExam.(map[string]interface{})["lang"].(string)] = ""
					}
				}
			}
		}
	}
}
