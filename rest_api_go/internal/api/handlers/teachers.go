package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	mo "restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
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

	db, err := sqlconnect.ConnectDB()

	if err != nil {
		http.Error(w, "Can not connect to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"

		var args []interface{}

		if firstName != "" {
			query += " AND first_name = ?"
			args = append(args, firstName)
		}

		if lastName != "" {
			query += " AND last_name = ?"
			args = append(args, lastName)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "Database query Error.", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		teacherList := make([]mo.Teacher, 0)

		for rows.Next() {
			var teacher mo.Teacher
			err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error Scanning Database Results", http.StatusInternalServerError)
				return
			}
			teacherList = append(teacherList, teacher)
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

		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Println("Not able to retrieve the id.")
		return
	}

	var teacher mo.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Database query Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(teacher)
}

func postTeacherHanddler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	db, err := sqlconnect.ConnectDB()

	if err != nil {
		http.Error(w, "Can not connect to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var newTeachers []mo.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES(?,?,?,?,?)")
	if err != nil {
		http.Error(w, "Error in Preparing SQL Query", http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	log.Printf("%#v\n", newTeachers)

	addedTeachers := make([]mo.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting into Database", http.StatusInternalServerError)
			return
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting Last inserted ID", http.StatusInternalServerError)
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
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
