package swagger

import (
	"fmt"
	"net/http"

	"github.com/fomichalopoulos/companiesMicroService/helpers"
	"github.com/fomichalopoulos/companiesMicroService/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := helpers.DecodeJSON(r, &user); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Username == "" || user.Password == "" {
		helpers.ErrorResponse(w, "Please provide both username and password for the login", http.StatusBadRequest)
		return
	}

	if user.Username == "Fotis" && user.Password == "5657" {
		tokenStr, err := createToken(user)
		if err != nil {
			helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenStr)
	} else {
		helpers.ErrorResponse(w, "Invalid creds", http.StatusUnauthorized)
	}

}
