package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortParam(param string) bool {
	return param == "first_name" || param == "last_name" || param == "email" || param == "class" || param == "subject"
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

		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"

		var args []interface{}

		query, args = addFilter(r, query, args)

		// .Get will return single value i.e. string. If multiple values are there then first value will be reurned.
		// without .Get all the values will be returned and will be of type slice.
		sortParams := r.URL.Query()["sortby"]

		query = addSorting(sortParams, query)
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

func addSorting(sortParams []string, query string) string {
	if len(sortParams) > 0 {
		query += " ORDER BY "
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			field, order := parts[0], parts[1]
			if !isValidSortOrder(order) || !isValidSortParam(field) {
				continue
			}
			if i > 0 {
				query += ","
			}
			query += " " + field + " " + order
		}
	}
	return query
}

func addFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += fmt.Sprintf(" AND %s = ?", dbField)
			args = append(args, value)
		}
	}
	return query, args
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
