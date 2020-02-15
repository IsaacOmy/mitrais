package handle

import (
	"log"
	"net/http"

	"github.com/IsaacOmy/mitrais/functions"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", functions.Registration).Methods("GET")
	myRouter.HandleFunc("/login", functions.Login).Methods("GET")
	myRouter.HandleFunc("/registration", functions.Registration).Methods("GET")
	myRouter.HandleFunc("/registration", functions.SubmitRegistration).Methods("POST")
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(myRouter)))

}
