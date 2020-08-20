package util

import (
	"fmt"
	"github.com/coreos/fcct/config/v1_1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filepath)
		return false
	}
	return true
}

func UnmarshalFCC(filepath string) (*v1_1.Config, error) {
	var fcc *v1_1.Config
	fccFileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error parsing input FCC")
	}

	err = yaml.Unmarshal(fccFileContents, &fcc)
	return fcc, err
}

func MarshalAndPrintFCC(fcc *v1_1.Config, inPlace bool) {
	o, err := yaml.Marshal(&fcc)
	if err != nil {
		fmt.Println(err)
		return
	}
	if inPlace {

	}
	fmt.Printf("%s\n", string(o))
}
