package controllers

import (
	"net/http"

	"github.com/edu2801/first-golang-crud/api/models"
	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []models.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (tc *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tc.tweets)
}

func (tc *tweetController) Create(ctx *gin.Context) {
	tweet := models.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	if tweet.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description field is required"})
		return
	}

	tc.tweets = append(tc.tweets, *tweet)

	ctx.JSON(http.StatusCreated, tweet)
}

func (tc *tweetController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, v := range tc.tweets {
		if v.ID == id {
			tweet := models.NewTweet()

			if err := ctx.BindJSON(&tweet); err != nil {
				return
			}

			tc.tweets[idx] = *tweet

			ctx.JSON(http.StatusOK, tweet)
			return
		}
	}
}

func (tc *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, v := range tc.tweets {
		if v.ID == id {
			tc.tweets = append(tc.tweets[0:idx], tc.tweets[idx+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}
