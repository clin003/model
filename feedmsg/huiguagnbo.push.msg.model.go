package feedmsg

type PushMsgModel struct {
	PicURL        string `json:"picURL"  form:"picURL"`               //图片地址
	Content       string `json:"content"  form:"content"`             //文本内容
	ServerChannel string `json:"serverChannel"  form:"serverChannel"` //频道 token
	ServerToken   string `json:"serverToken"  form:"serverToken"`     //验证 token
	ServerRouter  string `json:"serverRouter"  form:"serverRouter"`   //openAPI URL
}
