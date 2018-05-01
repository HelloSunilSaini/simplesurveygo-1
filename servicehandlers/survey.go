package servicehandlers

import (
	"fmt"
	"net/http"
	"simplesurveygo/dao"
)

type UserSurvayHandler struct {
}

func (p UserSurvayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserSurvayHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.Header["Token"]

	fmt.Println("TOKENS : ", tokens)
	fmt.Println("HEADERS : ", r.Header)

	surveys := dao.GetActiveSurveys()
	fmt.Println(surveys)
	return Response200OK(surveys)
}

func (p UserSurvayHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserSurvayHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
 