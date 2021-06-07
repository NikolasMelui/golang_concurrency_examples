package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Result struct {
	Sdk      map[string]string
	SdkChan  chan string
	Env      map[string]string
	EnvChan  chan string
	Lang     map[string]string
	LangChan chan string
	Done     int
	DoneChan chan bool
}

const (
	FILE_PATH = "./data/result.json"
)

func main() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	result := &Result{
		Sdk:      make(map[string]string),
		SdkChan:  make(chan string),
		Env:      make(map[string]string),
		EnvChan:  make(chan string),
		Lang:     make(map[string]string),
		LangChan: make(chan string),
		DoneChan: make(chan bool),
	}

	byteValue, _ := ioutil.ReadAll(file)
	data := &[]interface{}{}
	json.Unmarshal([]byte(byteValue), data)

	for _, element := range *data {
		celem := element.(map[string]interface{})
		go result.parseMap(&celem, 1)
	}

	for result.Done < len(*data) {
		select {
		case key := <-result.EnvChan:
			result.Env[key] = ""
		case key := <-result.LangChan:
			result.Lang[key] = ""
		case key := <-result.SdkChan:
			result.Sdk[key] = ""
		case <-result.DoneChan:
			result.Done++
		}
	}

	result.print()

}

func (r *Result) print() {
	fmt.Println("Sdk - ", r.Sdk)
	fmt.Println("Env - ", r.Env)
	fmt.Println("Lang - ", r.Lang)
}

func (r *Result) parseMap(element *map[string]interface{}, counter int) {
	if (*element)["children"] != nil {
		celem := (*element)["children"].([]interface{})
		r.parseArray(&celem, counter)
	}
	if (*element)["content"] != nil {
		celem := (*element)["content"].([]interface{})
		r.parseContent(&celem)
	}
	if counter == 1 {
		r.DoneChan <- true
	}
}

func (r *Result) parseArray(element *[]interface{}, counter int) {
	for _, elem := range *element {
		counter++
		celem := elem.(map[string]interface{})
		r.parseMap(&celem, counter)
	}
}

func (r *Result) parseContent(element *[]interface{}) {
	for _, elem := range *element {
		celem := elem.(map[string]interface{})
		if celem["kind"] == "content_source" {
			r.EnvChan <- celem["env"].(string)
			for _, exam := range celem["examples"].([]interface{}) {
				cexam := exam.(map[string]interface{})
				if cexam["title"] == nil {
					r.LangChan <- cexam["lang"].(string)
				} else {
					r.SdkChan <- cexam["title"].(string)
					for _, sdkExam := range cexam["examples"].([]interface{}) {
						r.LangChan <- sdkExam.(map[string]interface{})["lang"].(string)
					}
				}
			}
		}
	}
}
