package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mebackend/profileapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)


// NOTE :  ===****===****===****===****===****===****====== replace original db key
const connectionString = "REPLACE DATABASE KEY"

const dbName = "volunteer-Management-system"
const colName = "volunteer-details"

// MOST IMPORTANT
var collection *mongo.Collection

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)
	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping the database to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	// collection instance
	fmt.Println("Collection instance is ready")
}



// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// // Insert dummy-user data
// func insertUserData(data model.VolunteerDetails) {
// 	inserted, err := collection.InsertOne(context.Background(), data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted one volunteer data in DB with ID:", inserted.InsertedID)
// }

// // create data
// func CreateData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")

// 	var data model.VolunteerDetails
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	insertUserData(data)
// 	json.NewEncoder(w).Encode(data)
// }
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++



// get user data based on email
func GetUserData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// get email from query parameter
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

 	//  fmt.Println("Searching for email:", email)

	var user model.VolunteerDetails
	filter := bson.M{"email": email}

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// fmt.Println("User found:", user) 
	json.NewEncoder(w).Encode(user)
}
