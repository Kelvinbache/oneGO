package otros

import (
	"embed"
	"io/fs"
	"net/http"
)

var conte embed.FS

func repuesta() http.Handler {
	file := fs.FS(conte)
	html, _ := fs.Sub(file, "./public")

	return http.FileServer(http.FS(html))
}

func conexion() {
	mux := http.NewServeMux()
	mux.Handle("/", repuesta())
	http.ListenAndServe(":8080", nil)
}

/*
 buscar un modo de conentar los archivos publicos con el servidor
*/
