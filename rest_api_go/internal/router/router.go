package router

import (
	"net/http"
	ha "restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ha.RootHandler)

	mux.HandleFunc("/teachers/", ha.TeachersHandler)

	mux.HandleFunc("/students/", ha.StudentsHandler)
	mux.HandleFunc("/execs/", ha.ExecsHandler)

	return mux
}
