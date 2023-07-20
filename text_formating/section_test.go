package text_formating_test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	pb "github.com/SeaOfWisdom/sow_proto/ocr-srv"

	"google.golang.org/grpc"
)

func TestOCRService(t *testing.T) {
	visionCredentials := os.Getenv("VISION_CREDENTIALS")
	if visionCredentials == "" {
		t.Fatalf("Environment variable VISION_CREDENTIALS not set")
	}

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		t.Fatalf("Environment variable OPENAI_API_KEY not set")
	}

	// Create gRPC client
	conn, err := grpc.Dial("ocr:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOCRClient(conn)

	// Read pdf file
	imageBytes, err := ioutil.ReadFile("/app/test_data/test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	// Send image to the service and get the response
	resp, err := client.ExtractText(context.Background(), &pb.ExtractTextRequest{Image: imageBytes})
	if err != nil {
		t.Fatalf("Could not send document to service: %v", err)
	}

	// Log the response for visual evaluation
	log.Printf("Received response: %+v", resp)
}
