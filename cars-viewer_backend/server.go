package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// --- 1. Define the Data Structures ---

type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   int    `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

type CarModel struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int            `json:"manufacturerId"`
	CategoryID     int            `json:"categoryId"`
	Year           int            `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image          string         `json:"image"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Manufacturer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Founded int    `json:"foundingYear"`
}

type Data struct {
	CarModels     []CarModel     `json:"carModels"`
	Categories    []Category     `json:"categories"`
	Manufacturers []Manufacturer `json:"manufacturers"`
}

var data Data

func main() {
	// Load the JSON file
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("Error reading data.json:", err)
	}
	json.Unmarshal(file, &data)

	mux := http.NewServeMux()

	// Static files for images
	mux.Handle("/api/images/", http.StripPrefix("/api/images/", http.FileServer(http.Dir("img"))))

	// API Root
	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"models":        "/api/models",
			"categories":    "/api/categories",
			"manufacturers": "/api/manufacturers",
		})
	})

	// Routes
	mux.HandleFunc("GET /api/models", getModels)
	mux.HandleFunc("GET /api/models/{id}", getModelByID)
	mux.HandleFunc("GET /api/categories", getCategories)
	mux.HandleFunc("GET /api/categories/{id}", getCategoryByID)
	mux.HandleFunc("GET /api/manufacturers", getManufacturers)
	mux.HandleFunc("GET /api/manufacturers/{id}", getManufacturerByID)

	fmt.Println("Server is running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// --- 2. Handlers ---

func getModels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.CarModels)
}

func getModelByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	for _, item := range data.CarModels {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, `{"message": "Car model not found"}`, 404)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Categories)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	for _, item := range data.Categories {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, `{"message": "Category not found"}`, 404)
}

func getManufacturers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Manufacturers)
}

func getManufacturerByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	for _, item := range data.Manufacturers {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, `{"message": "Manufacturer not found"}`, 404)
}
