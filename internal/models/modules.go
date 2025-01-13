package models

type Request struct {
	Status       string `json:"status"`
	AdminComment string `json:"admin_comment"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateStatusResponse struct {
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}
