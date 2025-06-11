package models

type AddRequest struct {
	Number int `json:"number" binding:"required"`
}

type AddResponse struct {
	Success bool      `json:"success"`
	List    []float64 `json:"list"`
	Message string    `json:"message,omitempty"`
}

type ListResponse struct {
	List []float64 `json:"list"`
}
