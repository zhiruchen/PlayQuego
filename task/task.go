package task

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	que "github.com/bgentry/que-go"
)

func Add(j *que.Job) error {
	data := struct {
		A string `json:"a"`
		B string `json:"b"`
	}{}

	err := json.Unmarshal(j.Args, &data)
	if err != nil {
		return err
	}

	ioutil.WriteFile("log.txt", []byte(data.A+" "+data.B), 0666)
	return nil
}

func Mul(j *que.Job) error {
	data := struct {
		Str   string `json:"str"`
		Times int32  `json:"times"`
	}{}

	err := json.Unmarshal(j.Args, &data)
	if err != nil {
		return err
	}

	sl := []string{}
	var i int32 = 1
	for i <= data.Times {
		sl = append(sl, data.Str)
		i++
	}

	ioutil.WriteFile("log1.txt", []byte(strings.Join(sl, "***")), 0666)
	return nil
}
