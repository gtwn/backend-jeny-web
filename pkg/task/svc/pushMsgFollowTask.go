package svc

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/jenywebapp/pkg/task/model"
)

func PushMsgFollowTask(Task *model.Task,AccessToken string,OrderID string,UserID string) error {
	
	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",AccessToken)
	// ส่งแจ้งเตือนเรา
	msgSend := &[]model.Msg{
		{Type: "text",
		Text: "คุณได้กดติดตามงาน: "+Task.Task+"\nของคุณ"+Task.OrderTo+"แล้วค่ะ",
	}}
	// ส่งตามงาน
	msgFollow := &[]model.Msg{
		{Type: "text",
		Text: "คุณ: "+Task.OrderTo+"\nกรุณาส่งงาน: "+Task.Task+"ด้วยค่ะ\n"+"จาก"+Task.OrderBy,
	}}
	// ส่งแจ้งเตือนให้เราว่ามีการกดติดตามงาน
	pushSend := &model.PushMsg{
		To: UserID,
		Message: *msgSend}
	// ส่งแจ้งเตือนให้คนที่เราตามงาน

	// กรณีไม่มี User ในระบบให้ส่งตามใน Group แทน
	// if User == nil {
	// 	fmt.Printf("user is nil"+Task.GroupID)
	// 	pushToFollow := &model.PushMsg{
	// 		To: Task.GroupID,
	// 		Message: *msgFollow,
	// 	}
	// 	jsonFollow,err := json.Marshal(pushToFollow)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// ส่งหาคนสั่ง
	
	// 	if _,err := client.R().
	// 	SetHeaders(map[string]string{
	// 		"Content-Type": "application/json",
	// 		"Authorization" : auth,
	// 	}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
	// 		return err
	// 	}	
	// } else {	// มี user ในระบบ
	// 	fmt.Printf("user is not nil"+User.UserID)
	// 	pushToFollow := &model.PushMsg{
	// 		To: OrderID,
	// 		Message: *msgFollow,
	// 	}
	// 	jsonFollow,err := json.Marshal(pushToFollow)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// ส่งหาคนสั่ง
	
	// 	if _,err := client.R().
	// 	SetHeaders(map[string]string{
	// 		"Content-Type": "application/json",
	// 		"Authorization" : auth,
	// 	}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
	// 		return err
	// 	}	
	// }

	fmt.Printf("user is not nil"+OrderID)
		pushToFollow := &model.PushMsg{
			To: OrderID,
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
		}).SetBody(string(jsonFollow)).Post("https://api.line.me/v2/bot/message/push") ; err != nil {
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