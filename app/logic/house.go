package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"house/app/model"
	"net/http"
	"path/filepath"
	"strconv"
)

func GetRooms(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.DefaultQuery("pageSize", "6")
	fmt.Println(pageStr, pageSizeStr)
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return
	}

	ret, err := model.GetRooms(page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": ret})
}

func RoonInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "roominfo.tmpl", nil)
}

func GetRoomDetail(c *gin.Context) {
	// 获取 URL 参数 roomId
	roomIdStr := c.Query("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	// 查询房间详情
	room, err := model.GetRoomByID(roomId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

func GetRoomBySearch(c *gin.Context) {
	roomType := c.Query("roomType")
	fmt.Printf("roomType:%s\n", roomType)
	ret := model.GetRoomBySearch(roomType)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": ret})
}

func GetShopping(c *gin.Context) {
	c.HTML(http.StatusOK, "shopping.tmpl", nil)
}

func GetChat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.tmpl", nil)
}

func AdminAddRoom(c *gin.Context) {
	// 解析房间信息
	var room model.House
	room.Name = c.PostForm("name")                              // 获取名称字段
	room.Num, _ = strconv.Atoi(c.PostForm("num"))               // 获取并转换数量字段
	room.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64) // 获取并转换价格字段

	// 检查房间信息是否完整
	if room.Name == "" || room.Price == 0 || room.Num == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "房间信息不完整"})
		return
	}

	// 从请求中获取图片文件
	file, err := c.FormFile("image_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传房间图片"})
		return
	}

	// 生成唯一的文件名
	filename := uuid.New().String() + filepath.Ext(file.Filename)

	// 保存文件到指定目录
	if err := c.SaveUploadedFile(file, "app/images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 将文件名存储到房间信息中
	room.ImageUrl = filename

	// 调用 model 层方法插入房间数据
	if err := model.AdminAddRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "房间添加成功"})
}

func AdminDeleteRoom(c *gin.Context) {
	// 从请求参数中获取房间ID
	Name := c.PostForm("Name")
	if Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未提供房间名"})
		return
	}

	// 调用 model 层方法删除房间数据
	if err := model.AdminDeleteRoom(Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "房间删除成功"})
}

func AdminUpdateRoom(c *gin.Context) {
	// 从请求中解析房间信息
	var room model.House
	if err := c.ShouldBind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("room:%v\n", room)
	file, err := c.FormFile("Images")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传房间图片"})
		return
	}

	// 检查房间信息是否完整
	if room.Name == "" || room.Description == "" || room.Price == 0 || room.Num == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "房间信息不完整"})
		fmt.Printf("room:%v\n", room)
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
	room.ImageUrl = filename

	// 调用 model 层方法插入房间数据
	if err := model.AdminUpdateRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "房间添加成功"})
}
