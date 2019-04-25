/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*/
package go_tools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

//DecodeRequest mapea una peticion con body json a un struct
func DecodeBody(r *http.Request, modelo interface{}) error {
	jsn, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(jsn, &modelo)
	}
	return err
}

//DecodeRequest mapea una peticion con body json a un struct
func DecodeParams(r *http.Request, modelo interface{}) error {
	params := r.URL.Query().Encode()
	replacer := strings.NewReplacer("&", "\",\"", "=", "\":\"")
	return json.Unmarshal([]byte(`{"`+replacer.Replace(params)+`"}`), &modelo)
}
