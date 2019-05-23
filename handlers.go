package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Project struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
	Period      string `json:"period,omitempty"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string
		Password string
	}
	var response struct {
		Username string `json:"username,omitempty"`
		Password string `json:"-"`
	}
	var err error

	defer func() {
		r.Body.Close()
		if err != nil {
			io.WriteString(w, NewAppErr(err).String())
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}()

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	err = DB().Select("username", "password").From("users").
		Where("email=?", request.Email).LoadOne(&response)

	if err != nil {
		return
	}
	if !ComparePassword(response.Password, request.Password) {
		err = errors.New("password is wrong")
		return
	}
}

func HandleRegsiter(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string
		Email    string
		Password string
	}
	var response struct {
		Errcode int `json:"errcode"`
	}
	var err error

	defer func() {
		r.Body.Close()
		if err != nil {
			io.WriteString(w, NewAppErr(err).String())
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}()

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	count, _ := DB().Select("COUNT(*)").From("users").
		Where("email=?", request.Email).ReturnInt64()
	if count > 0 {
		err = errors.New("Email already exists")
		return
	}

	password, err := HashPassword(request.Password)
	if err != nil {
		return
	}

	_, err = DB().InsertInto("users").
		Columns("username", "email", "password", "created").
		Values(request.Username, request.Email, password, time.Now()).Exec()
}

func HandleProjects(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Page    uint64
		PerPage uint64
	}
	var response struct {
		Total    int64     `json:"total,omitempty"`
		Projects []Project `json:"projects,omitempty"`
	}

	var err error

	defer func() {
		r.Body.Close()
		if err != nil {
			io.WriteString(w, NewAppErr(err).String())
		} else {
			data, _ := json.Marshal(response)
			w.Write(data)
		}
	}()

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	if request.PerPage == 0 || request.PerPage > 100 {
		request.PerPage = 10
	}

	response.Total, _ = DB().Select("COUNT(*)").From("project").ReturnInt64()
	_, err = DB().Select("*").From("project").
		OrderDesc("created").
		Paginate(request.Page, request.PerPage).Load(&response.Projects)
}
