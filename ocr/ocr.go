package ocr

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"sow_ocr/text_formating"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// NewClient initializes a new OCR Client
func NewClient(visionClient *vision.ImageAnnotatorClient, formatter text_formating.TextFormatter) *Client {
	return &Client{visionClient: visionClient, formatter: formatter}
}

// Extracts first 'count' and last 'count' runes from the string
func extractStartAndEndSections(fullText string, count int) (string, string) {
	runes := []rune(fullText)
	if len(runes) <= 2*count {
		// if total length is less than twice the count just return full text as both
		return string(runes), string(runes)
	}

	return string(runes[:count]), string(runes[len(runes)-count:])
}

// PerformOCR performs OCR on the given image data and returns the recognized text
func (c *Client) PerformOCR(ctx context.Context, imageData []byte, isPaper bool) (*ExtractedPaper, error) {
	img, err := vision.NewImageFromReader(bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	response, err := c.visionClient.DetectDocumentText(ctx, img, nil)
	if err != nil {
		return nil, err
	}

	fullText, paragraphs := extractTextFromResponse(response)

	sections, err := c.formatter.ExtractSections(ctx, fullText)
	if err != nil {
		return nil, err
	}

	return &ExtractedPaper{
		Title:    sections.Title,
		Authors:  sections.Authors,
		Abstract: sections.Abstract,
		Keywords: sections.Keywords,
		Main:     paragraphs,
	}, nil
}

func (c *Client) PerformReviewerDocs(ctx context.Context, forwardAsBytes, backwardAsBytes []byte) (outForward *ExtractedValidatorForward, outBackward *ExtractedValidatorBackward, err error) {
	forwardImage, err := vision.NewImageFromReader(bytes.NewReader(forwardAsBytes))
	if err != nil {
		err = fmt.Errorf("OCR: err create a new forward image: %v", err)

		return
	}

	backwardImage, err := vision.NewImageFromReader(bytes.NewReader(backwardAsBytes))
	if err != nil {
		err = fmt.Errorf("OCR: err create a new backward image: %v", err)

		return
	}

	forwardResponse, err := c.visionClient.DetectDocumentText(ctx, forwardImage, nil)
	if err != nil {
		err = fmt.Errorf("OCR: err detect document text from forward image: %v", err)

		return
	}

	backwardResponse, err := c.visionClient.DetectDocumentText(ctx, backwardImage, nil)
	if err != nil {
		err = fmt.Errorf("OCR: err detect document text from backward image: %v", err)

		return
	}

	forwardFullText, _ := extractTextFromResponse(forwardResponse)
	backwardFullText, _ := extractTextFromResponse(backwardResponse)

	forward, err := c.formatter.ExtractDiplomaForward(ctx, forwardFullText)
	if err != nil {
		err = fmt.Errorf("OCR: err extract diploma forward info: %v", err)

		return
	}

	backward, err := c.formatter.ExtractDiplomaBackward(ctx, backwardFullText)
	if err != nil {
		err = fmt.Errorf("OCR: err extract diploma backward info: %v", err)

		return
	}

	number, err := strconv.ParseUint(forward.IssueId, 10, 64)
	if err != nil {
		err = fmt.Errorf("OCR: err parse diploma number to uint64: %v", err)

		return
	}

	forwardDate, err := time.Parse("2006-01-02", forward.IssueDate)
	if err != nil {
		err = fmt.Errorf("OCR: err parse forward date4: %v", err)

		return
	}

	backwardDate, err := time.Parse("2006-01-02", backward.IssueDate)
	if err != nil {
		err = fmt.Errorf("OCR: err parse backward date4: %v", err)

		return
	}

	outForward = &ExtractedValidatorForward{
		Number:  number,
		Date:    forwardDate,
		Science: forward.Science,
	}

	outBackward = &ExtractedValidatorBackward{
		DiplomaNumber:       backward.Number,
		DiplomaSerialNumber: backward.SerialNumber,
		Date:                backwardDate,
	}

	return
}

// extractTextFromResponse navigates the hierarchical structure of the OCR response
// and concatenates all the recognized text into a single string.
//
// The function moves through the response structure as follows:
// Pages -> Blocks -> Paragraphs -> Words -> Symbols.
// At each level, the relevant text is extracted and appended to the fullText variable.
//
// The output of this function is a string that contains the entire recognized text
// from the document.
func extractTextFromResponse(response *visionpb.TextAnnotation) (fullText string, paragraphs map[string]string) {
	paragraphs = make(map[string]string)
	i := 1

	for _, page := range response.Pages {
		for _, block := range page.Blocks {
			for _, paragraph := range block.Paragraphs {
				paragraphText := ""
				for _, word := range paragraph.Words {
					for _, symbol := range word.Symbols {
						fullText += symbol.Text
						paragraphText += symbol.Text
					}
					fullText += " "
					paragraphText += " "
				}
				fullText += "\n"
				paragraphs[fmt.Sprintf("paragraph%d", i)] = paragraphText
				i++
			}
		}
	}

	return
}
