package models

import (
	"database/sql"
	"fmt"
	"github.com/sflewis2970/accounts-api/config"
	"github.com/sflewis2970/accounts-api/messaging"
	"log"

	_ "github.com/lib/pq"
)

const (
	POSTGRESQL_DB_NAME_MSG string = "POSTGRESQL: "
)

const (
	POSTGRESQL_GET_CONFIG_ERROR      string = "Getting config error...: "
	POSTGRESQL_GET_CONFIG_DATA_ERROR string = "Getting config data error...: "
	POSTGRESQL_OPEN_ERROR            string = "Error opening database..."
	POSTGRESQL_INSERT_ERROR          string = "Error inserting record..."
	POSTGRESQL_GET_ERROR             string = "Error getting record..."
	POSTGRESQL_UPDATE_ERROR          string = "Error updating record..."
	POSTGRESQL_DELETE_ERROR          string = "Error deleting record..."
	POSTGRESQL_RESULTS_ERROR         string = "Error getting results...: "
	POSTGRESQL_ROWS_AFFECTED_ERROR   string = "Error getting rows affected...: "
	POSTGRESQL_PING_ERROR            string = "Error pinging database server..."
)

type PostgresDBModel struct {
	cfgData *config.CfgData
}

var postgresDBModel *PostgresDBModel
var dbDriverName string = "postgres"

// Open database
func (pdbm *PostgresDBModel) Open() (*sql.DB, error) {
	log.Println("Opening PostgreSQL database")

	// Open database connection
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pdbm.cfgData.PostGreSQL.Host, pdbm.cfgData.PostGreSQL.Port, pdbm.cfgData.PostGreSQL.User, "devAdmin", "accounts")
	db, openErr := sql.Open(dbDriverName, dataSourceName)

	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return nil, openErr
	}

	return db, nil
}

// Ping database server by verifying the database connection is active
func (pdbm *PostgresDBModel) Ping() error {
	db, openErr := pdbm.Open()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return openErr
	}
	defer db.Close()

	pingErr := db.Ping()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG + POSTGRESQL_PING_ERROR)
		return pingErr
	}

	return nil
}

// Insert a single record into table
func (pdbm *PostgresDBModel) Insert(userRequest messaging.UserRequest) (int64, error) {
	db, openErr := pdbm.Open()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return messaging.RESULTS_DEFAULT, openErr
	}
	defer db.Close()

	log.Print("Adding a new record to the database")
	queryStr := "insert into accounts VALUES ($1, $2, $3);"
	sqlDB, execErr := db.Exec(queryStr, userRequest.UserID, userRequest.Username, userRequest.Timestamp)
	if execErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG + POSTGRESQL_INSERT_ERROR)
		return messaging.RESULTS_DEFAULT, execErr
	}

	rowsAffected, rowsAffectedErr := sqlDB.RowsAffected()
	if rowsAffectedErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_ROWS_AFFECTED_ERROR, rowsAffectedErr.Error())
		return messaging.RESULTS_DEFAULT, rowsAffectedErr
	}

	return rowsAffected, nil
}

// Get a single record from table
func (pdbm *PostgresDBModel) Get(userID string) (messaging.UserResponse, error) {
	db, openErr := pdbm.Open()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return messaging.UserResponse{}, openErr
	}
	defer db.Close()

	var userResponse messaging.UserResponse

	log.Print("Getting a single record from the database")
	queryStr := "SELECT username, timestamp FROM accounts WHERE user_id = $1;"
	scanErr := db.QueryRow(queryStr, userID).Scan(&userResponse.Username, &userResponse.Timestamp)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_GET_ERROR, scanErr.Error())
		return messaging.UserResponse{}, scanErr
	}

	return userResponse, nil
}

// Update a single record in table
func (pdbm *PostgresDBModel) Update(userRequest messaging.UserRequest) (int64, error) {
	db, openErr := pdbm.Open()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return messaging.RESULTS_DEFAULT, openErr
	}
	defer db.Close()

	log.Println("Updating a single record in the database")
	queryStr := "UPDATE accounts SET username = $2, timestamp = $3 WHERE userid = $1"
	sqlDB, execErr := db.Exec(queryStr, userRequest.UserID, userRequest.Username, userRequest.Timestamp)
	if execErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_UPDATE_ERROR, execErr.Error())
		return messaging.RESULTS_DEFAULT, execErr
	}

	rowsAffected, rowsAffectedErr := sqlDB.RowsAffected()
	if rowsAffectedErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_ROWS_AFFECTED_ERROR, rowsAffectedErr.Error())
		return messaging.RESULTS_DEFAULT, nil
	}

	return rowsAffected, nil
}

// Delete a single record from table
func (pdbm *PostgresDBModel) Delete(userID string) (int64, error) {
	db, openErr := pdbm.Open()
	if openErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_OPEN_ERROR, openErr.Error())
		return messaging.RESULTS_DEFAULT, openErr
	}
	defer db.Close()

	log.Println("deleting a single record from the database")
	queryStr := "DELETE FROM accounts WHERE question_id = $1"
	sqlDB, execErr := db.Exec(queryStr, userID)
	if execErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_DELETE_ERROR, execErr.Error())
		return messaging.RESULTS_DEFAULT, execErr
	}

	rowsAffected, rowsAffectedErr := sqlDB.RowsAffected()
	if rowsAffectedErr != nil {
		log.Print(POSTGRESQL_DB_NAME_MSG+POSTGRESQL_ROWS_AFFECTED_ERROR, rowsAffectedErr.Error())
		return messaging.RESULTS_DEFAULT, nil
	}

	return rowsAffected, nil
}

func NewPostgresDBModel() *PostgresDBModel {
	log.Print("Creating PostgreSQL database model")

	// Initialize PostgreSQL database model
	postgresDBModel := new(PostgresDBModel)

	// load config data
	postgresDBModel.cfgData = config.NewConfig().LoadCfgData()

	return postgresDBModel
}
