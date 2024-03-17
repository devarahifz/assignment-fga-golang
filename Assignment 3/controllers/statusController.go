package controllers

import (
	"assignment3/models"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func UpdateStatus(status chan models.Status) {
	go func() {
		for {
			status <- models.Status{Water: rand.Intn(100) + 1, Wind: rand.Intn(100) + 1}
			<-time.After(15 * time.Second)
		}
	}()
}

func GetStatusText(s models.Status) (waterStatus, windStatus string) {
	if s.Water < 5 {
		waterStatus = "Aman"
	} else if s.Water >= 6 && s.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if s.Wind < 6 {
		windStatus = "Aman"
	} else if s.Wind >= 7 && s.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}
	return waterStatus, windStatus
}

func ShowHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, models.Status{Water: rand.Intn(100) + 1, Wind: rand.Intn(100) + 1})
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
