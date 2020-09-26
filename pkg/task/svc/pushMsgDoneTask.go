package svc

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/task/model"
)

func PushMsgDoneTask(Task *model.Task,AccessToken string,UserID string) error {
	
	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",AccessToken)
	// ส่งแจ้งเตือนเรา
	msgSend := &[]model.Msg{
		{Type: "text",
		Text: "คุณตรวจงาน: "+Task.Task+"\nให้คุณ: "+Task.OrderTo+"\nสถานะ: ผ่านการตรวจสอบ\n",
	}}
	// ส่งแจ้งเตือนอีกคน
	msgFollow := &[]model.Msg{
		{Type: "text",
		Text: "คุณ: "+Task.OrderBy+"ตรวจงาน: "+Task.Task+"\nสถานะ: ผ่านการตรวจสอบ",
	}}
	// ส่งหาเรา
	pushSend := &model.PushMsg{
		To: Task.FromID,
		Message: *msgSend}
	// ส่งหาคนส่งงาน
	pushToFollow := &model.PushMsg{
		To: UserID,
		Message: *msgFollow,
	}
	// spew.Dump(msgFollow)
	jsonSend,err := json.Marshal(pushSend)
	if err != nil{
		return err
	}
	fmt.Printf(string(jsonSend))
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
	}).SetBody(string(jsonSend)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
		return err
	}
	return nil
}