package testtextdatasimple

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	TestJsonFile = "testdata/test.json"
)

var (
	testFilename  = "testdata/test.txt"
	testFuncDatas = TestFuncDatas{}
	testID        = NullString
)

func init() {
	datas, err := GetTestDataConfig(TestJsonFile)
	if err != nil {
		log.Fatal(datas, err)
	}
	testFuncDatas = datas
}

func TestFuncId(t *testing.T) {
	setJson := GetJsonString()
	testSwitch(t, "TestFuncId")
	if log.IsLevelEnabled(log.DebugLevel) {
		SetJsonString(true)
		fmt.Println(testFuncDatas)
		SetJsonString(setJson)
	}
}

func TestFuncAlias(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	setJson := GetJsonString()
	testFuncData, error := GetTestDatasEntry(testFuncDatas, "TestFuncAlias")
	if error != nil || testFuncData.Alias == NullString {
		log.Error("NoAlias")
	}
	testSwitch(t, testFuncData.Alias)
	if log.IsLevelEnabled(log.DebugLevel) {
		SetJsonString(true)
		fmt.Println(testFuncDatas)
		SetJsonString(setJson)
	}
}

func TestFuncNoAlias(t *testing.T) {
	testFuncData, error := GetTestDatasEntry(testFuncDatas, "TestFuncNoAlias")
	if error != nil {
		log.Error("Error getting Alias")
	}
	testSwitch(t, "TestFuncNoAlias")
	if testFuncData.Alias != NullString {
		log.Error("No Alias but found one")
	}
	log.Debug("OK - No Alias was set")
}

func TestUseOfOut(t *testing.T) {
	testSwitch(t, "TestUseOfOut")
}

func testChecker(t *testing.T, keys []string, testData TestData, found bool, name string) {
	if !found {
		t.Error("nothing found for: ", name)
	}
	errors, check := TestDataChecker(t, keys, testData)
	if !check {
		t.Errorf("%v", errors)
	}

}
func testSwitch(t *testing.T, name string) {
	testFuncData, error := GetTestDatasEntry(testFuncDatas, name)
	if error != nil {
		t.Error(error)
		return
	}
	if len(testFuncData.Datas) == 0 {
		SetTestDataLevel(TestData{}, testFuncData.Level)
	} else {
		for _, testData := range testFuncData.Datas {
			SetTestDataLevel(testData, testFuncData.Level)
			foundData := false
			switch name {
			case "TestFuncAlias", "TestFuncId":
				idParam := GetTestDataParam(testData, "id")
				id := idParam
				found := id != NullString
				if found {
					testData.Out = map[string]string{
						"id": id,
						"in": "InValue",
					}
				}
				foundData = found
			case "TestFuncNoAlias":
			case "TestUseOfOut":
				testData.Out = map[string]string{
					"id":   "OutId",
					"name": "OutName",
				}
				foundData = true
			}
			testChecker(t, testFuncData.Keys, testData, foundData, name)
		}
	}
}
