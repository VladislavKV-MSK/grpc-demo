package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"sync"

	pb "demo/shared/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// In-memory хранилище задач
var (
	tasks = make(map[string]*pb.Task)
	mu    sync.RWMutex
)

type todoServer struct {
	pb.UnimplementedTodoServiceServer
}

// Генератор ID
func generateID() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generateID failed: %v", err)
	}
	return hex.EncodeToString(buf), nil
}

// Добавить задачу
func (s *todoServer) AddTask(ctx context.Context, req *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	id, err := generateID() // Функция для генерации ID (местная)
	if err != nil {
		log.Fatal(err)
	}

	task := &pb.Task{
		Id:        id,
		Title:     req.Title,
		Completed: false,
	}
	tasks[task.Id] = task

	return &pb.AddTaskResponse{Task: task}, nil
}

// Получить список задач
func (s *todoServer) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	mu.RLock()
	defer mu.RUnlock()

	var taskList []*pb.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	return &pb.GetTasksResponse{Tasks: taskList}, nil
}

// Обновить задачу
func (s *todoServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "task not found")
	}
	task.Completed = req.Completed

	return task, nil
}

// Удалить задачу
func (s *todoServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	mu.Lock()
	defer mu.Unlock()

	delete(tasks, req.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoServer{})
	log.Println("gRPC server started on :50051")
	s.Serve(lis)
}
