package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetRooms(page, pageSize int) ([]Houseinfo, error) {
	fmt.Printf("page: %d, pageSize: %d\n", page, pageSize)
	var ret []Houseinfo // 注意修改返回值类型为 []House
	err := Conn.Limit(pageSize).Offset((page - 1) * pageSize).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// GetRoomByID 定义数据库操作方法
func GetRoomByID(id int) (*Houseinfo, error) {
	var room Houseinfo
	if err := Conn.Table("houseinfo").Where("ID = ?", id).First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func AdminAddRoom(room *House) error {
	result := Conn.Create(room)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AdminDeleteRoom(Name string) error {
	// 根据房间名查询房间信息
	var room House
	if err := Conn.Where("name = ?", Name).First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("房间不存在")
		}
		return err
	}

	// 删除房间信息
	if err := Conn.Delete(&room).Error; err != nil {
		return err
	}

	return nil
}

func AdminUpdateRoom(room *House) error {
	// 根据房间ID查询房间信息
	var existingRoom House
	if err := Conn.First(&existingRoom, room.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("房间不存在")
		}
		return err
	}

	// 更新房间信息
	if err := Conn.Model(&existingRoom).Updates(room).Error; err != nil {
		return err
	}

	return nil
}
