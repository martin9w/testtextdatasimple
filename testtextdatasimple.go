package testtextdatasimple

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gookit/goutil/arrutil"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	NullString = ""
	OutVal     = "out value"
	ExpVal     = "exp value"
)

type (
	TestFuncData struct {
		Name  string   `json:"name"`
		Alias string   `json:"alias"`
		Keys  []string `json:"keys"`
		Level string   `json:"level"`
		Datas TestDatas
	}

	TestData struct {
		Name   string
		Level  string            `json:"level"`
		Params KeyValues         `json:"params"`
		In     KeyValue          `json:"in"`
		Out    map[string]string `json:"out"`
		Exp    map[string]string `json:"exp"`
	}
	KeyValue struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	TestFuncDatas []TestFuncData
	TestDatas     []TestData
	KeyValues     []KeyValue
)

var (
	jsonString = true
)

func jsonOut(jsonStr []byte, err error) string {
	if err != nil {
		log.Println("testdata: JSON parse error: ", err)
		return ""
	}
	return string(jsonStr)
}

func (a TestFuncData) String() string {
	if jsonString {
		return jsonOut(json.Marshal(a))
	}
	return fmt.Sprintf("Name: %s, Alias: %v, Keys: %v, Level: %v, Datas: %v",
		a.Level, a.Alias, a.Keys, a.Level, a.Datas)
}

func (a TestData) String() string {
	if jsonString {
		return jsonOut(json.Marshal(a))
	}
	return fmt.Sprintf("Level: %s, Params: %v, In: %v, Out: %v, Exp: %v",
		a.Level, a.Params, a.In, a.Out, a.Exp)
}

func (a KeyValue) String() string {
	if jsonString {
		return jsonOut(json.Marshal(a))
	}
	return fmt.Sprintf("Key: %s, Value: %v",
		a.Key, a.Value)
}

func init() {
}

func SetJsonString(onOff bool) {
	jsonString = onOff
}

func GetJsonString() bool {
	return jsonString
}

func GetTestDataConfig(filename string) ([]TestFuncData, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	config := []TestFuncData{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return nil, err
	}
	for iFunc := range config {
		for iData := range config[iFunc].Datas {
			if config[iFunc].Datas[iData].Name == NullString {
				config[iFunc].Datas[iData].Name = config[iFunc].Name
			}
		}
	}
	log.Debug(config)
	return config, err
}

func TestDataChecker(t *testing.T, keys []string, testData TestData) ([]string, bool) {
	check := true
	errors := []string{}
	for _, k := range keys {
		outVal, outOk := testData.Out[k]
		expVal, expOk := testData.Exp[k]
		if outOk && expOk {
			assert.Equal(t, outVal, expVal)
		} else {
			check = false
			errorMsg := "Errors in config of testData: "
			errorMsg = ""
			if !outOk {
				errorMsg += fmt.Sprintf("%s: missing %s", k, OutVal)
			}
			if !expOk {
				errorMsg += fmt.Sprintf("%s: missing %s", k, ExpVal)
			}
			errors = append(errors, errorMsg)
		}
	}
	return errors, check
}

func SetTestDataLevel(testData TestData, levelDefault string) {
	level := testData.Level
	if level == "" {
		level = levelDefault
	}
	switch level {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		if !arrutil.Contains([]string{"debug", "info"}, level) {
			log.SetLevel(log.WarnLevel)
		}
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		if !arrutil.Contains([]string{"debug"}, level) {
			log.SetLevel(log.InfoLevel)
		}
	default:
	}
}

func GetTestDatasEntry(testdatas []TestFuncData, name string) (*TestFuncData, error) {
	for _, data := range testdatas {
		if data.Name == name {
			return &data, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no testdatas found for: %s", name))
}

func GetTestDatasEntryAsAlias(testdatas []TestFuncData, name string) (*TestFuncData, error) {
	for _, data := range testdatas {
		if data.Alias != NullString {
			name = data.Alias
		}
		if data.Name == name {
			return &data, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no testdatas found for: %s", name))
}

func GetTestDataParam(testData TestData, name string) string {
	for _, param := range testData.Params {
		if param.Key == name {
			return param.Value
		}
	}
	return NullString
}

func GetTestDataIn(testData TestData) string {
	inType := testData.In.Key
	in := NullString
	switch inType {
	case "file":
		filename := testData.In.Value
		b, _ := ioutil.ReadFile(filename)
		in = string(b)
	case "string":
		in = testData.In.Value
	default:
		in = NullString
	}
	return in
}
