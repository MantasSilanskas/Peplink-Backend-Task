package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadDataFile(fileName string) error {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var data CryptoCurrencyData

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(byteValue)

	//for i := 0; i < len();i++{
	//
	//}

	return nil
}
