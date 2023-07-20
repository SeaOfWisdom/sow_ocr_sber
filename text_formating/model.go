package text_formating

type (
	PaperSections struct {
		Title    string   `json:"title"`
		Authors  []string `json:"authors"`
		Abstract string   `json:"abstract"`
		Keywords []string `json:"keywords"`
	}

	DiplomaForward struct {
		IssueId   string `json:"issue_id"`
		IssueDate string `json:"issue_date"`
		Science   string `json:"science"`
	}

	DiplomaBackward struct {
		IssueId      string `json:"issue_id"`
		IssueDate    string `json:"issue_date"`
		Number       string `json:"diploma_number"`
		SerialNumber string `json:"diploma_serial_number"`
	}
)
