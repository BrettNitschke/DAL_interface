package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gocql/gocql"

  	"git.sphere.ms/sphere/model/modelutil"
)

type mockSession struct{}

func (ms *mockSession) GetUser(userId gocql.uuid) (model.User){

  //Make fake user data return here
}

func TestGetUser(t *testing.T){
  rec := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/getUser", nil)

  uc := UsersController{session: mockSession}
  http.HandlerFunc(uc.GetUser).ServeHTTP(rec, req)

  expected := //Whatever GetUser above returns
  if expected != rec.Body {
    t.Errorf(//error message here)
  }
}
