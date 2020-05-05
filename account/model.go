package account

type AddReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
