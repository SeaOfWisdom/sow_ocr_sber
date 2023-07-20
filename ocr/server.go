package ocr

import (
	"context"
	"fmt"
	pb "github.com/SeaOfWisdom/sow_proto/ocr-srv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

func (s *Server) ExtractText(ctx context.Context, in *pb.ExtractTextRequest) (*pb.ExtractTextResponse, error) {
	isPaper := in.IsPaper
	paper, err := s.ocrClient.PerformOCR(ctx, in.GetImage(), isPaper)
	if err != nil {
		return nil, err
	}
	return &pb.ExtractTextResponse{
		Title:    paper.Title,
		Authors:  paper.Authors,
		Abstract: paper.Abstract,
		Keywords: paper.Keywords,
		Main:     paper.Main,
	}, nil
}

func (s *Server) ExtractValidatorText(ctx context.Context, in *pb.ExtractValidatorRequest) (out *pb.ExtractValidatorResponse, err error) {
	forward, backward, err := s.ocrClient.PerformReviewerDocs(ctx, in.ForwardImage, in.BackwardImage)
	if err != nil {
		return
	}

	out = &pb.ExtractValidatorResponse{
		BackwardInfo: &pb.ValidatorBackwardInfo{
			Number:       backward.DiplomaNumber,
			SerialNumber: backward.DiplomaSerialNumber,
			Date:         timestamppb.New(backward.Date),
		},
		ForwardInfo: &pb.ValidatorForwardInfo{
			Number:   forward.Number,
			Sciences: forward.Science,
			Date:     timestamppb.New(forward.Date),
		},
	}

	return
}

func RunGRPCServer(ocrClient *Client, port uint64) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOCRServer(s, &Server{ocrClient: ocrClient})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
