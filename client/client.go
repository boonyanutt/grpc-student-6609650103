package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetStudent(ctx, &pb.StudentRequest{Id: 101})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Phone: %s", res.Phone)

	listRes, err := c.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	marshal, err := json.Marshal(listRes)
	if err != nil {
		log.Fatalf("Error marshalling students: %v", err)
		return
	}

	log.Printf("Students: %s", marshal)
}
