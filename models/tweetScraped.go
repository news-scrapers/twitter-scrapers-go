package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type NewScraped struct {
	Page      int       `json:"page"`
	FullPage  bool      `json:"full_page"`
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
	ID        string    `json:"id"`
}

func (newScraped *NewScraped) Create() {

	db := GetDB()
	collection := db.Collection("TweetsScraped")

	// options := options.FindOneAndReplaceOptions{}
	// upsert := true
	// options.Upsert = &upsert
	// err := collection.FindOneAndReplace(context.Background(), bson.M{"url": newScraped.Url}, newScraped, &options)

	result := &NewScraped{}
	err := collection.FindOne(context.Background(), bson.M{"url": newScraped.Url}).Decode(result)

	if err != nil {
		err2, _ := collection.InsertOne(context.Background(), newScraped)
		if err2 != nil {
			fmt.Println("saved new with " + newScraped.Url)
		} else {
			fmt.Println("error saving new with url " + newScraped.Url)
			fmt.Println(err)
		}
	} else {
		fmt.Println("record already exists with url " + newScraped.Url)
	}

	//_, err := collection.InsertOne(context.Background(), newScraped)

}
