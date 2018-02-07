package main

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/tutley/cryptopages/app"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DBContextKey is exported but I'm not sure if it needs to be
type DBContextKey string

// DBKey is the key used to reference the database pointer in the context
var DBKey = DBContextKey("mogodb")

// UsernameContextKey is exported but I'm not sure if it needs to be
type UsernameContextKey string

// UsernameKey is the key used to reference the database pointer in the context
var UsernameKey = UsernameContextKey("thisiswheretheusernamegoes")

// MongoMiddleware adds a pointer to the database to the context
func MongoMiddleware(db *mgo.Database) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			newctx := context.WithValue(ctx, DBKey, db)
			return h(newctx, rw, req)
		}
	}
}

// GetDB grabs the db pointer off the context
func GetDB(ctx context.Context) *mgo.Database {
	db := ctx.Value(DBKey).(*mgo.Database)
	return db
}

// FindUserByUsername searches for an existing user with the passed ID
func FindUserByUsername(username string, db *mgo.Database) (*app.CryptopagesUser, error) {
	user := app.CryptopagesUser{}
	err := db.C("users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
