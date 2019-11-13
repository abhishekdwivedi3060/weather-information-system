package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	weat "github.com/abhishekdwivedi3060/weather-information-system/weather"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var apiKey string
var db *gorm.DB

func init() {
	var err error
	apiKey = os.Getenv("APIKEY")
	for {
		db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/weather?charset=utf8&parseTime=True")
		if err == nil {
			log.Print("Successfully connected to database")
			break
		} else {
			time.Sleep(5 * time.Second)
			log.Print("Waiting to connect to databse")
		}
	}
}

func main() {
	defer db.Close()
	db.AutoMigrate(&weat.RequiredData{})
	city := map[string]int{"Delhi": 110001, "Mumbai": 400004, "Jaipur": 302001, "Noida": 201203, "Bhopal": 462001, "Gandhinagar": 382002, "Pune": 411007, "Nasik": 422101, "Kochi": 682001,
		"Agra": 282001, "Ajmer": 305001, "Alibag": 402201, "Aligarh": 202001, "Allahabad": 211001, "Alwar": 683101, "Amaravati": 444601, "Ambala": 133003, "Amritsar": 143001, "Assam": 793011, "Aurangabad": 431001,
		"Bangalore": 560002, "Baramulla": 193101, "Belgaum": 590001, "Bhagalpur": 812001, "Bhuvaneswar": 751001, "Borivali": 400901, "KolKata": 700001, "Chandigarh": 160017, "Chittorgarh": 312001, "Coimbatore": 641001,
		"Darjeeling": 734101, "dholpur": 328001, "Durgapur": 314001, "Gangtok": 737101, "Gaya": 823001, "Ghaziabad": 201001, "Guwahati": 781001, "Gwalior": 474001, "Howrah": 711101, "Indore": 452001, "Itanagar": 791111,
		"Kanpur": 208001, "Kolhapur": 416003, "Kota": 324001, "Lucknow": 226003, "Ludhiana": 141001, "Meerut": 250001, "Mysore": 570001, "Nagpur": 341001, "Panvel": 410206, "Patna": 800001, "Port Blair": 744101, "Rajkot": 360001,
		"Satara": 415001, "Shimla": 171001, "Surat": 395003, "Thane": 400601, "Udaipur": 313001}
	go func(map[string]int) {
		w, err := weat.NewCurrent(apiKey)
		if err != nil {
			log.Fatalln(err)
		}
		for {
			for key, val := range city {
				data, _ := w.CurrentByName(key)
				data.Zipcode = val
				var d weat.RequiredData
				db.Where(" name = ?", key).Delete(&d)
				db.Create(&data)
			}
			time.Sleep(1 * time.Minute)
		}

	}(city)

	mux := mux.NewRouter()
	mux.HandleFunc("/weather/", getCityWeatherByName).Methods("GET").Queries("city", "{city}")
	mux.HandleFunc("/weather/", getCityWEatherByLocation).Methods("GET").Queries("lat", "{lat}", "lon", "{lon}")
	mux.HandleFunc("/weather/", getCityWeaherByZip).Methods("GET").Queries("zip", "{zip}")

	fmt.Printf("Running local server on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func getCityWeaherByZip(w http.ResponseWriter, r *http.Request) {
	zip := r.URL.Query()["zip"]
	var reqObj weat.RequiredData
	db.Where(" zipcode = ?", zip).Find(&reqObj)
	json.NewEncoder(w).Encode(reqObj)
}

func getCityWEatherByLocation(w http.ResponseWriter, r *http.Request) {
	lon := r.URL.Query()["lon"]
	lat := r.URL.Query()["lat"]
	var reqObj weat.RequiredData
	db.Where(" latitude = ? AND longitude =?", lat, lon).Find(&reqObj)
	json.NewEncoder(w).Encode(reqObj)
}

func getCityWeatherByName(w http.ResponseWriter, r *http.Request) {
	params := r.FormValue("city")
	var reqObj weat.RequiredData
	db.Where(" name = ?", params).Find(&reqObj)
	json.NewEncoder(w).Encode(reqObj)
}
