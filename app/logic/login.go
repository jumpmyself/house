package logic

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"house/app/model"
	"house/app/tools"
	"net/http"
)

func GetCaptcha(context *gin.Context) {

	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.ECode{
		Data: captcha,
	})
}

func GetCaptchaAdmin(context *gin.Context) {
	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.ECode{
		Data: captcha,
	})
}

func Order(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", nil)
}

func Messages(c *gin.Context) {
	c.HTML(http.StatusOK, "messages.tmpl", nil)
}

func GeRen(c *gin.Context) {
	c.HTML(http.StatusOK, "adminprofile.tmpl", nil)
}

func GetOrderRooms(c *gin.Context) {
	// 调用 model.AdminGetRooms() 获取房间数据
	data, err := model.AdminGetRooms()
	if err != nil {
		// 如果发生错误，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    10005,
			"message": err.Error(),
		})
		return
	}

	// 返回房间数据
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func GetOrderProducts(c *gin.Context) {
	// 调用 model.AdminGetRooms() 获取房间数据
	data, err := model.Admin1GetProductOrders()
	if err != nil {
		// 如果发生错误，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    10005,
			"message": err.Error(),
		})
		return
	}

	// 返回商品订单数据
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func MarkAsDelivered(c *gin.Context) {
	orderId := c.Param("orderId")

	// 调用 model 层的函数更新订单状态为已发货
	if err := model.UpdateOrderStatus(orderId, "已发货"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法更新订单状态",
		})
		return
	}

	// 发货成功的响应
	c.JSON(http.StatusOK, gin.H{
		"message": "订单状态已更新为已发货",
	})
}

func GetMessages(c *gin.Context) {
	messages, err := model.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func ReplyMessage(c *gin.Context) {
	// 获取 messageId 参数
	messageId := c.PostForm("messageId")

	// 从请求体中读取回复文本
	replyText := c.PostForm("replyText")

	// 根据 messageId 查找消息记录
	var message model.Messages
	err := model.Conn.Where("id = ?", messageId).First(&message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Message not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to find message",
			})
		}
		return
	}

	// 更新消息记录的 Message 字段
	message.Message = replyText
	err = model.Conn.Save(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save reply",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reply saved successfully",
	})
}
