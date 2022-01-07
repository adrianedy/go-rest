package movies

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adrianedy/go-rest/database"
	"github.com/adrianedy/go-rest/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "movies"

type Awards struct {
	Wins        int    `bson:"wins,omitempty" json:"wins,omitempty"`
	Nominations int    `bson:"nominations,omitempty" json:"nominations,omitempty"`
	Text        string `bson:"text,omitempty" json:"text,omitempty"`
}

type Imdb struct {
	Rating float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	Votes  int     `bson:"votes,omitempty" json:"votes,omitempty"`
	Id     int     `bson:"id,omitempty" json:"id,omitempty"`
}

type Viewer struct {
	Rating float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	Votes  int     `bson:"votes,omitempty" json:"votes,omitempty"`
	Meter  int     `bson:"meter,omitempty" json:"meter,omitempty"`
}

type Critic struct {
	Rating     float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	NumReviews int     `bson:"numReviews,omitempty" json:"numReviews,omitempty"`
	Meter      int     `bson:"meter,omitempty" json:"meter,omitempty"`
}

type Tomatoes struct {
	Viewer      Viewer             `bson:"viewer,omitempty" json:"viewer,omitempty"`
	Dvd         primitive.DateTime `bson:"dvd,omitempty" json:"dvd,omitempty"`
	Critic      Critic             `bson:"critic,omitempty" json:"critic,omitempty"`
	LastUpdated primitive.DateTime `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Consensus   string             `bson:"consensus,omitempty" json:"consensus,omitempty"`
	Rotten      int                `bson:"rotten,released" json:"rotten,omitempty"`
	Production  string             `bson:"production,omitempty" json:"production,omitempty"`
	Fresh       int                `bson:"fresh,omitempty" json:"fresh,omitempty"`
}

type Movie struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Plot               string             `bson:"plot,omitempty" json:"plot,omitempty"`
	Genres             []string           `bson:"genres,omitempty" json:"genres,omitempty"`
	Runtime            int                `bson:"runtime,omitempty" json:"runtime,omitempty"`
	Casts              []string           `bson:"casts,omitempty" json:"casts,omitempty"`
	Num_mflix_comments int                `bson:"num_mflix_comments,omitempty" json:"num_mflix_comments,omitempty"`
	Title              string             `bson:"title,omitempty" json:"title,omitempty"`
	Fullplot           string             `bson:"fullplot,omitempty" json:"fullplot,omitempty"`
	Countries          []string           `bson:"countries,omitempty" json:"countries,omitempty"`
	Released           primitive.DateTime `bson:"released,omitempty" json:"released,omitempty"`
	Directors          []string           `bson:"directors,omitempty" json:"directors,omitempty"`
	Writers            []string           `bson:"writers,omitempty" json:"writers,omitempty"`
	Rated              string             `bson:"rated,released" json:"rated,omitempty"`
	Awards             Awards             `bson:"awards,released" json:"awards,omitempty"`
	Lastupdated        string             `bson:"lastupdated,released" json:"lastupdated,omitempty"`
	Year               int                `bson:"year,omitempty" json:"year,omitempty"`
	Imdb               Imdb               `bson:"imdb,released" json:"imdb,omitempty"`
	Type               string             `bson:"type,released" json:"type,omitempty"`
	Tomatoes           Tomatoes           `bson:"tomatoes,released" json:"tomatoes,omitempty"`
}

func Read(w http.ResponseWriter, r *http.Request) {
	var filter bson.D
	queryString := r.URL.Query()
	rated := queryString.Get("rated")
	countries := queryString["countries[]"]
	languages := queryString["languages[]"]

	if rated != "" {
		filter = append(filter, primitive.E{Key: "rated", Value: rated})
	}

	if len(countries) > 0 {
		filter = append(filter, primitive.E{Key: "countries", Value: countries})
	}

	if len(languages) > 0 {
		filter = append(filter, primitive.E{Key: "languages", Value: languages})
	}

	limitQuery := queryString.Get("limit")
	if limitQuery == "" {
		limitQuery = "10"
	}
	limit, _ := strconv.ParseInt(limitQuery, 10, 64)

	opts := options.FindOptions{
		Limit: &limit,
	}

	var results []Movie
	cur, _ := database.Collection(collectionName).Find(context.TODO(), filter, &opts)

	for cur.Next(context.TODO()) {
		var movie Movie
		err := cur.Decode(&movie)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, movie)
	}

	response.JsonData(w, results)
}
