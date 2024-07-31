package networking

type Response struct {
	CardNumber string `json:"card_number"`
	IsValid    bool   `json:"is_valid"`
}

type Request struct {
	CardNumber string `json:"card_number"`
}
