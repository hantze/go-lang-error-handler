package gosqlx

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"gosqlx/internal/gosqlx/router"
)

type HttpServer struct {}

func (hs *HttpServer) Serve(db *sqlx.DB){
	r := router.NewV1Router(db)
	log.Printf("About to listen on 3333. Go to http://127.0.0.1:3333")
	http.ListenAndServe(":3333", r)
}

func NewHttpServer() *HttpServer{
	return &HttpServer{}
}