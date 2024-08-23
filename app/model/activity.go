package model

func AdminAddActivity(activity *Bulletin) error {

	if err := Conn.Create(activity).Error; err != nil {
		return err
	}
	return nil
}

func AdminDeleteActivity(activityID string) error {
	// 连接数据库并删除活动数据
	if err := Conn.Where("title = ?", activityID).Delete(&Bulletin{}).Error; err != nil {
		return err
	}
	return nil
}

func AdminUpdateActivity(activity *Bulletin) error {
	// 连接数据库并更新活动数据
	if err := Conn.Save(activity).Error; err != nil {
		return err
	}
	return nil
}
