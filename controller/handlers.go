package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task/module"
)

var users []module.User

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser module.User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users = append(users, newUser)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Пользватель  %s успешно создан!\n", newUser.Name)
	} else if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	} else {
		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
	}
}
