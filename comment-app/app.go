package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	httpDelivery "github.com/stanleynguyen/git-comment/comment-app/delivery/http"
	"github.com/stanleynguyen/git-comment/comment-app/repository/ghcli"
	"github.com/stanleynguyen/git-comment/comment-app/repository/persistence"
	"github.com/stanleynguyen/git-comment/comment-app/usecase/comment"
)

func startInDev() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := bootStrapDB()
	if err != nil {
		log.Fatal(err)
	}
	router := bootstrapApplication(db)
	log.Printf("Listening on %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func startInProd() {
	db, err := bootStrapDB()
	if err != nil {
		log.Fatal(err)
	}
	router := bootstrapApplication(db)
	log.Printf("Listening on %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func startInTest(bootstrappingDone chan<- bool) {
	db, err := bootStrapDB()
	if err != nil {
		log.Fatal(err)
	}
	router := bootstrapApplication(db)
	log.Printf("Listening on %s", os.Getenv("PORT"))
	bootstrappingDone<-true
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("PORT"), router))
}

func bootStrapDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, nil
}

func bootstrapApplication(db *sql.DB) http.Handler {
	postgresRepo := persistence.NewPostgresRepo(db)
	ghCli := ghcli.NewBasicGithubClient()
	cu := comment.NewCommentUsecase(postgresRepo, ghCli)
	router := httprouter.New()
	httpHandler := httpDelivery.Handler{Router: router}
	httpHandler.InitCommentsHandler(cu)

	return router
}
