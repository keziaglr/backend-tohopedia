package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/keziaglr/backend-tohopedia/graph"
	"github.com/keziaglr/backend-tohopedia/graph/generated"
	"github.com/keziaglr/backend-tohopedia/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB;

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000", "http://localhost:8080"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB : db,
	}}))
			
	router.Handle("/", playground.Handler("Tohopedia", "/query"))
	router.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
    if err != nil {
        panic(err)
    }
}

func initDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/tohopedia?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    }
	
	migrate();
	seeds();
}

func migrate(){
	db.Exec("DROP TABLE Users")
    db.AutoMigrate(&model.User{})	
}

func seeds(){
	seedUser();
}

func seedUser() {
	user := []model.User{
		{
			Email: "kezia@mail.com",
			Password: "kezia123",
		}, {
			Email: "gloria@mail.com",
			Password: "gloria123",
		},
	}

	db.Create(&user)
}

