package text_formating

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

const (
	paperPrompt = "From the full text of the following scientific paper, extract the title, authors, abstract and keywords, and" +
		" return them in a json-like format, such as {'title': <title>, 'authors':[<author1>, ..., <authorN>], 'abstract': <abstract>, 'keywords':[<keyword1>, ..., <keywordN>]}"
	EndOfPaperPrompt      = "\n\nEnd of paper. "
	diplomaForwardPrompt  = "Твоя задача найти следующую информацию в тексте: какого числа присуждена(формат даты: год-месяц-день) ученая степень, какой номер приказа(например №10, eNo 10. Верни только цифру) и по какой науке присуждена ученая степень.В случае, если название науки представлено некорректно или непонятно, найди наиболее подходящее соответствие из списка возможных наук (например, 'педагогические науки', 'физические науки', 'химические науки', 'биологические науки' и т.д.). Верни эту информацию исключительно в json формате в виде {\"issue_date\": <дата выдачи диплома>, \"issue_id\": <номер приказа>, \"science\": <наука>} без дополнительного текста"
	diplomaBackwardPrompt = "Твоя задача найти следующую информацию в тексте: номер диплома, серия диплома, номер и дату приказа(формат даты: год-месяц-день). Верти эту информацию в json формате в виде {\"diploma_number\": <номер диплома>, \"diploma_serial_number\": <серийный номер диплома>, \"issue_id\": <номер приказа>, \"issue_date\": <дата приказа>} без дополнительного текста"
)

type TextFormatter interface {
	ExtractSections(ctx context.Context, fullText string) (*PaperSections, error)
	ExtractDiplomaForward(ctx context.Context, fullText string) (*DiplomaForward, error)
	ExtractDiplomaBackward(ctx context.Context, fullText string) (*DiplomaBackward, error)
}

type formatter struct {
	client *openai.Client
}

func New(apiKey string) TextFormatter {
	return &formatter{
		client: openai.NewClient(apiKey),
	}
}
