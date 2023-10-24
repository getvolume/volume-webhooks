package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/getvolume/webhook-validator-go"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	router := gin.Default()

	router.PUT("/webhook", func(ctx *gin.Context) {

		// Get signarure from authorisationHeader
		authorisationHeader := ctx.GetHeader("Authorization")
		signature := signatureFromAuthRequest(authorisationHeader)

		// init validator and check body with signature
		pemCertUrl := os.Getenv("PEM_URL")
		bodyAsByteArray, _ := io.ReadAll(ctx.Request.Body)
		body := string(bodyAsByteArray)

		valid := validator.New(pemCertUrl)
		validationResult := valid.Run(string(body), signature)

		ctx.IndentedJSON(http.StatusCreated, validationResult)
	})

	router.Run(":8080")
}

func signatureFromAuthRequest(header string) string {
	authorisationHeaderValue := strings.Split(header, " ")
	var signature string
	if len(authorisationHeaderValue) > 0 {
		signature = authorisationHeaderValue[len(authorisationHeaderValue)-1]
	}
	return signature
}
