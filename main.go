package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/mvanbrummen/go-rest/middleware"
	"github.com/mvanbrummen/go-rest/repository"

	"github.com/gorilla/mux"
	"github.com/mvanbrummen/go-rest/handler"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/negroni"

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

	r.Use(middleware.ContentTypeMW)

	titlesRepository := repository.NewTitlesRepository(db)

	handler := handler.NewTitlesHandler(titlesRepository)

	r.HandleFunc("/titles/{id}", handler.GetTitle)
	r.Path("/titles").Queries("q", "{q}").HandlerFunc(handler.SearchTitle)
	r.Path("/titles").Queries("q", "{q}", "limit", "{limit}").HandlerFunc(handler.SearchTitle)

	n := negroni.Classic()
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
