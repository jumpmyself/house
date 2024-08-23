package model

// GetAllMessages 获取所有留言
func GetAllMessages() ([]Messages, error) {
	var messages []Messages
	if err := Conn.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
