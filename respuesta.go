package go_tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Responder escribe en el response la respuesta establecida
func Responder(w http.ResponseWriter, estado int, error interface{}, mensaje string, datos interface{}, esJSON bool, debug bool) {
	resp := struct {
		Estado  int         `json:"estado"`
		Error   interface{} `json:"error"`
		Mensaje string      `json:"mensaje"`
		Datos   interface{} `json:"datos"`
	}{
		estado,
		error,
		mensaje,
		datos,
	}
	log.Println(fmt.Sprintf("%#v", resp))
	if esJSON {
		if !debug {
			resp.Error = nil
		}
		js, err := json.Marshal(resp)
		if err == nil {
			CleanCache(w)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(estado)
			w.Write(js)
		} else {
			log.Println("Error Marshal JSON: ", err)
		}
	} else {
		w.WriteHeader(estado)
		w.Write([]byte(mensaje))
	}

}
