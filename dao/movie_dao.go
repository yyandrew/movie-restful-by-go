package dao

import (
	"log"

	. "github.com/lq201/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MovieDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MovieDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
func (m *MovieDao) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}
func (m *MovieDao) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MovieDao) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MovieDao) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *MovieDao) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
