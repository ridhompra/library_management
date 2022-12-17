package controller

import (
	"fmt"
	"net/http"
)

func Hompage(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintln("Tes Homepage")
}
