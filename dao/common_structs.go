package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Discription struct{
	Question string `json:"question"`
	Options []string `json:"options"`
}
type Survey struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	discription []Discription `json:"discription"`
}

func GetActiveSurveys ()[]Survey{
	session := MgoSession.Clone()
	defer session.Close()

	result := []Survey{}
	
	clctn := session.DB("simplesurveys").C("survey")
	clctn.Find(bson.M{"status":"active"}).All(&result)
	return result
}

func AuthenticateUser(cred UserCredentials) string {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)

	uuidStr := uuid.Must(uuid.NewV4()).String()
	sessionStruct := Session{cred.Username, uuidStr}
	if err != nil {
		return ""
	}

	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return uuidStr
}

func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}


func CreateUser(cred UserCredentials){
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("user")
	clctn.Insert(cred)
}