package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/nilerajput91/handlers"
	"github.com/nilerajput91/models"
)

//App has router and db instance
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize the predefined configuration
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	db, err := gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}
	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

//Set all requireds router
func (a *App) setRouters() {
	//routing for the handling the project
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employee/{title}/disable", a.DisableEmployee)
	a.Put("/employee/{title}/enable", a.EnableEmployee)
}

// Get method Wrap the router
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")

}

// Post method Wrap the router
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put method Wrap the router
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete method Wrap the router
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllEmployees(a.DB, w, r)
}

func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.CreateEmployee(a.DB, w, r)
}

func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.GetEmployee(a.DB, w, r)
}

func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateEmployee(a.DB, w, r)
}

func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteEmployee(a.DB, w, r)
}

func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handlers.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
