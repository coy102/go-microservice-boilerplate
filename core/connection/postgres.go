package connection

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/v2/util/log"

	// Postgres dialect
	_ "github.com/lib/pq"
)

type pgConnection struct {
	*connection
}

// NewPostgresConnection ...
func NewPostgresConnection() Connection {
	return &pgConnection{
		&connection{
			db: setPostgresDbConnection(),
		},
	}
}

func (connection *pgConnection) ExecuteDBInsertReturnID(qapi string, args ...interface{}) (CrudResult, error) {
	var id int64
	errConnection := connection.db.QueryRow(qapi, args...).Scan(&id)
	crud := CrudResult{}

	if errConnection != nil {
		log.Info("Failed execute Query", errConnection)
		return crud, errConnection
	}

	crud.InsertedID = id
	crud.RowsAffected = 1
	return crud, nil
}

func setPostgresDbConnection() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASS")
	timeout := os.Getenv("DB_TIMEOUT")
	idleConn, _ := strconv.Atoi(os.Getenv("DB_IDLE_CONN"))
	openConn, _ := strconv.Atoi(os.Getenv("DB_OPEN_CONN"))
	connLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_LIFETIME"))

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s connect_timeout=%s sslmode=disable",
		host, port, user, dbname, pass, timeout,
	)

	connDB, connError := sqlx.Open("postgres", connStr)
	if connError != nil {
		log.Errorf("Error establishing connection to %s database: %v", dbname, connError)
		panic(connError)
	}

	pingError := connDB.Ping()
	if pingError != nil {
		log.Errorf("Error connecting to %s database: %s", dbname, pingError)
		panic(pingError)
	}

	connDB.SetMaxIdleConns(idleConn)
	connDB.SetMaxOpenConns(openConn)
	connDB.SetConnMaxLifetime(time.Duration(connLifetime) * time.Second)

	return connDB
}
