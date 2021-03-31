package svc

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/from-api/model"
)

func GetUserIDInGroup(GroupId string,ChannelSecret string) (*model.GroupID,error) {

	client := resty.New()
	userId := model.GroupID{}

	url := fmt.Sprintf("https://api.line.me/v2/bot/group/%s/members/ids",GroupId)
	if _,err := client.R().
	SetHeader("Authorization",fmt.Sprintf("Bearer %s",ChannelSecret)).
	SetResult(&userId).Get(url); err != nil {

		return nil,err
	}

	return &userId,nil

}