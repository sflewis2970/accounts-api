package models

import (
	"github.com/sflewis2970/accounts-api/common"
	"github.com/sflewis2970/accounts-api/config"
	"github.com/sflewis2970/accounts-api/messaging"
	"log"
	"time"
)

type AcctModel struct {
	cfgData         *config.CfgData
	postgresDBModel *PostgresDBModel
}

var acctModel *AcctModel

func (am *AcctModel) AddUser(userRequest messaging.UserRequest) error {
	// Update timestamp
	userRequest.Timestamp = common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	_, insertErr := am.postgresDBModel.Insert(userRequest)
	if insertErr != nil {
		errMsg := "Error inserting record...: "
		log.Print(errMsg, insertErr)
	}

	return insertErr
}

func (am *AcctModel) GetUser(userID string) (messaging.UserResponse, error) {
	// AnswerResponse
	var userResponse messaging.UserResponse
	var getErr error

	// Send request to get question from Redis cache
	userResponse, getErr = am.postgresDBModel.Get(userID)
	if getErr != nil {
		errMsg := "Get record error...: "
		log.Print(errMsg, getErr)
		userResponse.Message = errMsg
		return userResponse, getErr
	} else {
		// Build AnswerResponse message
	}

	return userResponse, nil
}

func (am *AcctModel) UpdateUser(userRequest messaging.UserRequest) (messaging.UserResponse, error) {
	// AnswerResponse
	var userResponse messaging.UserResponse
	var getErr error

	// Update timestamp
	userRequest.Timestamp = common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	// Send request to get question from Redis cache
	_, getErr = am.postgresDBModel.Update(userRequest)
	if getErr != nil {
		errMsg := "Get record error...: "
		log.Print(errMsg, getErr)
		userResponse.Message = errMsg
		return userResponse, getErr
	} else {
		// Build AnswerResponse message
	}

	return userResponse, nil
}

func (am *AcctModel) DeleteUser(userID string) error {
	// Send request to delete question from Redis cache
	_, deleteErr := am.postgresDBModel.Delete(userID)
	if deleteErr != nil {
		errMsg := "Delete record error...: "
		log.Print(errMsg, deleteErr)
	}

	return deleteErr
}

func (am *AcctModel) CfgData() *config.CfgData {
	return am.cfgData
}

func NewAcctModel() *AcctModel {
	log.Print("Creating model object...")
	acctModel = new(AcctModel)

	// Get config data
	acctModel.cfgData = config.NewConfig().LoadCfgData()

	// New model (cacheModel)
	acctModel.postgresDBModel = NewPostgresDBModel()

	return acctModel
}
