package main

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"gopkg.in/mgo.v2"
)

// DBContextKey is exported but I'm not sure if it needs to be
type DBContextKey string

// DBKey is the key used to reference the database pointer in the context
var DBKey = DBContextKey("mogodb")

// MongoMiddleware adds a pointer to the database to the context
func MongoMiddleware(db *mgo.Database) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			newctx := context.WithValue(ctx, DBKey, db)
			return h(newctx, rw, req)
		}
	}
}
