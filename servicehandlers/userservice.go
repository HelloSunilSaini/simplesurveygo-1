package servicehandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
)

type UserServiceHandler struct {
}

func (p UserServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserServiceHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserServiceHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserServiceHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var cred dao.UserCredentials
	err = json.Unmarshal(body, &cred)
	token := dao.AuthenticateUser(cred)

	if token == "" {
		dao.CreateUser(cred)
		return Simple200OK("Registration Successful")
	} else {
		return SimpleBadRequest("User Already Exist")
	}
}
 