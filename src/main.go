package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	cfgPath := os.Args[1]

	file, err := os.Open(cfgPath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	config, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	m := map[string]interface{}{}
	json.Unmarshal(config, &m)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString(';')

	var re = regexp.MustCompile("(?m)\\s+")
	formattedSQL := re.ReplaceAllString(text, " ")

	for _, val := range m {

		v := val.(map[string]interface{})
		f := v["find"].(string)
		r := v["replace"].(string)
		formattedSQL = regexp.MustCompile(f).ReplaceAllString(formattedSQL, r)

	}

	fmt.Println(formattedSQL)
}
