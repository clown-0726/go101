package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type XInfoDataSaudiModel struct {
	Symbol string `bson:"SYMBOL"`
}

type ReportSaudi2 struct {
	XInfoData XInfoDataSaudiModel `bson:"x_info_data"`
}

func Main4SearchDataSaudi() {
	clientOptions := options.Client().ApplyURI("")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("成功断开与MongoDB的连接！")
	}()

	fmt.Println("成功连接到MongoDB！")

	// 获取集合
	collection := client.Database("filing_reports").Collection("filing_data")

	filter := bson.D{
		{"x_reported_at_utc_date", bson.D{
			{"$gte", "2023-07-01"},
			{"$lte", "2023-12-31"},
		}},
		{"x_spider_name", "stock_saudi"},
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start the search...with concrete struct...")
	var report ReportSaudi2
	for curInd := 1; cursor.Next(context.TODO()); curInd++ {
		err := cursor.Decode(&report)
		if err != nil {
			panic(err)
		}
		fmt.Println(curInd, report)
	}

	fmt.Println("Start the search...with bson.M...")
	var result bson.M
	for curInd := 1; cursor.Next(context.TODO()); curInd++ {
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Println(curInd, result)
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}
}
