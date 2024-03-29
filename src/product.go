package main

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Price  float64       `bson:"price" json:"price"`
	Status bool          `bson:"status" json:"status"`
}
