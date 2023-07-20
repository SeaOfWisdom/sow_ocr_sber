package text_formating

import (
	"context"
	"encoding/json"

	openai "github.com/sashabaranov/go-openai"
)

func (f *formatter) ExtractSections(ctx context.Context, fullText string) (*PaperSections, error) {
	response, err := f.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: paperPrompt + fullText + EndOfPaperPrompt,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	// The response from GPT-3 should be a JSON-like string, so we try to unmarshal it into a PaperSections struct
	var sections PaperSections
	err = json.Unmarshal([]byte(response.Choices[0].Message.Content), &sections)
	if err != nil {
		return nil, err
	}

	return &sections, nil
}
