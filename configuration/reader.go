package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mohit83k/rest-api-generator/utils"
)

//Config : Exported configration
var Config configuration

func init() {
	readconfig("./config.json")
	validateConfig()
}

func readconfig(fileLocation string) {
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Fatal(err)
	}
}

func validateConfig() {
	checkIfAPIFilesExist()
	checkIfTLSFilesExist()
}

func checkIfAPIFilesExist() {
	var exist bool
	for n, api := range Config.ApiRoutes {
		exist, Config.ApiRoutes[n].File = utils.FileExists(Config.MockDir, api.File)
		if !exist {
			log.Fatal(api.File, " Do not exist")
		}
	}
}

func checkIfTLSFilesExist() {
	if !Config.Server.Tls {
		return
	}
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "./"
	}
	certExist, _ := utils.FileExists(currentDir, Config.Server.Certificate)
	keyExist, _ := utils.FileExists(currentDir, Config.Server.PrivateKey)
	if !(certExist && keyExist) {
		fmt.Println("Either Certificate or private key Does Not exist. Setting TLS to false")
		Config.Server.Tls = false
	}

}
