package svc

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"strings"
	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/task/model"
)

func PushMsgDoneTask(Task *model.Task,AccessToken string) error {
	
	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",AccessToken)
	// ส่งแจ้งผู้สั่งงาน
	msgSend := &[]model.Msg{
		{Type: "text",
		Text: "คุณตรวจงาน: "+Task.Task+"\nให้คุณ: @"+strings.Join(Task.Member[:]," @")+"\nสถานะ: ผ่านการตรวจสอบ\n",
	}}
	// ส่งแจ้งเตือนหาผู้ส่งงาน
	msgFollow := &[]model.Msg{
		{Type: "text",
		Text: "คุณ: "+Task.OrderBy+"ตรวจงาน: "+Task.Task+"\nสถานะ: ผ่านการตรวจสอบ",
	}}
	// ส่งแจ้งผู้สั่งงาน
	pushSend := &model.PushMsg{
		To: Task.FromID,
		Message: *msgSend,
	}
	jsonSend,err := json.Marshal(pushSend)
	if err != nil{
		return err
	}

	// ส่งแจ้งเตือนหาผู้ส่งงาน
	pushToFollow := &model.PushMultiple{
		To: Task.MemberID,
		Message: *msgFollow,
	}
	jsonFollow,err := json.Marshal(pushToFollow)
	if err != nil {
		return err
	}

	// ส่งหาคนสั่งงาน
	
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonSend)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
		return err
	}	
	

	// ส่งหาคนส่งงาน
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Authorization" : auth,
	}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/multicast") ; err != nil {
		return err
	}
	
	return nil
}