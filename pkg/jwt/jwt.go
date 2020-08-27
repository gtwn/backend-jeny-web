package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/jenywebapp/pkg/jwt/model"
)

func DecodeIDToken(IDToken string) (*model.Payload,error) {

	data := model.Payload{}
	sp := strings.Split(IDToken, ".")
	payload := sp[1]

	payloadDecoded,err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil,err
	}
	if err := json.Unmarshal(payloadDecoded, &data) ; err != nil {
		return nil,err
	}

	return &data,nil

}