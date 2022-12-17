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
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusNotFound, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if err := models.DB.Create(&user).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, user)
	log.Printf("Creating success id :%d", user.Id)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if models.DB.Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		message := fmt.Sprintf("updating failed, id :%d not found", user.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}
	user.Id = id
	ResponseJson(w, http.StatusOK, user)
	log.Printf("Updating success id:%d", user.Id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	var user models.User
	if models.DB.Delete(&user, input["id"]).RowsAffected == 0 {
		message := fmt.Sprintf("Deleting Failed, id :%d not found", user.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}
	message := map[string]string{"message": "user berhasil dihapus"}
	ResponseJson(w, http.StatusOK, message)
	log.Printf("Deleting success id:%d", user.Id)
}
