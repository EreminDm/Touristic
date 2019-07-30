package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	Db *mgo.Database
}

func (productModel ProductModel) FindAll() ([]Product, error) {
	var products []Product
	err := productModel.Db.C("product").Find(bson.M{}).All(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
