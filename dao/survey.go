package dao

import (
	"gopkg.in/mgo.v2/bson"
)

type Discription struct{
	Question string `json:"question"`
	Options []string `json:"options"`
}
type Survey struct {
	Id bson.ObjectId `bson:"_id" json:"_id" ,omitempty`
	Title  string `json:"title"`
	Status string `json:"status"`
	Discription []Discription `json:"discription"`
}

type SurveyPostStruct struct{
	Title string
	Status string
	Discription []Discription
}

func GetActiveSurveys() []Survey{
	session := MgoSession.Clone()
	defer session.Close()

	result := []Survey{}
	clctn := session.DB("simplesurveys").C("survey")
	clctn.Find(bson.M{"status":"active"}).All(&result)
	return result
}

func CreateSurvey(t SurveyPostStruct) {
	session := MgoSession.Clone()
	defer session.Close()
	clctn := session.DB("simplesurveys").C("survey")
	id := bson.NewObjectId()
	survey_post_var := Survey{
		Id : id,
		Title : t.Title,
		Status : t.Status,
		Discription : t.Discription,
	}
	clctn.Insert(survey_post_var)
}

type TakeSurveyStruct struct {
	Id bson.ObjectId `bson:"_id" json:"_id" ,omitempty`
	SurveyId string `json:"surveyid"`
	User string `json:"user"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}

type TakeSurveyPostStruct struct {
	SurveyId string
	Question string 
	Answer string 
}

func TakeSurvey(t TakeSurveyPostStruct,user string) {
	session := MgoSession.Clone()
	defer session.Close()
	clctn := session.DB("simplesurveys").C("takesurvey")
	id := bson.NewObjectId()
	take_survey_var := TakeSurveyStruct{
		Id : id,
		SurveyId : t.SurveyId,
		User : user,
		Question : t.Question,
		Answer : t.Answer,
	}
	clctn.Insert(take_survey_var)
}

func GetUserTakenSurvey(user string) []TakeSurveyStruct {
	session := MgoSession.Clone()
	defer session.Close()
	clctn := session.DB("simplesurveys").C("takesurvey")
	result := []TakeSurveyStruct{}
	clctn.Find(bson.M{"user":user}).All(&result)
	return result
}