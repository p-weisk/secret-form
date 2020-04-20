package rest

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/p-weisk/secret-form/answer"
	"github.com/p-weisk/secret-form/form"
	"github.com/p-weisk/secret-form/result"
	"io"
	"log"
	"net/http"
	"time"
)

func NewForm(w http.ResponseWriter, r *http.Request) {
	//	allowCors(w)
	f, err := createFormFromJson(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%+v : Failed to parse form from json request body: %+v", time.Now(), err)
		return
	}
	form, ferr := form.CreateForm(f.Content, f.PublicKey)
	if ferr != nil {
		http.Error(w, "An error occured while trying to create purchase", http.StatusInternalServerError)
		log.Println(ferr.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	setJsonContentType(w)
	fmt.Fprintf(w, "{\"id\":\"%s\", \"location\":\"/api/form/%s\"}", form.ID.String(), form.ID.String())
	log.Printf("%+v : HTTP 201/CREATED %s on POST /form", form.ID.String(),time.Now())
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	//	allowCors(w)
	idStr := mux.Vars(r)["fid"]
	id, ierr := uuid.Parse(idStr)
	if ierr != nil {
		http.Error(w, ierr.Error(), http.StatusNotFound)
		log.Println(ierr.Error)
		return
	}
	form, ferr := form.GetForm(id)
	if ferr != nil {
		log.Println(ferr.Error)
		http.Error(w, ferr.Error(), http.StatusInternalServerError)
		return
	}
	serr := sendJsonResponse(w, form)
	if serr != nil {
		log.Println(serr.Error())
		return
	}
	log.Printf("%+v : HTTP 200/OK on GET /form/%s", time.Now(), idStr)
}

func AddAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["fid"]
	id, ierr := uuid.Parse(idStr)
	if ierr != nil {
		http.Error(w, ierr.Error(), http.StatusNotFound)
		log.Println(ierr.Error)
		return
	}
	a, err := createAnswerFromJson(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%+v : Failed to parse answer from json request body: %+v", time.Now(), err)
		return
	}
	aerr := answer.CreateAnswer(a.Data, id)
	if aerr != nil {
		http.Error(w, "An error occured while trying to create answer", http.StatusInternalServerError)
		log.Println(aerr.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Printf("%+v : HTTP 201/CREATED on POST /form/%s/answers", idStr, time.Now())
}

func FormResult(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["fid"]
	id, ierr := uuid.Parse(idStr)
	if ierr != nil {
		http.Error(w, ierr.Error(), http.StatusNotFound)
		log.Println(ierr.Error)
		return
	}
	result, rerr := result.GetResult(id)
	if rerr != nil {
		log.Println(rerr.Error)
		http.Error(w, rerr.Error(), http.StatusInternalServerError)
		return
	}
	serr := sendJsonResponse(w, result)
	if serr != nil {
		log.Println(serr.Error())
		return
	}
	log.Printf("%+v : HTTP 200/OK on GET /form/%s/result", time.Now(), idStr)
}

func createFormFromJson(r io.ReadCloser) (form.Form, error) {
	decoder := json.NewDecoder(r)
	f := form.Form{}
	err := decoder.Decode(&f)
	return f, err
}

func createAnswerFromJson(r io.ReadCloser) (answer.Answer, error) {
	decoder := json.NewDecoder(r)
	a := answer.Answer{}
	err := decoder.Decode(&a)
	return a, err
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func sendJsonResponse(w http.ResponseWriter, payload interface{}) error {
	j, jerr := json.Marshal(payload)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
		return jerr
	} else {
		setJsonContentType(w)
		w.WriteHeader(200)
		fmt.Fprint(w, string(j))
		return nil
	}
}
