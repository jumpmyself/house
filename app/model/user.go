package model

import (
	"fmt"
	"github.com/go-pay/util/snowflake"
	"strconv"
	"strings"
	"time"
)

func GetUser(name string) *User {
	var ret User
	if err := Conn.Table("user").Where("username=?", name).Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}

// GetUserV1  原生sql优化
func GetUserV1(name string) *User {
	var ret User
	err := Conn.Raw("select * from user where name = ? limit 1", name).Scan(&ret).Error
	if err != nil {
		return nil
	}
	return &ret
}

func CreateUser(user *User) error {
	if err := Conn.Create(user).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}

// CreateUserHouse 创建用户订房订单信息
func CreateUserHouse(data *UserHouse) error {
	// 调用数据库连接的 Create 方法创建用户订房订单信息
	if err := Conn.Create(data).Error; err != nil {
		return err
	}

	return nil
}

// CreateUserProduct 创建用户订购商品订单信息
func CreateUserProduct(data *UserProduct) error {
	// 执行插入订单数据
	if err := Conn.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID int) (User, error) {
	var user User
	if err := Conn.Where("id = ?", userID).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// GetRoomBySearch 根据搜索条件返回符合条件的房间列表
func GetRoomBySearch(roomType string) []*Houseinfo {
	var rooms []*Houseinfo
	// 使用 GORM 查询数据库，检查 name 字段是否包含 roomType，并返回符合条件的房间列表
	Conn.Where("name LIKE ?", "%"+roomType+"%").Find(&rooms)
	return rooms
}

func AdminAddUser(newUser *User) error {
	// 连接数据库并插入用户数据
	if err := Conn.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}
func AdminDeleteUser(userID string) error {
	// 连接数据库并根据用户ID删除用户数据
	if err := Conn.Where("username = ?", userID).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func AdminUpdateUser(user *User) error {
	// 连接数据库并更新用户数据
	if err := Conn.Model(&User{}).Where("Username = ?", user.Username).Updates(map[string]interface{}{"Username": user.Username, "Password": user.Password, "Email": user.Email, "Telephone": user.Telephone, "City": user.City, "Age": user.Age}).Error; err != nil {
		return err
	}
	return nil
}

func GetRoomPrice(roomId string) (name string, price int64, err error) {
	var room Houseinfo
	if err := Conn.Where("house_id = ?", roomId).First(&room).Error; err != nil {
		// 处理数据库查询错误
		return "", 0, err
	}

	// 输出调试信息
	fmt.Println("房间名:", room.Name)
	fmt.Println("价格:", room.Price)

	return room.Name, int64(room.Price), nil
}

func BuyRoom(name, UserName, HouseId, HouseName, Phone, Notes, Num, Price, TotalPriceStr, CreatTime string) error {
	num, _ := strconv.Atoi(Num)
	totalPrice, _ := strconv.ParseFloat(TotalPriceStr, 64)
	houseId, _ := strconv.Atoi(HouseId)

	// 创建一个新的雪花节点
	node, err := snowflake.NewNode(1) // 这里的参数是节点的 ID
	if err != nil {
		fmt.Println("uuid创建失败:", err)
		return nil
	}

	// 生成一个唯一的 ID
	uuid := node.Generate().Int64()

	// 移除单位部分
	priceStr := strings.Replace(Price, "元/天", "", 1)

	// 将价格部分转换为浮点数
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("解析价格失败:", err)
		return nil
	}

	fmt.Println("价格:", price)

	order := UserHouse{
		UUID:        uuid,
		UserName:    UserName,
		Name:        name,
		HouseId:     int64(houseId),
		HouseName:   HouseName,
		Phone:       Phone,
		Notes:       Notes,
		Num:         int64(num),
		Price:       price,
		TotalPrice:  totalPrice,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	// 插入数据到数据库
	if err := Conn.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

// GetUserOrdersByName 根据用户名获取订单信息
func GetUserOrdersByName(name string) ([]UserHouse, error) {
	var orders []UserHouse
	// 查询数据库，根据用户名查询订单信息
	if err := Conn.Where("name = ?", name).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetUserProductOrdersByName 根据用户名获取订单信息
func GetUserProductOrdersByName(name string) ([]UserProduct, error) {
	var orders []UserProduct
	// 查询数据库，根据用户名查询订单信息
	if err := Conn.Where("username = ?", name).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func UserRegister(name, password, email string) error {
	// 将用户数据保存到数据库
	result := Conn.Create(&User{
		Username: name,
		Password: password,
		Email:    email,
	})

	// 检查保存过程中是否发生错误
	if result.Error != nil {
		return result.Error
	}

	return nil
}
