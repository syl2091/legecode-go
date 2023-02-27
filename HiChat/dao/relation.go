package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
)

// 获取好友信息
func FriendList(userId uint) (*[]models.UserBasic, error) {
	relations := make([]models.Relation, 0)
	if tx := global.DB.Where("owner_id = ?", userId).Find(&relations); tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation数据")
		return nil, errors.New("未查到好友关系")
	}

	userID := make([]uint, 0)
	for _, v := range relations {
		userID = append(userID, v.TargetId)
	}
	user := make([]models.UserBasic, 0)
	if tx := global.DB.Where("id in ?", userID).Find(&user); tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation好友关系")
		return nil, errors.New("未查到好友")
	}
	return &user, nil
}

// 通过id添加好友
func AddFriend(userId, targetId uint) (int, error) {
	if userId == targetId {
		return -2, errors.New("userID和targetId相等")
	}
	//通过id查询用户
	targetUser, err := FindUserID(targetId)
	if err != nil {
		return -1, errors.New("未查询到用户")
	}
	if targetUser.ID == 0 {
		zap.S().Info("未查询到用户")
		return -1, errors.New("未查询到用户")
	}
	relation := models.Relation{}
	if tx := global.DB.Where("owner_id = ? and target_id = ?", userId, targetId).First(&relation); tx.RowsAffected == 1 {
		zap.S().Info("该好友存在")
		return 0, errors.New("好友已经存在")
	}
	//开启事务
	tx := global.DB.Begin()
	relation.OwnerId = userId
	relation.TargetId = targetUser.ID
	relation.Type = 1
	if t := tx.Create(&relation); t.RowsAffected == 0 {
		zap.S().Info("创建失败")
		tx.Rollback()
		return 0, errors.New("创建好友记录失败")
	}

	//提交事务
	tx.Commit()
	return 1, nil
}

// 通过昵称添加好友
func AddFriendByName(userId uint, name string) (int, error) {
	user, err := FindUserByName(name)
	if err != nil {
		return -1, errors.New("该用户不存在")
	}
	if userId == 0 {
		zap.S().Info("未查询到用户")
		return -1, errors.New("该用户不存在")
	}
	return AddFriend(userId, user.ID)
}
