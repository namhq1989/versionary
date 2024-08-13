package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string

	UserSubscription        string
	UserSubscriptionHistory string
}{
	User: "user.users",

	UserSubscription:        "subscription.userSubscriptions",
	UserSubscriptionHistory: "subscription.userSubscriptionHistories",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
