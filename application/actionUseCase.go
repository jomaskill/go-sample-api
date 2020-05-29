package actionUseCase

import (
	"TesteWeb/domain"
	dbConnection "TesteWeb/framework"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var db = dbConnection.Start()

func init(){
	action.InitialMigration()
}

func Create(writer http.ResponseWriter, request *http.Request){

	data := action.Action{}

	reqBody, err := ioutil.ReadAll(request.Body)
	if err!=nil {
		fmt.Println(writer, "")
	}

	json.Unmarshal(reqBody, &data)

	newAction, err := action.NewAction(data)
	if err != nil{
		json.NewEncoder(writer).Encode(err.Error())
	}

	db.Create(newAction)

	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(newAction)

}

func Index(writer http.ResponseWriter, request *http.Request){

	var actions []action.Action

	db.Find(&actions)

	json.NewEncoder(writer).Encode(actions)

}

func Show(writer http.ResponseWriter, request *http.Request){

	uuid := mux.Vars(request)["uuid"]
	var actions action.Action

	db.First(&actions, "uuid = ?", uuid)

	json.NewEncoder(writer).Encode(actions)
}

func Update(writer http.ResponseWriter, request *http.Request){

	var actions action.Action
	uuid := mux.Vars(request)["uuid"]
	data := make([]string, 0)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err!=nil {
		fmt.Println(writer, "")
	}
	json.Unmarshal(reqBody, &data)

	db.First(&actions, "uuid = ?", uuid)
	db.Model(&actions).Update(&data)

	json.NewEncoder(writer).Encode(actions)
}

func Destroy(writer http.ResponseWriter, request *http.Request){

	var actions action.Action
	uuid := mux.Vars(request)["uuid"]

	db.First(&actions, "uuid = ?", uuid)
	db.Delete(&actions)

	writer.WriteHeader(204)
}
