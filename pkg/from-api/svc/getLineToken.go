package svc

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/labstack/echo/v4"
)

func GetLineToken(LineAPI string,ChannelID string,ChannelSecret string,c echo.Context) (*model.Profile,error) {

	client := resty.New()
	code := c.QueryParam("code")
	fmt.Println("code : ",code)
	fmt.Println("ID : ",ChannelID)
	fmt.Println("Secret  : ",ChannelSecret)

	authSuccess := model.AuthSuccess{}
	resp,err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  "",
		"client_id":     ChannelID,
		"client_secret": ChannelSecret,
		}).
		SetResult(&authSuccess). // or SetResult(AuthSuccess{}).
		Post("https://api.line.me/oauth2/v2.1/token") 
	if err != nil {
			return nil,err
	}
	fmt.Println(resp)
	fmt.Println("AccessToken: ",authSuccess.AccessToken)
	profile := model.Profile{}
	if _,err := client.R().
		SetHeader("Authorization", "Bearer "+authSuccess.AccessToken).
		SetResult(profile). // or SetResult(AuthSuccess{}).
		Get("https://api.line.me/v2/profile") ; err != nil {
			return nil,err
	}

	return &profile,nil
}