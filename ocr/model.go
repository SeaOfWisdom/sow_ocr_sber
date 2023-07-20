package ocr

import (
	vision "cloud.google.com/go/vision/apiv1"
	pb "github.com/SeaOfWisdom/sow_proto/ocr-srv"
	"sow_ocr/text_formating"
	"time"
)

type Client struct {
	visionClient *vision.ImageAnnotatorClient
	formatter    text_formating.TextFormatter
}

type Server struct {
	pb.UnimplementedOCRServer
	ocrClient *Client
}

type ExtractedPaper struct {
	Title    string            `json:"title"`
	Authors  []string          `json:"authors"`
	Abstract string            `json:"abstract"`
	Keywords []string          `json:"keywords"`
	Main     map[string]string `json:"main"`
}

type (
	ExtractedValidatorForward struct {
		Number  uint64    `json:"number"`
		Date    time.Time `json:"date"`
		Science string    `json:"science"`
	}

	ExtractedValidatorBackward struct {
		Date                time.Time `json:"date"`
		DiplomaSerialNumber string    `json:"diploma_serial_number"`
		DiplomaNumber       string    `json:"diploma_number"`
	}
)
