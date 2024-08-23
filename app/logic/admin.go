package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"house/app/model"
	"house/app/tools"
	"net/http"
	"strconv"
)

type Admin struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func DoAdminLogin(ctx *gin.Context) {
	var admin Admin
	if err := ctx.ShouldBind(&admin); err != nil {
		ctx.JSON(http.StatusOK, tools.ECode{
			Message: "err.Error()", //有风险
		})
		return
	}

	fmt.Printf("admin:%+v\n", admin)
	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: admin.CaptchaId,
		Data:      admin.CaptchaValue,
	}) {
		ctx.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "验证码校验失败",
		})
		return
	}

	ret := model.GetAdminByName(admin.Name)
	fmt.Printf("ret:%+v\n", ret)
	if ret.ID < 1 || ret.Password != tools.EncryptV1(admin.Password) {
		ctx.JSON(http.StatusOK, tools.UserErr)
		return
	}

	ctx.SetCookie("name", admin.Name, 3600, "/", "", true, false)
	ctx.SetCookie("Id", fmt.Sprint(ret.ID), 3600, "/", "", true, false)

	_ = model.SetSession(ctx, admin.Name, ret.ID)

	ctx.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
		Data:    ret,
	})
	return
}

func AdminIndex(c *gin.Context) {
	c.HTML(200, "admin.tmpl", nil)
}

func AddAnnounce(c *gin.Context) {
	var b model.Bulletin
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行插入操作
	result := model.Conn.Create(&b)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 获取插入行的 ID
	id := b.ID

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func AdminAnnounce(c *gin.Context) {
	c.HTML(http.StatusOK, "aannounce.tmpl", nil)
}

func AdminGetRooms(c *gin.Context) {
	ret, _ := model.AdminGetRooms()
	c.JSON(http.StatusOK, ret)
}
func AdminGetRooms1(c *gin.Context) {
	ret := model.AdminGetRooms1()
	c.JSON(http.StatusOK, ret)
}

func AdminGetActivities(c *gin.Context) {
	ret := model.AdminGetActivities()
	c.JSON(http.StatusOK, ret)
}

func AdminGetProducts(c *gin.Context) {
	ret := model.AdminGetProducts()
	c.JSON(http.StatusOK, ret)
}
func GetUsers(c *gin.Context) {
	ret := model.AdminGetUsers()
	c.JSON(http.StatusOK, ret)
}

func GetHouseOrder(c *gin.Context) {
	ret := model.AdminGetHouseOrder()
	c.JSON(http.StatusOK, ret)
}

func GetProductOrder(c *gin.Context) {
	ret := model.AdminGetProductOrder()
	c.JSON(http.StatusOK, ret)
}

func AdminGetAllMessages(c *gin.Context) {
	messages, err := model.AdminGetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func AdminSearchMessage(c *gin.Context) {
	// 获取前端传递的用户名参数
	username := c.Query("username")
	fmt.Println(username)
	// 在这里可以添加逻辑验证用户名参数的有效性，例如检查是否为空或是否符合要求

	// 调用服务层或数据访问层的函数，根据用户名搜索留言数据
	messages, err := model.SearchMessagesByUsername(username)
	if err != nil {
		// 处理搜索过程中的错误，例如返回错误信息给前端
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索留言时发生错误"})
		return
	}

	// 返回搜索结果给前端
	c.JSON(http.StatusOK, messages)

}

func AdminProfile(c *gin.Context) {
	// 获取前端传递的用户名参数
	name := c.Query("name")
	if name == "" {
		name = "admin"
	}
	fmt.Println(name)
	// 在这里可以添加逻辑验证用户名参数的有效性，例如检查是否为空或是否符合要求
	// 调用服务层或数据访问层的函数，根据用户名搜索留言数据
	messages, err := model.AdminGetProfile(name)
	if err != nil {
		// 处理搜索过程中的错误，例如返回错误信息给前端
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索管理员信息发生错误"})
		return
	}

	// 返回搜索结果给前端
	c.JSON(http.StatusOK, messages)
}

func AdminUpdateProfile(c *gin.Context) {
	// 从请求中获取表单数据
	username := c.PostForm("Username")
	password := c.PostForm("Password")
	email := c.PostForm("Email")

	// 执行个人信息更新的逻辑
	err := model.UpdateUserProfile(username, password, email)
	if err != nil {
		// 处理更新失败的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "更新个人信息时出错",
			"message": err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "个人信息已成功更新",
	})
}

func AdminLogout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/user/login") // 假设登录页面的路由为 "/login"
}

func GetData(c *gin.Context) {
	target := c.Query("target")
	var data interface{}
	var err error

	switch target {
	case "room":
		data, err = model.AdminGetRooms()
	case "activity":
		data, err = model.Admin1GetActivities()
	case "product":
		data, err = model.Admin1GetProducts()
	case "user":
		data, err = model.Admin1GetUsers()
	case "order":
		data, err = model.Admin1GetOrders()
	case "chatroom":
		data, err = model.Admin1GetChatMessages()
	case "profile":
		userIDStr := c.Query("userID")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
			return
		}
		data, err = model.Loadprofile(userID)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func SearchData(c *gin.Context) {
	// 获取请求参数
	target := c.Query("target")   // 目标类型
	keyword := c.Query("keyword") // 搜索关键词

	// 根据不同的目标类型执行不同的数据库查询操作
	switch target {
	case "room":
		var rooms []model.House
		model.Conn.Where("name LIKE ?", "%"+keyword+"%").Find(&rooms)
		c.JSON(http.StatusOK, rooms)
	case "activity":
		var activities []model.Bulletin
		model.Conn.Where("title LIKE ?", "%"+keyword+"%").Find(&activities)
		c.JSON(http.StatusOK, activities)
	case "product":
		var products []model.Products
		model.Conn.Where("name LIKE ?", "%"+keyword+"%").Find(&products)
		c.JSON(http.StatusOK, products)
	case "user":
		var user []model.User
		model.Conn.Where("username LIKE ?", "%"+keyword+"%").Find(&user)
		c.JSON(http.StatusOK, user)
	case "order":
		var orders []model.UserHouse
		model.Conn.Where("user_name LIKE ? OR house_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&orders)
		c.JSON(http.StatusOK, orders)
	case "chatroom":
		var messages []model.Messages
		model.Conn.Where("username LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&messages)
		c.JSON(http.StatusOK, messages)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
	}
}

func DeleteItem(c *gin.Context) {
	// 获取请求参数
	id := c.PostForm("id")         // 待删除项的ID
	target := c.PostForm("target") // 目标类型

	// 根据不同的目标类型执行不同的数据库删除操作
	switch target {
	case "room":
		// 执行删除房间的操作
		if err := model.Conn.Where("id = ?", id).Delete(&model.House{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	case "activity":
		// 执行删除活动的操作
		if err := model.Conn.Where("id = ?", id).Delete(&model.Bulletin{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	case "product":
		// 执行删除商品的操作
		if err := model.Conn.Where("id = ?", id).Delete(&model.Products{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	case "user":
		// 执行删除用户的操作
		if err := model.Conn.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	case "order":
		// 执行删除订单的操作
		if err := model.Conn.Where("house_id = ?", id).Delete(&model.UserHouse{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	case "chatroom":
		// 执行删除留言的操作
		if err := model.Conn.Where("id = ?", id).Delete(&Message{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
		return
	}

	// 删除成功，返回成功消息
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetItemData(c *gin.Context) {
	// 获取请求参数
	target := c.Query("target") // 目标类型
	itemID := c.Query("id")     // 数据项的ID

	// 根据不同的目标类型执行不同的数据库查询操作
	switch target {
	case "room":
		var room model.House
		if err := model.Conn.First(&room, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
			return
		}
		c.JSON(http.StatusOK, room)
	case "activity":
		var activity model.Bulletin
		if err := model.Conn.First(&activity, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		c.JSON(http.StatusOK, activity)
	case "product":
		var product model.Products
		if err := model.Conn.First(&product, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	case "user":
		var user model.User
		if err := model.Conn.First(&user, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	case "order":
		var order model.UserHouse
		if err := model.Conn.First(&order, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	case "chatroom":
		var message model.Messages
		if err := model.Conn.First(&message, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
			return
		}
		c.JSON(http.StatusOK, message)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
	}
}

func EditData(c *gin.Context) {
	// 获取请求参数
	target := c.PostForm("target") // 目标类型
	itemID := c.PostForm("id")     // 数据项的ID

	// 解析请求体中的数据
	var requestBody interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 根据不同的目标类型执行不同的数据库更新操作
	switch target {
	case "room":
		var room model.House
		if err := model.Conn.First(&room, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
			return
		}
		// 更新房间数据
		if err := model.Conn.Model(&room).Updates(requestBody).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully"})
	case "activity":
		var activity model.Bulletin
		if err := model.Conn.First(&activity, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		// 更新活动数据
		if err := model.Conn.Model(&activity).Updates(requestBody).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
	case "product":
		var product model.Products
		if err := model.Conn.First(&product, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		// 更新产品数据
		if err := model.Conn.Model(&product).Updates(requestBody).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
	case "user":
		var user model.User
		if err := model.Conn.First(&user, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// 更新用户数据
		if err := model.Conn.Model(&user).Updates(requestBody).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
	}
}
func AddData(c *gin.Context) {
	// 获取请求参数
	target := c.PostForm("target") // 目标类型

	// 根据不同的目标类型执行不同的数据库插入操作
	switch target {
	case "room":
		var room model.House
		if err := c.Bind(&room); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&room).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add room"})
			return
		}
		c.JSON(http.StatusOK, room)
	case "activity":
		var activity model.Bulletin
		if err := c.Bind(&activity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&activity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add activity"})
			return
		}
		c.JSON(http.StatusOK, activity)
	case "product":
		var product model.Products
		if err := c.Bind(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
			return
		}
		c.JSON(http.StatusOK, product)
	case "user":
		var user User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
			return
		}
		c.JSON(http.StatusOK, user)
	case "order":
		var order model.UserHouse
		if err := c.Bind(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order"})
			return
		}
		c.JSON(http.StatusOK, order)
	case "chatroom":
		var message Message
		if err := c.Bind(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := model.Conn.Create(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add message"})
			return
		}
		c.JSON(http.StatusOK, message)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
	}
}

//func AddData(c *gin.Context) {
//	target := c.Query("target")
//	var err error
//
//	switch target {
//	case "room":
//		var room model.Room
//		if err := c.BindJSON(&room); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room data"})
//			return
//		}
//		err = model.AddRoom(room)
//	case "activity":
//		var activity model.Activity
//		if err := c.BindJSON(&activity); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity data"})
//			return
//		}
//		err = model.AddActivity(activity)
//	// 添加其他目标的新增逻辑
//	default:
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
//		return
//	}
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add data"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully"})
//}
//
//func EditData(c *gin.Context) {
//	target := c.Query("target")
//	var err error
//
//	switch target {
//	case "room":
//		var room model.Room
//		if err := c.BindJSON(&room); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room data"})
//			return
//		}
//		err = model.EditRoom(room)
//	case "activity":
//		var activity model.Activity
//		if err := c.BindJSON(&activity); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity data"})
//			return
//		}
//		err = model.EditActivity(activity)
//	// 添加其他目标的编辑逻辑
//	default:
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
//		return
//	}
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit data"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Data edited successfully"})
//}
//}
//
//func DeleteData(c *gin.Context) {
//	target := c.Query("target")
//	id := c.Param("id") // 假设从路径中获取要删除数据的 ID
//
//	var err error
//
//	switch target {
//	case "room":
//		err = model.DeleteRoom(id)
//	case "activity":
//		err = model.DeleteActivity(id)
//	// 添加其他目标的删除逻辑
//	default:
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
//		return
//	}
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
//}

func AddRoom(c *gin.Context) {
	var ret model.House
	if err := c.ShouldBindJSON(&ret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.AddRoom(ret.Name, int(ret.Num), ret.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room added successfully"})
}
