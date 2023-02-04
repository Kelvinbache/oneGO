package otros

import (
	"log"
	"net/http"
)

func servidor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hola</h1>"))
}

func repuesta() {
	http.HandleFunc("/", servidor)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
