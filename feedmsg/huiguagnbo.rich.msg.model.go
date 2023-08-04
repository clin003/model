package feedmsg

import (
	"fmt"
)

type FeedRichMsgModel struct {
	Msgtype string                `json:"msgtype"  form:"msgtype"` //rich text image video
	MsgID   string                `json:"msgID"  form:"msgID"`
	MsgTime string                `json:"msgTime"  form:"msgTime"`
	Text    FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image   FeedRichMsgImageModel `json:"image"  form:"image"`
	Video   FeedRichMsgVideoModel `json:"video"  form:"video"`
	Link    string                `json:"link"  form:"link"`
	// (Optional)
	FeedChatId int64 `json:"feed_chat_id,omitempty"`
}
type FeedRichMsgTextModel struct {
	Content         string `json:"content"  form:"content"`
	ContentEx       string `json:"contentEx"  form:"contentEx"`
	ContentExPic    string `json:"contentExPic"  form:"contentExPic"`
	ContentMarkdown string `json:"contentMarkdown"  form:"contentMarkdown"`
	ContentHTML     string `json:"contentHTML"  form:"contentHTML"`
}
type FeedRichMsgImageModel struct {
	PicURL   string `json:"picURL"  form:"picURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgVideoModel struct {
	FileURL  string `json:"fileURL"  form:"fileURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}

func (msg *FeedRichMsgModel) ToString() (res string) {
	res = fmt.Sprintf("msgID:%s,msgType:%s,msgTime:%s\n", msg.MsgID, msg.Msgtype, msg.MsgTime)
	if len(msg.Text.Content) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Text.Content)
	}
	if len(msg.Image.PicURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Image.PicURL)
	}
	if len(msg.Video.FileURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Video.FileURL)
	}

	return
}

func NewFeedRichMsg(msgType, msgID, msgTime, textContent, textContentEx, textContentExPic, imagePicURL, imageFilePath string) FeedRichMsgModel {
	return FeedRichMsgModel{
		Msgtype: msgType,
		MsgID:   msgID,
		MsgTime: msgTime,
		Text: FeedRichMsgTextModel{
			Content:      textContent,
			ContentEx:    textContentEx,
			ContentExPic: textContentExPic,
		},
		Image: FeedRichMsgImageModel{
			PicURL:   imagePicURL,
			FilePath: imageFilePath,
		},
	}
}
func NewFeedRichMsgV2(msgType, msgID, msgTime, textContent, textContentEx, textContentExPic, imagePicURL, imageFilePath, link string) FeedRichMsgModel {
	return FeedRichMsgModel{
		Msgtype: msgType,
		MsgID:   msgID,
		MsgTime: msgTime,
		Text: FeedRichMsgTextModel{
			Content:      textContent,
			ContentEx:    textContentEx,
			ContentExPic: textContentExPic,
		},
		Image: FeedRichMsgImageModel{
			PicURL:   imagePicURL,
			FilePath: imageFilePath,
		},
		Link: link,
	}
}

func NewFeedRichMsgV3(msgType, msgID, msgTime, textContent, textContentEx, textContentExPic, textContentMarkdown, imagePicURL, imageFilePath, imageCaption, videoFileURL, videoFilePath, videoCaption, link string) FeedRichMsgModel {
	return FeedRichMsgModel{
		Msgtype: msgType,
		MsgID:   msgID,
		MsgTime: msgTime,
		Text: FeedRichMsgTextModel{
			Content:         textContent,
			ContentEx:       textContentEx,
			ContentExPic:    textContentExPic,
			ContentMarkdown: textContentMarkdown,
		},
		Image: FeedRichMsgImageModel{
			PicURL:   imagePicURL,
			FilePath: imageFilePath,
			// (Optional)
			Caption: imageCaption,
		},
		Video: FeedRichMsgVideoModel{
			FileURL:  videoFileURL,
			FilePath: videoFilePath,
			// (Optional)
			Caption: videoCaption,
		},
		Link: link,
	}
}
