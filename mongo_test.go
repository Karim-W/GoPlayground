package main

// Use
//  docker run -d --name mongobench -p 27017:27017 mongo
// to create a suitable test instance and access it via
//  docker exec -it mongobench mongo

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	ID      primitive.ObjectID `bson:"_id"`
	Counter int
}

const DefaultHost = "localhost:27017"
const DefaultDB = "test"
const DefaultCollection = "bench"

var client *mongo.Client

func init() {
	var err error

	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + DefaultHost))
	if err != nil {
		log.Fatalf("setting up client: %s", err)
	}

	if err = client.Connect(context.Background()); err != nil {
		log.Fatalf("connecting with client: %s", err)
	}
	go func() {
		for {
			client.Ping(context.Background(), nil)
			time.Sleep(1 * time.Second)
		}
	}()
}

func BenchmarkMongoInsert(b *testing.B) {
	c := client.Database(DefaultDB).Collection("mongosimple")
	if _, err := c.DeleteMany(context.Background(), bson.M{}); err != nil {
		b.Logf("cleaning collection 'mongosimple': %s", err)
		b.FailNow()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.InsertOne(context.Background(), &Data{ID: primitive.NewObjectID(), Counter: i})
	}
}

func BenchmarkMongoBulk(b *testing.B) {

	c := client.Database(DefaultDB).Collection("mongobulk")
	if _, err := c.DeleteMany(context.Background(), bson.M{}); err != nil {
		b.Logf("cleaning collection 'mongosimple': %s", err)
		b.FailNow()
	}

	d := make([]mongo.WriteModel, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d[i] = mongo.NewInsertOneModel().SetDocument(Data{ID: primitive.NewObjectID(), Counter: i})
	}

	if _, err := c.BulkWrite(context.Background(), d); err != nil {
		b.Logf("inserting bulk: %s", err)
		b.FailNow()
	}

}

func BenchmarkMongoFind(b *testing.B) {
	c := client.Database(DefaultDB).Collection("mongobulk1")
	if _, err := c.DeleteMany(context.Background(), bson.M{}); err != nil {
		b.Logf("cleaning collection 'mongosimple': %s", err)
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		c.InsertOne(context.Background(), &Data{ID: primitive.NewObjectID(), Counter: i})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		now := time.Now()
		c.FindOne(context.Background(), bson.M{"counter": i})
		b.Logf("find %d: %s", i, time.Since(now))
	}
}
