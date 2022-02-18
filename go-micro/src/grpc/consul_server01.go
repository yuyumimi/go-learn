package main

import (
	"fmt"
	"google.golang.org/grpc"
	"idea-test-project/src/protobuf/protoc"
	"log"
	"net"
)

func main() {

	fmt.Println("hll")

}

type ManTeacher struct {
}

func (t *ManTeacher) SayName(ctx context.Context, teacher *protoc.Teacher) (*protoc.Teacher, error) {
	teacher.Name += " say"
	return teacher, nil

}
func (t *ManTeacher) GetCourse(ctx context.Context, teacher *protoc.Teacher) (*protoc.Teacher, error) {
	teacher.Courses[0] = "english"
	return teacher, nil

}
func grpcServer() {
	server := grpc.NewServer()
	proto.RegisterSayNameServer(server, new(ManTeacher))

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	server.Serve(listener)
}
