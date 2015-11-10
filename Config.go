package configLoader

import (
	"encoding/json"
	"strconv"
)

type Config struct {
	values map[string]string
	nested map[string]*Config
}

func NewConfig() *Config {
	config := Config{
		values: make(map[string]string),
		nested: make(map[string]*Config),
	}
	return &config
}

func (this *Config) Size() int {
	return len(this.values) + len(this.nested)
}

func (this *Config) Parse(jsonData []byte) error {
	var decoded interface{}

	err := json.Unmarshal(jsonData, &decoded)
	if err != nil {
		return err
	}

	switch val := decoded.(type){
	case []interface{}:
		this.parseArr(val)
	case map[string]interface{}:
		this.parseMap(val)
	}

	return nil
}

func (this *Config) parseArr(decoded interface{}) error {
	decodedArr := decoded.([]interface{})

	for key, value := range decodedArr {
		switch v := value.(type) {
		case string:
			this.values[strconv.Itoa(key)] = v
		default:
			cnf := NewConfig()
			jsn, err := json.Marshal(v)
			if err != nil {
				return err
			}
			cnf.Parse(jsn)
			this.nested[strconv.Itoa(key)] = cnf
		}
	}
	return nil
}

func (this *Config) parseMap(decoded interface{}) error {
	decodedMap := decoded.(map[string]interface{})

	for key, value := range decodedMap {
		switch v := value.(type) {
		case string:
			this.values[key] = v
		default:
			cnf := NewConfig()
			jsn, err := json.Marshal(v)
			if err != nil {
				return err
			}
			cnf.Parse(jsn)
			this.nested[key] = cnf
		}
	}

	return nil
}

func (this *Config) GetValue(key string) string {
	 if value, ok := this.values[key]; ok {
		 return value
	 }

	return ""
}

func (this *Config) GetNested(key string) *Config {
	if value, ok := this.nested[key]; ok {
		return value
	}

	return NewConfig()
}

