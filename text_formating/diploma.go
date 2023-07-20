package text_formating

import (
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
)

func (f *formatter) ExtractDiplomaForward(ctx context.Context, fullText string) (*DiplomaForward, error) {
	response, err := f.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: diplomaForwardPrompt + fullText + EndOfPaperPrompt,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	// The response from GPT-3 should be a JSON-like string, so we try to unmarshal it into a PaperSections struct
	var diplomaForward DiplomaForward
	err = json.Unmarshal([]byte(response.Choices[0].Message.Content), &diplomaForward)
	if err != nil {
		return nil, err
	}

	return &diplomaForward, nil
}

func (f *formatter) ExtractDiplomaBackward(ctx context.Context, fullText string) (*DiplomaBackward, error) {
	response, err := f.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: diplomaBackwardPrompt + fullText + EndOfPaperPrompt,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	// The response from GPT-3 should be a JSON-like string, so we try to unmarshal it into a PaperSections struct
	var diplomaBackward DiplomaBackward
	err = json.Unmarshal([]byte(response.Choices[0].Message.Content), &diplomaBackward)
	if err != nil {
		return nil, err
	}

	return &diplomaBackward, nil
}
