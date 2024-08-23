package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"house/app/model"
	"net/http"
	"path/filepath"
)

// AdminAddActivity 用于处理管理员添加活动的请求
func AdminAddActivity(c *gin.Context) {
	var newActivity model.Bulletin

	newActivity.Title = c.PostForm("title")
	newActivity.Content = c.PostForm("content")
	newActivity.DateTime = c.PostForm("time")
	newActivity.Place = c.PostForm("location")

	// 检查活动信息是否完整
	if newActivity.Title == "" || newActivity.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "活动标题和描述不能为空"})
		return
	}

	// 检查是否上传了图片
	file, err := c.FormFile("image_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传活动图片：" + err.Error()})
		return
	}

	// 生成唯一的文件名
	filename := uuid.New().String() + filepath.Ext(file.Filename)

	// 将文件保存到指定目录
	if err := c.SaveUploadedFile(file, "app/images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存活动图片失败：" + err.Error()})
		return
	}

	// 将文件名存储到活动信息中
	newActivity.ImageUrl = filename

	// 调用模型层方法插入活动数据
	if err := model.AdminAddActivity(&newActivity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "插入活动数据失败：" + err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "活动添加成功", "activity": newActivity})
}
func AdminDeleteActivity(c *gin.Context) {
	// 从请求中获取活动名称
	activityID := c.PostForm("Title")

	// 调用模型层方法删除活动数据
	if err := model.AdminDeleteActivity(activityID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}

func AdminUpdateActivity(c *gin.Context) {
	// 从请求中解析活动信息
	var activity model.Bulletin
	if err := c.Bind(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := c.FormFile("Images")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传房间图片"})
		return
	}

	// 生成唯一的文件名
	filename := uuid.New().String() + filepath.Ext(file.Filename)

	// 将文件保存到指定目录
	if err := c.SaveUploadedFile(file, "app/images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 将文件名存储到房间信息中
	activity.ImageUrl = filename
	// 调用模型层方法更新活动数据
	if err := model.AdminUpdateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}
