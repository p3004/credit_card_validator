package networking

type Response struct {
	CardNumber string `json:"card_number"`
	IsValid    bool   `json:"is_valid"`
	Issuer     string `json:"issuer"`
}

type Request struct {
	CardNumber string `json:"card_number"`
}

type IssuerValid struct {
	IsValidIssuer bool   `json:"is_valid_issuer"`
	Issuer        string `json:"issuer"`
}
