package servicehandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
	"reflect"
)

type UserSurvayHandler struct {
}

func (p UserSurvayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserSurvayHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.URL.Query()["Token"]

	token := tokens[0]
	user := dao.GetSessionDetails(token)
	if reflect.DeepEqual(user,dao.UserCredentials{}) {
		return SimpleBadRequest("Login Again")
	}else{
		surveys := dao.GetActiveSurveys()
		return Response200OK(surveys)
	}
}

func (p UserSurvayHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserSurvayHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	tokens, _ := r.URL.Query()["Token"]
	var cred dao.SurveyPostStruct
	err = json.Unmarshal(body, &cred)

	token := tokens[0]
	user := dao.GetSessionDetails(token)
	if reflect.DeepEqual(user,dao.UserCredentials{}) {
		return SimpleBadRequest("Login Again")
	}else{
		dao.CreateSurvey(cred)
		return Simple200OK("Survey Created")
	}
}
 