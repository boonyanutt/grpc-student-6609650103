package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb" // เรียกใช้โค้ดที่คอมสร้างให้

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

// GetStudent: คืนค่าข้อมูลนักศึกษา 1 คน (Task 2: เพิ่มเบอร์โทรศัพท์)
func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {
	log.Printf("Received request for ID: %d", req.Id)
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
		Phone: "081-234-5678", // Task 2
	}, nil
}

// ListStudents: คืนค่ารายชื่อนักศึกษาทุกคน (Task 1)
func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {
	log.Println("Received request for ListStudents")

	students := []*pb.StudentResponse{
		{Id: 101, Name: "Alice Johnson", Major: "CS", Email: "alice@tu.ac.th", Phone: "081-111-1111"},
		{Id: 102, Name: "Bob Smith", Major: "CS", Email: "bob@tu.ac.th", Phone: "082-222-2222"},
	}

	return &pb.StudentListResponse{Student: students}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // รันบนพอร์ต 50051 [cite: 172]
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &server{})
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
