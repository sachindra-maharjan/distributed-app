package grades

import "net/http"

type studentsHandler struct{}

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}
