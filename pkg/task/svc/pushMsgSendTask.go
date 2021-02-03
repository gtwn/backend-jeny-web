package svc

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/task/model"
)

func PushMsgSendTask(Task *model.Task,AccessToken string,Display string) error {
	
	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",AccessToken)
	msgSend := &[]model.Msg{
		{Type: "text",
		Text: "คุณส่งงาน: "+Task.Task+"\nให้คุณ: "+Task.OrderBy+"\nสถานะ: รอการตรวจสอบ\n",
	}}
	msgFollow := &[]model.Msg{
		{Type: "text",
		Text: "คุณ: "+Display+"ส่งงานให้คุณ\n กรุณาตรวจสอบงานด้วยค่ะ",
	}}
	// ผู้ส่งงาน
	pushSend := &model.PushMultiple{
		To: Task.MemberID,
		Message: *msgSend}
	// ผู้สั่งงาน
	pushToFollow := &model.PushMsg{
		To: Task.FromID,
		Message: *msgFollow,
	}
	jsonSend,err := json.Marshal(pushSend)
	if err != nil{
		return err
	}
	jsonFollow,err := json.Marshal(pushToFollow)
	if err != nil {
		return err
	}

	// ส่งหาคนสั่ง
	
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
		return err
	}	
	// ส่งหาเรา
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonSend)).Post("https://api.line.me/v2/bot/message/multicast") ; err != nil {
		return err
	}
	return nil
}