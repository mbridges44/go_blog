package main
import(
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

func ParseJSON(file string, config Server_Config) {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", configFile)
	err = json.Unmarshal(configFile, &config)
	fmt.Printf("%s", config.Html_templates)
	if err != nil {
		log.Fatal(err)
	}
}