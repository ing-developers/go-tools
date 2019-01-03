/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*/
package tools

import (
	"net/http"
	"time"
)

//CleanCache Establece encabezados para evitar el cache en el navegador
func CleanCache(w http.ResponseWriter) {
	w.Header().Set("Expires: Mon", "5 Jan 1993 05:00:00 GMT")
	w.Header().Set("Last-Modified", time.Now().Format(time.RFC1123)+" GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
}
