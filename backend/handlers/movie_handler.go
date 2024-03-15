package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitishmakam/embox/database"
	"github.com/nitishmakam/embox/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMovie(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var movie types.Movie
	err := database.MoviesCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusNoContent)
			return
		}
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, movie)
}

func GetMovies(c *gin.Context) {
	movies, err := database.MoviesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var moviejson []types.Movie
	if err := movies.All(context.TODO(), &moviejson); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, moviejson)
}

func AddMovie(c *gin.Context) {
	var newMovie types.Movie
	if err := c.BindJSON(&newMovie); err != nil {
		log.Fatal(err)
	}
	newMovie.ID = primitive.NewObjectID()
	result, err := database.MoviesCollection.InsertOne(context.TODO(), newMovie)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted row with ID: %v", result.InsertedID)
	c.IndentedJSON(http.StatusOK, newMovie)
}

func UpdateMovie(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var movie types.Movie
	var movieInDb types.Movie
	if err := c.BindJSON(&movie); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"code": "Request Body is invalid."})
	}
	err := database.MoviesCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&movieInDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusBadRequest, map[string]string{"code": "RESOURCE_NOT_FOUND", "message": "Movie with ID: " + c.Param("id") + " does not exist."})
			return
		}
	}
	movie.ID = id
	database.MoviesCollection.ReplaceOne(context.TODO(), bson.D{{Key: "_id", Value: id}}, movie)
	c.IndentedJSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var movieInDb types.Movie
	err := database.MoviesCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&movieInDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusBadRequest, map[string]string{"code": "RESOURCE_NOT_FOUND", "message": "Movie with ID: " + c.Param("id") + " does not exist."})
			return
		}
	}
	database.MoviesCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	c.Status(http.StatusNoContent)
}
