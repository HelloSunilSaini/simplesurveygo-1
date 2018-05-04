package servicehandlers

import (
	"net/http"
	"simplesurveygo/dao"
)

type SessionHandler struct {
}

func (p SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p SessionHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.URL.Query()["Token"]

	token := tokens[0]

	user := dao.GetSessionDetails(token)
	return Response200OK(user)
}

func (p SessionHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p SessionHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
