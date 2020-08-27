package model

type Payload struct {
	Iss 	string 		`json:"iss"`
	Sub		string 		`json:"sub"`
	Aud		string 		`json:"aud"`
	Exp		int64 		`json:"exp"`
	Iat		int64		`json:"iat"`
	Amr		[]string	`json:"amr"`
	Name	string		`json:"name"`
	Picture	string		`json:"picture"`
}