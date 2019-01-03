package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

//Responder escribe en el response la respuesta establecida
func Responder(w http.ResponseWriter, estado bool, error interface{}, mensaje string, datos interface{}, esJSON bool, debug bool) {
	resp := struct {
		Estado  bool
		Error   interface{}
		Mensaje string
		Datos   interface{}
	}{
		estado,
		error,
		mensaje,
		datos,
	}
	if !debug && esJSON {
		resp.Error = nil
	}
	if esJSON {
		js, err := json.Marshal(resp)
		if err == nil {
			CleanCache(w)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			log.Println("Error Marshal JSON: ", err)
		}
	} else {
		w.Write([]byte(mensaje))
	}
}
