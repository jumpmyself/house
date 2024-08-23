package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetAdminByName(name string) *Admin {
	var ret Admin
	if err := Conn.Table("admin").Where("username=?", name).Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}

//func AdminGetRooms() []Houseinfo {
//	var ret []Houseinfo
//	if err := Conn.Table("Houseinfo").Find(&ret).Error; err != nil {
//		fmt.Printf("err:%s", err.Error())
//		fmt.Println("Houseinfo:", ret)
//	}
//	return ret
//}

func AdminGetRooms1() []House {
	var ret []House
	if err := Conn.Table("House").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		fmt.Println("House:", ret)
	}
	return ret
}

func AdminGetActivities() []Bulletin {
	var ret []Bulletin
	if err := Conn.Table("bulletin").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func AdminGetProducts() []Products {
	var ret []Products
	if err := Conn.Table("Products").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func AdminGetUsers() []User {
	var ret []User
	if err := Conn.Table("user").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func AdminGetHouseOrder() []UserHouse {
	var ret []UserHouse
	if err := Conn.Table("user_house").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func AdminGetProductOrder() []UserProduct {
	var ret []UserProduct
	if err := Conn.Table("user_product").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析HTTP请求中的文件
	err := r.ParseMultipartForm(10 << 20) // 设置最大文件大小为10MB
	if err != nil {
		fmt.Println("Error parsing form:", err)
		return
	}

	// 获取上传的文件
	files := r.MultipartForm.File["Images"]
	for _, file := range files {
		// 打开上传的文件
		src, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer src.Close()

		// 创建一个新文件来保存上传的文件内容
		dst, err := os.Create("images/" + file.Filename) // 将文件保存到指定位置，这里假设为 images 文件夹下
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer dst.Close()

		// 将上传的文件内容拷贝到新文件中
		_, err = io.Copy(dst, src)
		if err != nil {
			fmt.Println("Error copying file:", err)
			continue
		}

		// 根据操作类型更新数据库中的字段值
		action := r.FormValue("action")
		switch action {
		case "add-room":
			// 更新添加房间的逻辑
			result := Conn.Model(&Houseinfo{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating room:", result.Error)
				continue
			}
			// 处理其他逻辑...
		case "update-room":
			result := Conn.Model(&Houseinfo{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating room:", result.Error)
				continue
			}
			// 更新更新房间的逻辑
			// 处理其他逻辑...
		case "add-activity":
			// 更新添加活动的逻辑
			result := Conn.Model(&Bulletin{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating activity:", result.Error)
				continue
			}
			// 处理其他逻辑...
		case "edit-activity":
			result := Conn.Model(&Bulletin{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating room:", result.Error)
				continue
			}
			// 更新编辑活动的逻辑
			// 处理其他逻辑...
		case "add-product":
			// 更新添加商品的逻辑
			result := Conn.Model(&Products{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating product:", result.Error)
				continue
			}
			// 处理其他逻辑...
		case "edit-product":
			result := Conn.Model(&Products{}).Update("image_url", gorm.Expr("CONCAT(image_url, ?, ',')", file.Filename))
			if result.Error != nil {
				fmt.Println("Error updating room:", result.Error)
				continue
			}
			// 更新编辑商品的逻辑
			// 处理其他逻辑...
		}
	}

	fmt.Fprintf(w, "Files uploaded successfully")
}

func AdminGetAllMessages() ([]Messages, error) {
	var messages []Messages
	if err := Conn.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func SearchMessagesByUsername(username string) ([]Messages, error) {
	var messages []Messages

	// 查询用户名匹配的留言数据
	if err := Conn.Where("username = ?", username).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

// AdminGetProfile 根据用户名搜索个人资料数据
func AdminGetProfile(name string) (Admin, error) {
	// 在这里添加根据用户名查询个人资料的逻辑
	// 假设你已经正确设置了数据库连接和初始化了GORM

	// 创建Profile结构体用于存储个人资料数据
	profile := Admin{}

	// 执行查询
	if err := Conn.Where("username = ?", name).First(&profile).Error; err != nil {
		// 处理查询过程中的错误，例如返回错误信息给逻辑层
		return Admin{}, err
	}

	// 返回个人资料数据给逻辑层
	return profile, nil
}

// UpdateUserProfile 更新用户个人信息
func UpdateUserProfile(username, password, email string) error {
	// 执行数据库操作，更新管理员个人信息
	err := Conn.Model(&Admin{}).Where("username = ?", username).Updates(map[string]interface{}{
		"password": password,
		"email":    email,
	}).Error

	return err
}

func AdminGetRooms() ([]House, error) {
	var ret []House
	if err := Conn.Table("House").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		fmt.Println("House:", ret)
	}
	return ret, nil
}

func Admin1GetActivities() ([]Bulletin, error) {
	var ret []Bulletin
	if err := Conn.Table("bulletin").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret, nil
}

func Admin1GetProducts() ([]Products, error) {
	var ret []Products
	if err := Conn.Table("Products").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret, nil
}

func Admin1GetUsers() ([]User, error) {
	var ret []User
	if err := Conn.Table("user").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret, nil
}

func Admin1GetOrders() ([]UserHouse, error) {
	var orders []UserHouse
	if err := Conn.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
func Admin1GetProductOrders() ([]UserProduct, error) {
	var orders []UserProduct
	if err := Conn.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func Admin1GetChatMessages() ([]Messages, error) {
	var messages []Messages
	if err := Conn.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func Admin1GetProfiles(c *gin.Context) {
	// 从Cookie中获取用户ID
	userIDCookie, err := c.Cookie("Id")
	if err != nil {
		// 处理无法获取用户ID的情况
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法获取用户ID"})
		return
	}

	// 将字符串类型的用户ID转换为整数类型
	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		// 处理无法将字符串转换为整数的情况
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户ID无效"})
		return
	}

	user, err := GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load profile"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Loadprofile(userID int) (User, error) {
	var user User
	if err := Conn.First(&user, userID).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// AddRoom 向数据库添加房间记录
func AddRoom(name string, num int, price float64) error {
	room := House{Name: name, Num: num, Price: price}
	if err := Conn.Create(&room).Error; err != nil {
		return err
	}
	return nil
}
