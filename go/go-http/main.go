package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/getvolume/webhook-validator-go"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	logger.Printf("Started simple http server")
	http.HandleFunc("/webhook", handlePutWebhook)
	http.ListenAndServe(":8080", nil)
}

func handlePutWebhook(response http.ResponseWriter, request *http.Request) {
	if isPutCall(request, response) {
		return
	}

	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(response, "Error reading body", http.StatusInternalServerError)
		return
	}

	defer request.Body.Close()

	// Get signarure from authorisationHeader 
	signature := signatureFromRequest(request)

	// init validator and check body with signature
	pemCertUrl := os.Getenv("PEM_URL")
	valid:= validator.New(pemCertUrl)
	validationResult := valid.Run(string(body), signature)

	response.WriteHeader(http.StatusOK)

	responseData, err := json.Marshal(validationResult)
	if err != nil {
		http.Error(response, "Failed to generate response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(responseData)
}

func signatureFromRequest(request *http.Request) string {
	authorisationHeader := request.Header.Get("Authorization")

	authorisationHeaderValue := strings.Split(authorisationHeader, " ")

	var signature string
	if len(authorisationHeaderValue) > 0 {
		signature = authorisationHeaderValue[len(authorisationHeaderValue)-1]
	}
	return signature
}

func isPutCall(r *http.Request, w http.ResponseWriter) bool {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
		return true
	}
	return false
}
