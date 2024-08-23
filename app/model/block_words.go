package model

import "fmt"

type BlockingWords struct {
	ID      int    `gorm:"primary_key"`
	Content string `gorm:"type:varchar(255);not null;unique"`
}

func Block() ([]string, error) {
	var blockList []BlockingWords
	if err := Conn.Find(&blockList).Error; err != nil {
		fmt.Printf("屏蔽词获取失败 err:%s", err.Error())
		return nil, err
	}

	var contentList []string
	result := Conn.Model(&BlockingWords{}).Pluck("content", &contentList)
	if result.Error != nil {
		fmt.Printf("内容获取失败 err:%s", result.Error.Error())
		return nil, result.Error
	}
	return contentList, nil

}
