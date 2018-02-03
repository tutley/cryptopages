package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bradialabs/shortid"
	"github.com/goadesign/goa"
	"github.com/tutley/cryptopages/app"
	"golang.org/x/crypto/bcrypt"
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

// User structure for working with users
type User struct {
	ID             string   `json:"id" bson:"_id,omitempty"`
	Email          Email    `json:"email" bson:"email"`
	Password       string   `json:"-" bson:"password"`
	Name           string   `json:"name" bson:"name"`
	Username       string   `json:"username" bson:"username"`
	Available      bool     `json:"available" bson:"available"`
	Location       Location `json:"location" bson:"location"`
	Coins          Coins    `json:"coins" bson:"coins"`
	OtherCoin      string   `json:"otherCoin" bson:"otherCoin"`
	JobCategory    string   `json:"jobCategory" bson:"jobCategory"`
	Skills         []string `json:"skills" bson:"skills"`
	JobDescription string   `json:"jobDescription" bson:"jobDescription"`
}

// Email the object to hold emails
type Email struct {
	MakePublic bool   `json:"makePublic" bson:"makePublic"`
	Value      string `json:"value" bson:"value"`
}

// Location the object to hold locations
type Location struct {
	MakePublic bool   `json:"makePublic" bson:"makePublic"`
	Value      string `json:"value" bson:"value"`
}

// Coins the object to hold coins
type Coins struct {
	BTC   bool `json:"btc" bson:"btc"`
	LTC   bool `json:"ltc" bson:"ltc"`
	BCC   bool `json:"bcc" bson:"bcc"`
	ETH   bool `json:"eth" bson:"eth"`
	NEO   bool `json:"neo" bson:"neo"`
	XRP   bool `json:"xrp" bson:"xrp"`
	XLM   bool `json:"xlm" bson:"xml"`
	Other bool `json:"other" bson:"other"`
}

// NewUser creates a new User and saves it in the database
func NewUser(u User, db *mgo.Database) (*User, error) {
	//verify the username isn't already being used
	user, _ := FindUserByUsername(u.Username, db)
	if user != nil {
		return nil, errors.New("user: The user already exists")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("user: %+v", err)
	}
	sid := shortid.New()
	sid.SetSeed(200)

	newUser := User{
		ID:             sid.Generate(),
		Email:          u.Email,
		Password:       string(passHash),
		Name:           u.Name,
		Username:       u.Username,
		Available:      u.Available,
		Location:       u.Location,
		Coins:          u.Coins,
		OtherCoin:      u.OtherCoin,
		JobCategory:    u.JobCategory,
		Skills:         u.Skills,
		JobDescription: u.JobDescription,
	}

	err = newUser.Save(db)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

// FindUserByUsername searches for an existing user with the passed ID
func FindUserByUsername(username string, db *mgo.Database) (*User, error) {
	user := User{}
	err := db.C("users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByEmail searches for an existing user with the passed Email
func FindUserByEmail(email string, db *mgo.Database) (*User, error) {
	user := User{}
	err := db.C("users").Find(bson.M{"email": bson.M{"value": email}}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SearchUsers finds users by the specified criteria and returns an array of results
func SearchUsers(db *mgo.Database) (*[]User, error) {
	users := []User{}
	err := db.C("users").Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Save Upserts the user into the database
func (user *User) Save(db *mgo.Database) error {
	_, err := db.C("users").UpsertId(user.ID, user)
	return err
}

// Delete deletes the user from the database
func (user *User) Delete(db *mgo.Database) error {
	err := db.C("users").Remove(bson.M{"_id": user.ID})
	return err
}

// CheckPassword will check a passed password string with the stored hash
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// MapToCryptopagesUser takes a User object and converts it into an app.CryptopagesUser object
func (user *User) MapToCryptopagesUser() (app.CryptopagesUser, error) {
	// Something tells me I need to add error checking in here or I'm going to get null pointer problems
	u := app.CryptopagesUser{
		Username:       user.Username,
		Name:           user.Name,
		Available:      &user.Available,
		JobCategory:    user.JobCategory,
		JobDescription: &user.JobDescription,
		Skills:         user.Skills,
	}
	u.Location.Value = &user.Location.Value
	u.Location.MakePublic = &user.Location.MakePublic
	u.Email.Value = &user.Email.Value
	u.Email.MakePublic = &user.Email.MakePublic
	u.Coins.Btc = &user.Coins.BTC
	u.Coins.Ltc = &user.Coins.LTC
	u.Coins.Bcc = &user.Coins.BCC
	u.Coins.Eth = &user.Coins.ETH
	u.Coins.Neo = &user.Coins.NEO
	u.Coins.Xrp = &user.Coins.XRP
	u.Coins.Xlm = &user.Coins.XLM
	u.Coins.Other = &user.Coins.Other
	u.OtherCoin = &user.OtherCoin

	return u, nil
}
