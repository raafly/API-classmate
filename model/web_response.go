package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type StudentResponse struct {
	Nama	string	`json:"nama"`
	Absen	int		`json:"absen"`
	Gender	string	`json:"gender"`
	Nis		int		`json:"nis"`
}