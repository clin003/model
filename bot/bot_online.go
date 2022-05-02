package bot

import (
	"gitee.com/lyhuilin/model"
)

// 机器人在线表模型
type BotOnline struct {
	model.Model

	// 机器人类型：wx qq
	Botype string `json:"botype" form:"botype"  gorm:"column:botype;"`
	// 机器人id
	Botid string `json:"botid" form:"botid"  gorm:"column:botid;"`
	// 机器人token 等同 OwnerUUID
	Botoken string `json:"botoken" form:"botoken"  gorm:"column:botoken;"`
	// 最后一次登录时间
	LastloginTime string `json:"lastlogin_time" form:"lastlogin_time"  gorm:"column:lastlogin_time;"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}

// // 表名
// func (c *BotOnline) TableName() string {
// 	return "bot_online"
// }

// // 创建
// func (c *BotOnline) Create() error {
// 	return model.DB.Self.Create(&c).Error
// }

// // 更新
// func (c *BotOnline) Save() error {
// 	if d, err := c.Get(); err != nil {
// 		return c.Create()
// 	} else {
// 		c.ID = d.ID
// 	}
// 	return c.Update()
// }

// // 删除
// func (c *BotOnline) Delete() error {
// 	switch {
// 	case c.ID > 0:
// 		return model.DB.Self.Where("botoken=?", c.Botoken).Delete(&c).Error
// 	default:
// 		return model.DB.Self.Where("botid = ?  AND botoken=?", c.Botid, c.Botoken).Delete(&BotOnline{}).Error
// 	}
// }

// // 改
// func (c *BotOnline) Update() error {
// 	return model.DB.Self.Model(&c).Omit("botoken").Omit("botid").Where("botoken = ?", c.Botoken).Updates(&c).Error
// }

// // 获取名单信息
// func (c *BotOnline) Get() (*BotOnline, error) {
// 	u := &BotOnline{}

// 	switch {
// 	case len(c.Botoken) > 0 && len(c.Botid) > 0:
// 		d := model.DB.Self.Where("botid = ? AND botoken = ?", c.Botid, c.Botoken).First(&u)
// 		return u, d.Error
// 	default:
// 		d := model.DB.Self.Where("botid = ?", c.Botid).First(&u)
// 		return u, d.Error
// 	}
// }

// // func (c *BotOnline) IsHave() bool {
// // 	var count int64

// // 	switch {
// // 	case len(c.Botoken) > 0 && len(c.AccountId) > 0:
// // 		model.DB.Self.Model(&BotOnline{}).Where("botoken = ? AND account_id = ?", c.Botoken, c.AccountId).Count(&count)
// // 		return count > 0
// // 	default:
// // 		model.DB.Self.Model(&BotOnline{}).Where("account_id = ?", c.AccountId).Count(&count)
// // 		return count > 0
// // 	}
// // }

// // 获取名单列表
// func (c *BotOnline) GetList() ([]*BotOnline, int64, error) {
// 	var count int64
// 	if c.Size <= 0 || c.Size > 100 {
// 		c.Size = 10
// 	}
// 	if c.Page <= 0 {
// 		c.Page = -1
// 	}
// 	size := c.Size
// 	offset := c.Page - 1
// 	if offset > 0 {
// 		offset = offset * size
// 	}

// 	list := make([]*BotOnline, 0)

// 	switch {
// 	case len(c.Botoken) > 0 && len(c.Filter) > 0:
// 		err := model.DB.Self.Select("id", "botype", "botid", "lastlogin_time").Where("botoken = ? AND botid = ?", c.Botoken, c.Filter).Limit(size).Offset(offset).Find(&list).Error
// 		model.DB.Self.Model(&BotOnline{}).Where("botoken = ? AND account_id = ?", c.Botoken, c.Filter).Count(&count)
// 		return list, count, err

// 	case len(c.Botoken) > 0:
// 		err := model.DB.Self.Select("id", "botype", "botid", "lastlogin_time").Where("botoken = ?", c.Botoken).Limit(size).Offset(offset).Find(&list).Error
// 		model.DB.Self.Model(&BotOnline{}).Where("botoken = ?", c.Botoken).Count(&count)
// 		return list, count, err

// 	default:
// 		err := model.DB.Self.Select("id", "botype", "botid", "lastlogin_time").Limit(size).Offset(offset).Find(&list).Error
// 		model.DB.Self.Model(&BotOnline{}).Count(&count)
// 		return list, count, err
// 	}
// }
