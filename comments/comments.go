package comments

import (
	"context"
	"net/http"
	"time"

	"github.com/adrianedy/go-rest/database"
	"github.com/adrianedy/go-rest/response"
	"github.com/adrianedy/go-rest/router/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName string = "s"

type comment struct {
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	MovieId   primitive.ObjectID `bson:"movie_id,omitempty" json:"movie_id,omitempty"`
	Text      string             `bson:"text,omitempty" json:"text,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	movieId, _ := primitive.ObjectIDFromHex(r.FormValue("movie_id"))

	comment := comment{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		MovieId:   movieId,
		Text:      r.FormValue("text"),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	database.Collection(collectionName).InsertOne(context.TODO(), comment)

	response.JsonData(w, make(map[string]interface{}))
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(utils.GetParameter(r, "id"))
	movieId, _ := primitive.ObjectIDFromHex(r.FormValue("movie_id"))

	comment := comment{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		MovieId: movieId,
		Text:    r.FormValue("text"),
	}

	database.Collection(collectionName).UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{{Key: "$set", Value: comment}},
	)

	response.JsonData(w, make(map[string]interface{}))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(utils.GetParameter(r, "id"))

	database.Collection(collectionName).DeleteOne(
		context.TODO(),
		bson.M{"_id": id},
	)

	response.JsonData(w, make(map[string]interface{}))
}
