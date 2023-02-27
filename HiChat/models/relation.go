package models

type Relation struct {
	Model
	OwnerId  uint   //谁的关系信息
	TargetId uint   //对应的谁
	Type     int    //关系类型： 1表示好友关系 2表示群关系
	Desc     string //描述
}

func (r *Relation) RelTableName() string {
	return "relation"
}
