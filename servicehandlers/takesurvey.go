package servicehandlers

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
	"reflect"
)

type TakeSurvayHandler struct {
}

func (p TakeSurvayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p TakeSurvayHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.URL.Query()["Token"]

	fmt.Println("TOKENS : ", tokens)
	token := tokens[0]
	user := dao.GetSessionDetails(token)
	if reflect.DeepEqual(user,dao.UserCredentials{}) {
		return SimpleBadRequest("Login Again")
	}else{
		surveys := dao.GetUserTakenSurvey(user.Username)
		fmt.Println(surveys)
		return Response200OK(surveys)
	}
}

func (p TakeSurvayHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p TakeSurvayHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	tokens, _ := r.URL.Query()["Token"]

	var cred dao.TakeSurveyPostStruct
	err = json.Unmarshal(body, &cred)
	token := tokens[0]
	user := dao.GetSessionDetails(token)
	if reflect.DeepEqual(user,dao.UserCredentials{}) {
		return SimpleBadRequest("Login Again")
	}else{
		dao.TakeSurvey(cred,user.Username)
		return Simple200OK("Your Answer has Submited")
	}
}
 