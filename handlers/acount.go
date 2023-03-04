package handlers

import (
	"encoding/json"
	"github.com/sflewis2970/accounts-api/messaging"
	"github.com/sflewis2970/accounts-api/models"
	"log"
	"net/http"
)

type AccountHandler struct {
	acctModel *models.AcctModel
}

var acctHandler *AccountHandler

func (ah *AccountHandler) RegisterUser(rw http.ResponseWriter, r *http.Request) {
	// Display a log message
	log.Print("data received from client...")

	var userRequest messaging.UserRequest
	var userResponse messaging.UserResponse

	// Read JSON from stream
	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	if decodeErr != nil {
		log.Print("Error decoding json...: ", decodeErr)

		// Update AnswerResponse
		userResponse.Message = decodeErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Process API Get Request
	acctErr := ah.acctModel.AddUser(userRequest)
	if acctErr != nil {
		log.Print("Error encoding json...:", acctErr)

		// Update QuestionResponse struct
		userResponse.Message = acctErr.Error()

		// Update HTTP header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Update HTTP Header
	rw.WriteHeader(http.StatusCreated)

	// Write JSON to stream
	encodeResponse(rw, userResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func (ah *AccountHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	var userRequest messaging.UserRequest
	var userResponse messaging.UserResponse

	// Read JSON from stream
	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	if decodeErr != nil {
		log.Print("Error decoding json...: ", decodeErr)

		// Update AnswerResponse
		userResponse.Message = decodeErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send a request to the model for the answer
	var getErr error
	userResponse, getErr = ah.acctModel.GetUser(userRequest.UserID)

	if getErr != nil {
		log.Print("Error getting api answer...: ", getErr)

		// Update AnswerResponse
		userResponse.Message = getErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send OK status
	rw.WriteHeader(http.StatusOK)

	// Encode response
	encodeResponse(rw, userResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func (ah *AccountHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userRequest messaging.UserRequest
	var userResponse messaging.UserResponse

	// Read JSON from stream
	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	if decodeErr != nil {
		log.Print("Error decoding json...: ", decodeErr)

		// Update AnswerResponse
		userResponse.Message = decodeErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send a request to the model for the answer
	var getErr error
	userResponse, getErr = ah.acctModel.UpdateUser(userRequest)

	if getErr != nil {
		log.Print("Error getting api answer...: ", getErr)

		// Update AnswerResponse
		userResponse.Message = getErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send OK status
	rw.WriteHeader(http.StatusOK)

	// Encode response
	encodeResponse(rw, userResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func (ah *AccountHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var userRequest messaging.UserRequest
	var userResponse messaging.UserResponse

	// Read JSON from stream
	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	if decodeErr != nil {
		log.Print("Error decoding json...: ", decodeErr)

		// Update AnswerResponse
		userResponse.Message = decodeErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send a request to the model for the answer
	var getErr error
	getErr = ah.acctModel.DeleteUser(userRequest.UserID)

	if getErr != nil {
		log.Print("Error getting api answer...: ", getErr)

		// Update AnswerResponse
		userResponse.Message = getErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeResponse(rw, userResponse)
		return
	}

	// Send OK status
	rw.WriteHeader(http.StatusOK)

	// Encode response
	encodeResponse(rw, userResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func encodeResponse(rw http.ResponseWriter, userResponse messaging.UserResponse) {
	// Write JSON to stream
	encodeErr := json.NewEncoder(rw).Encode(userResponse)
	if encodeErr != nil {
		log.Print("Error encoding json...:", encodeErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func NewAccountHandler() *AccountHandler {
	acctHandler = new(AccountHandler)

	// Create api model
	acctHandler.acctModel = models.NewAcctModel()

	return acctHandler
}
