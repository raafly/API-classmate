package model

type StudentCreateRequest struct {
	Nama	string	`json:"nama"`
	Absen	int		`json:"absen"`
	Gender 	string	`json:"gender"`
	Nis		int		`json:"nis"`
}

type StudentCreateResponse struct {
	Nama	string	`json:"nama"`
	Absen	int		`json:"absen"`
	Gender 	string	`json:"gender"`
	Nis		int		`json:"nis"`
}