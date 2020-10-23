package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/asdine/storm"
	"github.com/gorilla/mux"
)

// Building such as home or garage
type Building struct {
	IDBuilding    int    `storm:"id,increment"`
	BuildingName  string `storm:"index"`
	BuildingOrder int
}

// Buildings is a collection of buildings
type Buildings struct {
	Building []Building
}

func BuildingList(w http.ResponseWriter, r *http.Request) {
	db, _ := storm.Open("database.db")
	defer db.Close()
	var buildings []Building
	err := db.All(&buildings)
	if err != nil {
		log.Println(err)
	}
	tmpl.ExecuteTemplate(w, "BuildingList", buildings)
}

func BuildingAdd(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "BuildingAdd", "")
}
func BuildingDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])

	db, _ := storm.Open("database.db")
	defer db.Close()

	var building Building
	err := db.One("IDBuilding", i, &building)

	if err != nil {
		log.Println(err)
	}
	tmpl.ExecuteTemplate(w, "BuildingDelete", building)
}
func BuildingRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])

	db, _ := storm.Open("database.db")
	defer db.Close()

	var building Building
	err := db.One("IDBuilding", i, &building)
	if err != nil {
		log.Println(err)
	}
	db.DeleteStruct(&building)
	http.Redirect(w, r, "/ui/manage/buildings", 301)
}

func BuildingEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])

	db, _ := storm.Open("database.db")
	defer db.Close()

	var building Building
	err := db.One("IDBuilding", i, &building)

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "BuildingEdit", building)
}
func BuildingUpdate(w http.ResponseWriter, r *http.Request) {

	db, _ := storm.Open("database.db")
	defer db.Close()
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("idbuilding"))
		var building Building
		err := db.One("IDBuilding", i, &building)
		db.UpdateField(&Building{IDBuilding: i}, "BuildingName", r.FormValue("buildingname"))
		if err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/ui/manage/buildings", 301)
}

func BuildingInsert(w http.ResponseWriter, r *http.Request) {
	db, _ := storm.Open("database.db")
	defer db.Close()
	if r.Method == "POST" {
		b := Building{
			BuildingName:  r.FormValue("buildingname"),
			BuildingOrder: 1,
		}
		err := db.Save(&b)
		if err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/ui/manage/buildings", 301)
}
