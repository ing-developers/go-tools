/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*/
package tools

import (
	"encoding/json"
	"os"
)

//Decode mapea un archivo json a un struct
func Decode(ruta string, modelo interface{}) error {
	var js *os.File
	js, err := os.Open(ruta)
	if err == nil {
		defer js.Close()
		deco := json.NewDecoder(js)
		err = deco.Decode(&modelo)
	}
	return err
}
