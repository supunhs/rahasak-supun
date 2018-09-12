package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

type Weather struct{
	SerialNo string `json:"serialNo,omitempty"`
	TimeStamp string `json:"timeStamp,omitempty"`
	Temp string `json:"temparature,omitempty"`
	Humid string `json:"humidity,omitempty"`
	PM2 string `json:"pm2,omitempty"`
	Hchco string `json:"hchco,omitempty"`
	Ozone string `json:"ozone,omitempty"`
	Co2 string `json:"co2,omitempty"`
	Tvoc string `json:"tvoc,omitempty"`

}

var weatherDeatils []Weather

func GetDataEndpoint(w http.ResponseWriter,req *http.Request){
	params := mux.Vars(req)
	for _, item := range weatherDeatils{
		if item.SerialNo == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Weather{})
}

func GetAllDataEndpoint(w http.ResponseWriter,req *http.Request){
	json.NewEncoder(w).Encode(weatherDeatils)

	
}


func main(){
	router :=mux.NewRouter()
	weatherDeatils = append(weatherDeatils, Weather{SerialNo:"0001", TimeStamp:"123456", Temp:"29.8", Humid:"3", PM2:"232", Hchco:"322", Ozone:"323233", Co2:"31", Tvoc:"3232"})
	weatherDeatils = append(weatherDeatils, Weather{SerialNo:"0002", TimeStamp:"456123", Temp:"30.5", Humid:"4", PM2:"456", Hchco:"385", Ozone:"25635", Co2:"36", Tvoc:"2563"})
	router.HandleFunc("/weather/{id}", GetDataEndpoint).Methods("GET")
	router.HandleFunc("/weather", GetAllDataEndpoint).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":12345",router))

}





