package main
import(
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseJSON(file string, config interface{}) {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)

	if err != nil {
		log.Fatal(err)
	}
}