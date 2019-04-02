/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*/
package go_tools

import (
	"html/template"
	"log"
	"net/http"
)

//RenderTemplate renderiza template
func RenderTemplate(w http.ResponseWriter, data interface{}, multiples bool, layout string, tpl ...string) {
	CleanCache(w)
	t, err := template.ParseFiles(tpl...)
	if err != nil {
		log.Println("Error ParseFiles: ", err)
		w.Write([]byte("Ocurrio un error al renderizar la vista solicitada"))
		return
	}

	if multiples {
		err = t.ExecuteTemplate(w, layout, data)
	} else {
		err = t.Execute(w, data)
	}

	if err != nil {
		log.Println("Error Execute / ExecuteTemplate: ", err)
		w.Write([]byte("Ocurrio un error al renderizar la vista solicitada"))
	}
}
