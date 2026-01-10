package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	mo "restapi/internal/models"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]mo.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = mo.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Maths",
	}
	nextID++
	teachers[nextID] = mo.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = mo.Teacher{
		ID:        nextID,
		FirstName: "Alice",
		LastName:  "Bracken",
		Class:     "12C",
		Subject:   "Chemistry",
	}
	nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")
		teacherList := make([]mo.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}
		response := struct {
			Status string       `json:"status"`
			Count  int          `json:"count"`
			Data   []mo.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Println("Not able to retrieve the id.")
		return
	}

	teacher, exist := teachers[id]

	if !exist {
		http.Error(w, "Teacher Not found!", http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(teacher)
	}
}

func postTeacherHanddler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var newTeachers []mo.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	log.Printf("%#v\n", newTeachers)

	addedTeachers := make([]mo.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string       `json:"status"`
		Count  int          `json:"count"`
		Data   []mo.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)

}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write when used with http.Error as it will flush w writer
	// w.Write([]byte("Hello, teachers route.\n"))
	// http: superfluous response.WriteHeader call from main.getTeachersHandler (server.go:99)
	log.Println("Hello, teachers route.")

	switch r.Method {
	case http.MethodGet:
		// queryHandler(r)
		// Call get handler function
		getTeachersHandler(w, r)
	case http.MethodPost:
		//w.Write([]byte("Hello Post method on Teachers route.\n"))
		// Post request Handler
		postTeacherHanddler(w, r)
	case http.MethodPut:
		w.Write([]byte("Hello Put method on Teachers route.\n"))
	case http.MethodPatch:
		w.Write([]byte("Hello Put method on Teachers route.\n"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on Teachers route.\n"))
	default:
		w.Write([]byte("Unknown HTTP Method is called."))
	}
}
