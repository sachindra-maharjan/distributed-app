package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type studentsHandler struct{}

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

// /students - all students
// /students/{id} - single student
// /students/{id}/grades - grade of a student
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.Split(r.URL.Path, "/")
	switch len(pathSegment) {
	case 2:
		sh.getAll(w, r)
	case 3:
		id, err := strconv.Atoi((pathSegment[2]))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi((pathSegment[2]))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (sh studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentMutex.Lock()
	defer studentMutex.Unlock()

	data, err := sh.toJson(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) toJson(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize students: %q", err)
	}
	return b.Bytes(), nil
}

func (sh studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentMutex.Lock()
	defer studentMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data, err := sh.toJson(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize students: %q", err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentMutex.Lock()
	defer studentMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var g Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	student.Grades = append(student.Grades, g)

	w.WriteHeader(http.StatusCreated)
	data, err := sh.toJson(g)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
