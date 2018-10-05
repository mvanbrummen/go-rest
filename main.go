package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mvanbrummen/go-rest/repository"

	"github.com/gorilla/mux"
	"github.com/mvanbrummen/go-rest/http"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	port := viper.GetInt("application.port")

	dbHost := viper.GetString("db.host")
	dbPort := viper.GetInt("db.port")
	dbName := viper.GetString("db.name")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbSchema := viper.GetString("db.schema")

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSchema)

	log.Println("DB connection string" + connectionString)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to DB successfully")

	r := mux.NewRouter()

	titlesRepository := repository.NewTitlesRepository(db)

	handler.NewTitlesHandler(r, titlesRepository)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
