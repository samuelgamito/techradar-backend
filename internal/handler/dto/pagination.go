package dto

type (
	Pagination struct {
		Content  interface{} `json:"content"`
		Page     int         `json:"page"`
		PageSize int         `json:"page_size"`
		Total    int         `json:"total"`
	}
)
