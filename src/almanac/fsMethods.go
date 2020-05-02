package almanac

import (
	"github.com/hisitra/hedron/src/configs"
	"io/ioutil"
	"log"
	"os"
)

func CreateBaseDir() {
	info, err := os.Stat(configs.Storage.BaseLocation)
	if !os.IsNotExist(err) && info.IsDir() {
		log.Println("Base storage location (", configs.Storage.BaseLocation, ") found created.")
		return
	}
	log.Println("Creating base storage location (", configs.Storage.BaseLocation, ")...")
	err = os.MkdirAll(configs.Storage.BaseLocation, 0755)
	if err != nil {
		log.Fatalln("Failed to create storage directory because: " + err.Error())
	}
	log.Println("Successfully created base storage location.")
}

func readFile(path string) []byte {
	log.Println("Reading file:", path)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Failed to read file:", path, "because:", err)
		return nil
	}
	log.Println("File read successful.")
	return contents
}

func writeFile(path string, content []byte) error {
	log.Println("Writing file:", path)
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Println("Failed to write file:", path, "because:", err)
		return err
	}
	log.Println("File write successful.")
	return nil
}
