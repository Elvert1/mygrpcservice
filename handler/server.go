package handler

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	pb "mygprcservice/proto"
)

type Server struct {
	pb.UnimplementedFileServiceServer
}

func (s *Server) UploadFile(ctx context.Context, in *pb.UploadRequest) (*pb.UploadResponse, error) {
	fileName := in.GetName()
	content := in.GetContent()

	err := ioutil.WriteFile("files/"+fileName, content, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
		return &pb.UploadResponse{Message: "Failed to save file"}, err
	}

	return &pb.UploadResponse{Message: "File saved successfully"}, nil
}

func (s *Server) DownloadFile(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	fileName := in.GetName()

	content, err := ioutil.ReadFile("files/" + fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return &pb.DownloadResponse{}, err
	}

	return &pb.DownloadResponse{Content: content}, nil
}

func (s *Server) ListFiles(ctx context.Context, in *pb.Empty) (*pb.FilesList, error) {
	files, err := ioutil.ReadDir("files/")
	if err != nil {
		log.Fatal(err)
	}

	var filesList pb.FilesList

	for _, f := range files {
		file := &pb.File{
			Name:         f.Name(),
			CreationDate: f.ModTime().Format(time.RFC3339),
			UpdateDate:   f.ModTime().Format(time.RFC3339), // assuming that the file is not being modified after creation
		}
		filesList.Files = append(filesList.Files, file)
	}

	return &filesList, nil
}
