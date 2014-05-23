package main

import (
	"encoding/json"
	"fmt"
	"github.com/rdxiang/tangelo-server/router"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

var tangelosCollection *mgo.Collection

type Tangelo struct {
	ID        bson.ObjectId `json:"id",bson:"_id"`
	StartTime time.Time     `json:"startTime",bson:"length"`
	Length    time.Duration `json:"length",bson:"length"`
}

func (t *Tangelo) saveToDatabase() error {
	err := tangelosCollection.Insert(t)
	return err
}

func retrieveTangeloById(id bson.ObjectId) (retrievedTangelo *Tangelo, err error) {
	retrievedTangelo = new(Tangelo)
	err = tangelosCollection.FindId(id).One(retrievedTangelo)
	return retrievedTangelo, err
}
func saveNewTangeloToDatabase() (newTangelo *Tangelo, err error) {
	newTangelo = new(Tangelo)
	newTangelo.ID = bson.NewObjectId()
	newTangelo.StartTime = time.Now()
	newTangelo.Length, err = time.ParseDuration("25m")
	if err != nil {
		return newTangelo, err
	}
	err = newTangelo.saveToDatabase()
	return newTangelo, err
}

func createTangeloHandler(w http.ResponseWriter, r *http.Request) {
	var response APIResponse
	newTangelo, err := saveNewTangeloToDatabase()
	if err != nil {
		response.Error.ErrorMessage = fmt.Sprintln("There was an error saving the new Tangelo:", err)
		response.Error.ErrorCode = StatusDatabaseErrorCode
		response.setResponseCode(http.StatusInternalServerError)
		response.send(w)
	} else {
		response.Data, err = json.Marshal(newTangelo)
		response.setResponseCode(http.StatusCreated)
		response.send(w)
	}
}
func setupTangeloRoutes() {
	s := router.APIRouter.PathPrefix("/tangelos").Subrouter()
	s.HandleFunc("/create", createTangeloHandler)

}

func init() {
	tangelosCollection = db.C(TangelosCollectionName)
}
