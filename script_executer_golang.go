package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee struct to represent MongoDB document
type Employee struct {
	EmpID       int    `bson:"emp_id"`
	Department  string `bson:"department"`
	Salary      int    `bson:"salary"`
	WorkingMode string `bson:"working_mode"`
}

// empData function to generate employee data
func empData() []Employee {
	var employees []Employee
	for i := 0; i < 20; i++ {
		employee := Employee{
			EmpID:       i,
			Department:  randomChoice([]string{"IT", "HR", "Sales", "Operations"}),
			Salary:      rand.Intn(90001) + 10000, // Generates random salary between 10000 and 100000
			WorkingMode: randomChoice([]string{"Hybrid", "onsite"}),
		}
		log.Printf("%+v\n", employee)
		employees = append(employees, employee)
	}
	return employees
}

// writeData function to insert data into MongoDB
func writeData(employees []Employee) error {
	connStr := "mongodb://chetan_user:chetan123@db.stg.bdb.ai:27017/pipeline?authMechanism=SCRAM-SHA-256"

	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	collection := client.Database("pipeline").Collection("emp_data")

	// Convert []Employee to []interface{}
	var interfaceSlice []interface{} = make([]interface{}, len(employees))
	for i, v := range employees {
		interfaceSlice[i] = v
	}

	// Inserting data into MongoDB
	_, err = collection.InsertMany(ctx, interfaceSlice)
	return err
}

// randomChoice function to choose a random element from a string slice
func randomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed for random number generation

	// Generate employee data
	employees := empData()
	fmt.Println("=============================Employee data read===============================")

	// Write data to MongoDB
	err := writeData(employees)
	if err != nil {
		log.Fatal("Error writing data to MongoDB:", err)
		return
	}

	fmt.Println("=============================Data Written=====================================")
	fmt.Println("Data Read and written successfully.")
}
