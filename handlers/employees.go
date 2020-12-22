package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/nilerajput91/models"
)

func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []models.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	respondJSON(w, http.StatusCreated, employee)
}

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["Name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["Name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	respondJSON(w, http.StatusOK, employee)
}

func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Disable()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Enable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)

}

// getEmployeeOr404 gets a employee instance if exists or respond the 404 error otherwise

func getEmployeeOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *models.Employee {
	employee := models.Employee{}
	if err := db.First(&employee, models.Employee{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}
