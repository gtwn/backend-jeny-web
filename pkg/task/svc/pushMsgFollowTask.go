package svc

import (
	"encoding/json"
	"fmt"
	"strings"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/task/model"
)

func PushMsgFollowTask(Task *model.Task,AccessToken string,UserID string) error {
	
	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",AccessToken)
	// ส่งแจ้งเตือนคนสั่งงาน
	msgSend := &[]model.Msg{
		{Type: "text",
		Text: "คุณได้กดติดตามงาน: "+Task.Task+"\nของคุณ"+strings.Join(Task.Member[:],",")+"แล้วค่ะ",
	}}
	// ส่งตามงานคนถูกสั่งงาน
	msgFollow := &[]model.Msg{
		{Type: "text",
		Text: "คุณ: "+strings.Join(Task.Member[:],",")+"\nกรุณาส่งงาน: "+Task.Task+"ด้วยค่ะ\n"+"จาก"+Task.OrderBy,
	}}
	// ส่งแจ้งเตือนให้เราว่ามีการกดติดตามงาน
	pushSend := &model.PushMsg{
		To: UserID,
		Message: *msgSend}

	pushToFollow := &model.PushMultiple{
		To: Task.MemberID,
		Message: *msgFollow,
	}
	jsonFollow,err := json.Marshal(pushToFollow)
	if err != nil {
		return err
	}

	// ส่งหาคนโดนสั่ง
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/multicast") ; err != nil {
		return err
	}	
	
	jsonSend,err := json.Marshal(pushSend)
	if err != nil{
		return err
	}
	
	// ส่งหาเรา
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonSend)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
		return err
	}
	
	return nil
}