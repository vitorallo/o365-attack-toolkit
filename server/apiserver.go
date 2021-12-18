package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	//log "github.com/sirupsen/logrus"

	"github.com/vitorallo/o365-attack-toolkit/api"
	"github.com/vitorallo/o365-attack-toolkit/database"
	"github.com/vitorallo/o365-attack-toolkit/logging"
	"github.com/vitorallo/o365-attack-toolkit/model"

	"github.com/gorilla/mux"
)

func StartAPIServer(config model.Config, l *logrus.Logger) {

	// Start the update token function like in the standard server
	go api.RecursiveTokenUpdate(logging.GetLogger())

	l.Info(fmt.Sprintf("Starting API Server on %s:%d", config.Server.Host, config.Server.ApiPort))

	route := mux.NewRouter()
	route.HandleFunc("/users", GetUsersAPI).Methods("GET")
	route.HandleFunc("/config", GetConfig).Methods("GET")
	route.HandleFunc("/connect", Connect).Methods("GET")

	//route.HandleFunc(model.IntAbout, GetAbout).Methods("GET")

	// Routes for Users
	//route.HandleFunc(model.IntGetAll, GetUsers).Methods("GET")

	// Route for files
	//route.HandleFunc(model.IntUserFiles, GetUserFiles).Methods("GET")
	//route.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir("downloads/"))))

	// Route for Live Interaction
	//oute.HandleFunc(model.IntLiveMain, GetLiveMain).Methods("GET")
	//route.HandleFunc(model.IntLiveSearchMail, GetLiveEmails).Methods("GET")
	//route.HandleFunc(model.IntLiveSendMail, SendEmail).Methods("POST")
	//route.HandleFunc(model.IntLiveOpenMail, GetEmail).Methods("GET")
	//route.HandleFunc(model.IntLiveSearchFiles, GetLiveFiles).Methods("GET")
	//route.HandleFunc(model.IntLiveDownloadFile, DownloadFileHandler).Methods("GET")
	//route.HandleFunc(model.IntLiveReplaceFile, ReplaceFile).Methods("POST")

	//Route for emails
	//	route.HandleFunc(model.IntUserEmails, GetUserEmails).Methods("GET")
	//	route.HandleFunc(model.IntUserEmails, SearchUserEmails).Methods("POST") //  For searching
	//	route.HandleFunc(model.IntAllEmails, GetAllEmails).Methods("GET")
	//	route.HandleFunc(model.IntAllEmails, SearchEmails).Methods("POST") // For Searching
	// Removed this as we are going to use the Live thing.
	//route.HandleFunc(model.IntUserEmail, GetEmail).Methods("GET")

	// The route for the file downloads.

	//route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.ApiPort),
		Handler: route,
	}

	server.ListenAndServe()

}

func GetUsersAPI(w http.ResponseWriter, r *http.Request) {
	logging.Log.Debug("Fetching users from the DB")
	users := database.GetUsers()

	//alternative manual way to get json out of a slice
	//data, err := json.Marshal(&users)
	//if err == nil {
	//logging.Log.Trace(string(data))
	//w.Write(data)
	//} else {
	//logging.Log.Error("Something horrible happened while fetching users: ", err)
	///}

	// a bit more safe way
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	logging.Log.Debug("Serving the external server config")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GlbConfig)
}

func Connect(w http.ResponseWriter, r *http.Request) {
	type obj struct {
		Status string
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	logging.Log.Debug("Acknowledging connection of new client")
	json.NewEncoder(w).Encode(obj{"OK"})
}
