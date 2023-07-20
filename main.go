package main

import (
	"context"
	"log"

	"sow_ocr/config"
	"sow_ocr/ocr"
	"sow_ocr/text_formating"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func main() {
	cfg := config.NewConfig()

	visionClient, err := vision.NewImageAnnotatorClient(context.Background(), option.WithCredentialsFile(cfg.VisionFilepath))
	if err != nil {
		log.Fatal(err)
	}

	formatter := text_formating.New(cfg.OpenAIApiKey)

	ocrClient := ocr.NewClient(visionClient, formatter)

	ocr.RunGRPCServer(ocrClient, cfg.GrpcPort)
}
