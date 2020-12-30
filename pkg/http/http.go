package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/nivaldogmelo/web-api-tester/pkg/requests"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, "Hello World\n")
	// fmt.Fprint(w, "Hello World\n")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request root.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := requests.Save(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusCreated, "Saved with Success")
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	content, err := requests.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusCreated, content)
}

func Serve(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/", saveHandler).Methods("POST")
	r.HandleFunc("/requests", getAllHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(port, r))
}
