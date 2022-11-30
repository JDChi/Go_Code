package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	var myConfig MyConfig
	var filePath = "./src/yaml/hello.yaml"
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Stat err is %v", err)
	}
	config, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ReadFile err is %v", err)
	}

	err = yaml.Unmarshal(config, &myConfig)
	if err != nil {
		fmt.Printf("Unmarshal err is %v\n", err)
	}
	fmt.Printf("MyConfig is %v\n", myConfig)

}

type MyConfig struct {
	MyNumber int64    `yaml:"myNumber"`
	MyArray  []string `yaml:"myArray"`
	MyStruct MyStruct `yaml:"myStruct"`
}

type MyStruct struct {
	Name string `yaml:"name"`
	Age  int8   `yaml:"age"`
}
