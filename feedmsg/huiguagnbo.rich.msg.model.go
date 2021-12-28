package feedmsg

import (
	// "encoding/json"
	"fmt"
	// "strings"
	// "sync"
	// "time"
	// "gitee.com/lyhuilin/open_api/model/kvdb"
	// "gitee.com/lyhuilin/open_api/pkg/errno"
	// "gitee.com/lyhuilin/open_api/pkg/wsserver"
	// "gitee.com/lyhuilin/util"
)

type FeedRichMsgModel struct {
	Msgtype string                `json:"msgtype"  form:"msgtype"` //rich text image
	MsgID   string                `json:"msgID"  form:"msgID"`
	MsgTime string                `json:"msgTime"  form:"msgTime"`
	Text    FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image   FeedRichMsgImageModel `json:"image"  form:"image"`
	Link    string                `json:"link"  form:"link"`
}
type FeedRichMsgTextModel struct {
	Content      string `json:"content"  form:"content"`
	ContentEx    string `json:"contentEx"  form:"contentEx"`
	ContentExPic string `json:"contentExPic"  form:"contentExPic"`
}
type FeedRichMsgImageModel struct {
	PicURL   string `json:"picURL"  form:"picURL"`
	FilePath string `json:"filePath"  form:"filePath"`
}

func (msg *FeedRichMsgModel) ToString() (res string) {
	res = fmt.Sprintf("msgID:%s,msgType:%s,msgTime:%s\n", msg.MsgID, msg.Msgtype, msg.MsgTime)
	if len(msg.Text.Content) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Text.Content)
	}
	if len(msg.Image.PicURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Image.PicURL)
	}
	// if len(msg.Link) > 0 {
	// 	res = fmt.Sprintf("%s\n%s", res, msg.Link)
	// }
	return
}
