package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/couchbase/gocb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Car struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Year         string `json:"year,omitempty"`
}

type Part struct {
	Engine string `json:"engine"`
	Tires  string `json:"tires"`
	Lights string `json:"lights"`
}

type N1qlCar struct {
	Car Car `json:"car"`
}

var bucket *gocb.Bucket

func GetCarEndpoint(w http.ResponseWriter, req *http.Request) {
	var n1qlParams []interface{}
	query := gocb.NewN1qlQuery("SELECT * FROM cars AS car WHERE META(car).id = $1")
	params := mux.Vars(req)
	n1qlParams = append(n1qlParams, params["id"])
	rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
	var row N1qlCar
	rows.One(&row)
	json.NewEncoder(w).Encode(row.Car)

}

func GetCarsEndpoint(w http.ResponseWriter, req *http.Request) {
	var car []Car
	query := gocb.NewN1qlQuery("SELECT * FROM cars AS car")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		fmt.Printf("failed to query couchbase: %s\n", err)
		http.Error(w, err.Error(), 400)
		return
	}
	var row N1qlCar
	for rows.Next(&row) {
		car = append(car, row.Car)
	}
	json.NewEncoder(w).Encode(car)

}

func CreateCarEndpoint(w http.ResponseWriter, req *http.Request) {

	var car Car
	var n1qlParams []interface{}
	_ = json.NewDecoder(req.Body).Decode(&car)
	query := gocb.NewN1qlQuery("INSERT INTO `cars` (KEY, VALUE) VALUES ($1, {'name': $2, 'manufacturer': $3, 'year': $4})")
	n1qlParams = append(n1qlParams, uuid.New().String())
	n1qlParams = append(n1qlParams, car.Name)
	n1qlParams = append(n1qlParams, car.Manufacturer)
	n1qlParams = append(n1qlParams, car.Year)
	_, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(car)
}

func UpdateCarEndpoint(w http.ResponseWriter, req *http.Request) {
	return

}

func DeleteCarEndpoint(w http.ResponseWriter, req *http.Request) {
	return

}

func GetCarPartsEndpoint(w http.ResponseWriter, req *http.Request) {
	parts := []Part{
		{
			Engine: "2021 Camry Engine",
			Tires:  "2021 Camry Wheelset",
			Lights: "2020 Light Package",
		},
	}
	json.NewEncoder(w).Encode(parts)

}

func main() {
	host := "localhost"
	if hostEnv := os.Getenv("DB_HOST"); hostEnv != "" {
		host = hostEnv
	}

	router := mux.NewRouter()
	cluster, _ := gocb.Connect(fmt.Sprintf("%s:8091", host))
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "password",
	})
	//cluster, _ := gocb.Connect("172.17.0.5")
	//bucket, _ := cluster.OpenBucket("cars", "password")
	bucket, _ = cluster.OpenBucket("cars", "")
	router.HandleFunc("/cars", GetCarsEndpoint).Methods("GET")
	router.HandleFunc("/car/{id}", GetCarEndpoint).Methods("GET")
	router.HandleFunc("/car/camry/2021", GetCarPartsEndpoint).Methods("GET")
	router.HandleFunc("/car", CreateCarEndpoint).Methods("PUT")
	router.HandleFunc("/car/{id}", UpdateCarEndpoint).Methods("POST")
	router.HandleFunc("/car/{id}", DeleteCarEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
