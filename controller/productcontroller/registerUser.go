package productcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/library_Management/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusNotFound, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, book)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if err := models.DB.Create(&book).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, book)
	log.Printf("Creating success id :%d", book.Id)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		message := fmt.Sprintf("updating failed, id :%d not found", book.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}
	book.Id = id
	ResponseJson(w, http.StatusOK, book)
	log.Printf("Updating success id:%d", book.Id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	var book models.Book
	if models.DB.Delete(&book, input["id"]).RowsAffected == 0 {
		message := fmt.Sprintf("Deleting Failed, id :%d not found", book.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}
	message := map[string]string{"message": "book berhasil dihapus"}
	ResponseJson(w, http.StatusOK, message)
	log.Printf("Deleting success id:%d", book.Id)
}
