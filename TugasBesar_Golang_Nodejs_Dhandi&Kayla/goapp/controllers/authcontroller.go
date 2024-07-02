package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DhandiAdam/OnnlineFood/config"
	"github.com/DhandiAdam/OnnlineFood/entities"
	"github.com/DhandiAdam/OnnlineFood/libraries"
	"github.com/DhandiAdam/OnnlineFood/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `validate:"required"`
	Password string `valida	e:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Redirect to Node.js server for the purchase page
	http.Redirect(w, r, "http://localhost:3000", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// proses login
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)

		} else {

			var user entities.User
			userModel.Where(&user, "username", UserInput.Username)

			var message error
			if user.Username == "" {
				message = errors.New("Username atau Password salah!")
			} else {
				// pengecekan password
				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Username atau Password salah!")
				}
			}

			if message != nil {

				data := map[string]interface{}{
					"error": message,
				}

				temp, _ := template.ParseFiles("views/login.html")
				temp.Execute(w, data)
			} else {
				// set session
				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["username"] = user.Username
				session.Values["nama_lengkap"] = user.NamaLengkap

				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

	}

}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi

		// mengambil inputan form
		r.ParseForm()

		user := entities.User{
			NamaLengkap: r.Form.Get("nama_lengkap"),
			Email:       r.Form.Get("email"),
			Username:    r.Form.Get("username"),
			Password:    r.Form.Get("password"),
			Cpassword:   r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)

		data := map[string]interface{}{
			"user": user,
		}

		if errorMessages != nil {
			data["validation"] = errorMessages
			data["alert"] = "Pendaftaran gagal. Silakan periksa kembali data Anda."
		} else {
			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			// insert ke database
			userModel.Create(user)

			data["alert"] = "Registrasi berhasil!"
		}

		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, data)
	}
}

func Tampilan_Utama(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		data := map[string]interface{}{
			"nama_lengkap": session.Values["nama_lengkap"],
		}

		temp, err := template.ParseFiles("views/Tampilan_Utama.html")
		if err != nil {
			http.Error(w, "Could not parse template", http.StatusInternalServerError)
			return
		}

		temp.Execute(w, data)
	}
}

func SendDataToNodeJS(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello from Golang!",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:3000/receive-data", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Fprintln(w, "Data sent to Node.js successfully!")
	} else {
		http.Error(w, "Failed to send data to Node.js", http.StatusInternalServerError)
	}
}

// func receiveData(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var data Data
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Printf("Received data: %+v\n", data)

// 	// Create response data
// 	responseData := Response{
// 		Message: "Data received successfully",
// 		Status:  "success",
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(responseData)
// }
