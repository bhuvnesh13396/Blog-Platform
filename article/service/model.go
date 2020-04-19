package service

type AddReq struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateReq struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
