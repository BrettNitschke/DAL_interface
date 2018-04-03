package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/omidnikta/logrus"

	"git.sphere.ms/sphere/model"
	"services/dal"
	"services/encryption"
)

type (
	// UsersController represents the controller for operating on the User resource
	UsersController struct {
		session cassandra.Dal
	}
)

func NewUsersController() *UsersController {
	session, err := cassandra.GetSession()
	if err != nil {
		logrus.Errorf("Failed to connect to Cassandra: %v.", err)
	}

	return &UsersController{session}
}

func (uc *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	// Access parameters from route
	var p httprouter.Params
	if params := context.Get(r, "params"); params != nil {
		p = params.(httprouter.Params)
	} else {
		logrus.Error("Failed to get URL parameter(s).")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "no url parameters")
		return
	}

	// Get userId from parameters
	userId, err := gocql.ParseUUID(p.ByName("UserId"))
	if err != nil {
		logrus.Errorf("Failed to parse UserId parameter: %v.", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "error")
		return
	}

	//DATABASE CALL
	uc.session.GetUserByUuid(userId)

	//DO rest of stuff here blah blah
}
