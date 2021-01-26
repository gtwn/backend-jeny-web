package svc

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/from-api/model"
)

func GetUsersProfile(UsersId *model.GroupID,ChannelSecret string) (*[]model.Profile,error) {

	client := resty.New()
	profiles := []model.Profile{}

	for _,id := range UsersId.MemberIDs {
		profile := model.Profile{}
		url := fmt.Sprintf("https://api.line.me/v2/bot/profile/%s",id)
	
		if _,err := client.R().
		SetHeader("Authorization",fmt.Sprintf("Bearer %s",ChannelSecret)).
		SetResult(&profile).Get(url); err != nil {
			return nil,err
		}

		profiles = append([]model.Profile{profile},profiles...)
	}
	

	return &profiles,nil

}