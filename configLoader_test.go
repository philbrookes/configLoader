package configLoader_test

import (
	"testing"
	"os"
	"github.com/philbrookes/configLoader"
)

func TestLoad(t *testing.T){
	configLoader := configLoader.NewConfigLoader()

	file := os.Getenv("GOPATH") + "/src/github.com/philbrookes/configLoader/test_resources/good_config.json"

	config, err := configLoader.GetConfigFor(file)

	if err != nil {
		t.Error(err)
	}

	users := config.GetNested("users")

	if users.Size() == 0 {
		t.Error("nested users values not retrieved succesfully")
	}

	username := config.GetNested("users").GetNested("0").GetValue("name")

	if username != "Phil Brookes" {
		t.Error("Value not retrieved succesfully")
	}
}
