package svc

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/from-api/model"
	md "github.com/jenywebapp/pkg/jwt/model"
	"github.com/jenywebapp/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func GetLineToken(LineAPI string,ChannelID string,ChannelSecret string,c echo.Context) (*model.AuthSuccess,*md.Payload,*model.Profile,error) {

	client := resty.New()
	code := c.QueryParam("code")
	url := c.Request().Host
	redirectURI := fmt.Sprintf("https://%s/callback",url)
	

	authSuccess := model.AuthSuccess{}
	if _,err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  redirectURI,
		"client_id":     ChannelID,
		"client_secret": ChannelSecret,
		}).
		SetResult(&authSuccess). // or SetResult(AuthSuccess{}).
		Post("https://api.line.me/oauth2/v2.1/token") ; err != nil {
			return nil,nil,nil,err
	}
	
	fmt.Println("AccessToken: ",authSuccess.IDToken)
	profile := model.Profile{}
	if _,err := client.R().
		SetHeader("Authorization", "Bearer "+authSuccess.AccessToken).
		SetResult(&profile). // or SetResult(AuthSuccess{}).
		Get("https://api.line.me/v2/profile") ; err != nil {
			return nil,nil,nil,err
	}

	payload,err := jwt.DecodeIDToken(authSuccess.IDToken) 
	if err != nil{
		return nil,nil,nil,err
	}



	return &authSuccess,payload,&profile,nil
}