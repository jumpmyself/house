package logic

import (
	_ "database/sql"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"house/app/model"
	"house/app/tools"
	"house/config"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func UserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "ulogin.tmpl", nil)
}

func DoUserLogin(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusOK, tools.ECode{

			Message: "err.Error()", //有风险

		})
		return
	}

	fmt.Printf("user:%+v\n", user)
	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: user.CaptchaId,
		Data:      user.CaptchaValue,
	}) {
		ctx.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "验证码校验失败",
		})
		return
	}

	ret := model.GetUser(user.Name)
	fmt.Printf("ret:%+v\n", ret)
	if ret.ID < 1 || ret.Password != user.Password {
		ctx.JSON(http.StatusOK, tools.UserErr)
		return
	}

	ctx.SetCookie("name", user.Name, 3600, "/", "", true, false)
	ctx.SetCookie("Id", fmt.Sprint(ret.ID), 3600, "/", "", true, false)

	_ = model.SetSession(ctx, user.Name, ret.ID)

	ctx.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
		Data:    ret,
	})
	return
}

func DoUserRegister(c *gin.Context) {
	// 从请求中获取用户提交的数据
	name := c.PostForm("name")
	password := c.PostForm("password")
	captchaValue := c.PostForm("captcha_value")

	// 调用逻辑层的注册函数处理注册逻辑
	err := model.UserRegister(name, password, captchaValue)
	if err != nil {
		// 发生错误，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	// 注册成功，返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
	})
}

func UserIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "uindex.tmpl", nil)
}

func UserAnnounce(c *gin.Context) {
	c.HTML(http.StatusOK, "uannounce.tmpl", nil)
}

// GetUserAnnounce 根据日期获取活动公告信息
func GetUserAnnounce(c *gin.Context) {
	dateStr := c.Query("date") // 从查询参数中获取日期字符串

	fmt.Printf("dateStr:%s\n", dateStr)
	// 获取该日期当天的活动公告信息
	var bulletins []model.Bulletin
	if err := model.Conn.Where("date_time >= ? AND date_time < DATE_ADD(?, INTERVAL 1 DAY)", dateStr, dateStr).Find(&bulletins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bulletins)
}

func Loadprofile(c *gin.Context) {
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

	// 从数据库中查询用户信息
	var user model.User
	if err := model.Conn.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法加载用户信息"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.tmpl", nil)
}

func OrderHouse(c *gin.Context) {

	// 绑定请求数据到结构体
	var data model.UserHouse
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	// 生成雪花算法的唯一 ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成唯一 ID 失败"})
		return
	}
	uuid := node.Generate().Int64()

	// 将生成的 uid 设置到数据结构中
	data.UUID = uuid
	// 获取当前时间
	currentTime := time.Now()

	// 将当前时间设置到数据结构中
	data.CreatedTime = currentTime
	data.UpdatedTime = currentTime

	fmt.Printf("data:%+v\n\n", data)
	// 调用模型层创建用户订房订单信息
	err = model.CreateUserHouse(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加用户订房订单信息失败"})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "订房成功",
		Data:    nil,
	})
}

func OrderProduct(c *gin.Context) {
	// 绑定请求数据到结构体
	var data model.UserProduct
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	// 生成雪花算法的唯一 ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成唯一 ID 失败"})
		return
	}
	uuid := node.Generate().Int64()

	// 将生成的 UUID 设置到数据结构中
	data.Uuid = strconv.FormatInt(uuid, 10)

	// 获取当前时间
	currentTime := time.Now()

	// 将当前时间设置到数据结构中
	data.CreatedAt = currentTime
	data.UpdatedAt = currentTime

	// 打印订单数据，方便调试
	fmt.Printf("data:%+v\n", data)

	// 调用模型层创建用户订购商品订单信息
	err = model.CreateUserProduct(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加用户订购商品订单信息失败"})
		return
	}

	c.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "订购商品成功",
		Data:    nil,
	})
}

func UserLogout(c *gin.Context) {
	// 在此处执行注销操作，例如清除会话或删除令牌

	// 使用 c.Redirect() 将用户重定向到登录页面
	c.Redirect(http.StatusFound, "/user/login") // 假设登录页面的路由为 "/login"
}

// UpdateProfile 更新用户个人资料的处理程序
func UpdateProfile(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node, err := snowflake.NewNode(1) // 这里的参数是节点的 ID
	if err != nil {
		fmt.Println("uid创建失败:", err)
		return
	}

	// 生成一个唯一的 ID
	uid := node.Generate().Int64()

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
	// 更新用户信息
	result := model.Conn.Model(&model.User{}).Where("id = ?", userID).Updates(model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Uid:       uid,
		City:      user.City,
		Age:       user.Age,
		Telephone: user.Telephone,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 检查是否有更新
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户没有更新数据"})
		return
	}

	c.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "更新资料成功，请牢记您的密码",
		Data:    nil,
	})
}

func AdminAddUser(c *gin.Context) {
	// 从请求中解析用户信息
	var newUser model.User
	newUser.Username = c.PostForm("username")
	newUser.Password = c.PostForm("password")
	newUser.Email = c.PostForm("email")
	newUser.Telephone = c.PostForm("phone")
	newUser.City = c.PostForm("address")

	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" || newUser.Telephone == "" || newUser.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不能为空"})
		return
	}

	file, err := c.FormFile("image_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传头像"})
		return
	}
	// 生成唯一的文件名
	filename := uuid.New().String() + filepath.Ext(file.Filename)

	// 保存文件到指定目录
	if err := c.SaveUploadedFile(file, "app/images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	newUser.ImageUrl = filename
	// 调用模型层方法插入用户数据
	if err := model.AdminAddUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "新建用户成功"})
}

func AdminDeleteUser(c *gin.Context) {
	// 从请求中获取要删除的用户ID
	userID := c.PostForm("Username")

	// 调用模型层方法删除用户
	if err := model.AdminDeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func AdminUpdateUser(c *gin.Context) {
	// 从请求中解析更新的用户信息
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user:", user)

	// 调用模型层方法更新用户信息
	if err := model.AdminUpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func GetBookPrice(c *gin.Context) {
	roomId := c.Query("roomId")
	if roomId == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "未找到图书id"})
		return
	}
	name, price, err := model.GetRoomPrice(roomId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "message": "获取图书信息失败：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": price, "name": name})
}

func BuyRoom(c *gin.Context) {
	// 获取表单数据
	UserName := c.PostForm("username")
	HouseId := c.PostForm("houseId")
	HouseName := c.PostForm("room")
	Phone := c.PostForm("phone")
	Notes := c.PostForm("notes")
	Num := c.PostForm("quantity")
	PriceStr := c.PostForm("unit-price")
	TotalPriceStr := c.PostForm("total-price")
	name := c.PostForm("name")
	fmt.Println("name:", name)
	CreatTime := time.Now().Format("2006-01-02 15:04:05")

	// 获取url进行支付
	client, err := alipay.NewClient(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		log.Println("支付宝初始化错误")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "支付宝初始化错误"})
		return
	}
	client.SetCharset("utf-8").SetSignType(alipay.RSA2).SetNotifyUrl(config.NotifyURL).SetReturnUrl(config.ReturnURL)

	ts := time.Now().UnixMilli() //生成唯一订单号
	outTradeNo := fmt.Sprintf("%d", ts)

	bm := make(gopay.BodyMap)
	bm.Set("subject", "图书商城支付页面")
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("total_amount", TotalPriceStr)
	bm.Set("product_code", config.ProductCode)
	//bm.Set("timeout_express", "30m")
	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "支付链接生成失败"})
		return
	}

	fmt.Println("houseid:", HouseId)
	fmt.Println("PriceStr:", PriceStr)
	// 调用模型的函数来处理订单信息
	err = model.BuyRoom(name, UserName, HouseId, HouseName, Phone, Notes, Num, PriceStr, TotalPriceStr, CreatTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	// 返回购买成功的 JSON 响应，并返回支付链接
	c.JSON(http.StatusOK, gin.H{"message": "恭喜您，买书成功", "payUrl": payUrl})

}

func GetUserOrder(c *gin.Context) {
	// 声明一个结构体来接收请求体中的数据
	var requestData struct {
		Name string `json:"name"`
	}

	// 解析请求体中的JSON数据到结构体中
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从requestData中获取用户名
	name := requestData.Name
	fmt.Println("name:", name)

	// 调用model层的函数直接根据用户名获取订单信息
	orders, err := model.GetUserOrdersByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回订单信息给客户端
	c.JSON(http.StatusOK, orders)
}

func GetUserProductOrder(c *gin.Context) {
	// 声明一个结构体来接收请求体中的数据
	var requestData struct {
		Name string `json:"name"`
	}

	// 解析请求体中的JSON数据到结构体中
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从requestData中获取用户名
	name := requestData.Name
	fmt.Println("name:", name)

	// 调用model层的函数直接根据用户名获取订单信息
	orders, err := model.GetUserProductOrdersByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回订单信息给客户端
	c.JSON(http.StatusOK, orders)
}

func GetMessage(c *gin.Context) {
	var messages []model.Messages

	// 从数据库中获取留言数据
	result := model.Conn.Find(&messages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func SendMessage(c *gin.Context) {
	// 解析请求中的参数
	var message model.Messages
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打印请求的 JSON 数据
	jsonData, _ := json.Marshal(message)
	fmt.Println(string(jsonData))

	// 将留言保存到数据库
	result := model.Conn.Create(&message)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func GetRoom(c *gin.Context) {
	id := c.Query("id") // 从请求中获取传入的房间 ID

	var room model.House
	if err := model.Conn.First(&room, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch room data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": room,
	})
}

func GetActivity(c *gin.Context) {
	// 从请求中获取传入的活动 ID
	id := c.Query("id")

	var activity model.Bulletin // 假设您有名为 Activity 的模型结构体

	// 使用 GORM 查询活动信息
	if err := model.Conn.First(&activity, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch activity data",
		})
		return
	}

	// 返回查询到的活动信息给前端
	c.JSON(200, gin.H{
		"data": activity,
	})
}

func GetProduct(c *gin.Context) {
	// 从请求中获取传入的产品 ID
	id := c.Query("id")

	var product model.Products // 假设您有名为 Product 的模型结构体

	// 使用 GORM 查询产品信息
	if err := model.Conn.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch product data",
		})
		return
	}

	// 返回查询到的产品信息给前端
	c.JSON(200, gin.H{
		"data": product,
	})
}

func GetUser(c *gin.Context) {
	// 从请求中获取传入的用户 ID
	id := c.Query("id")

	var user model.User // 假设您有名为 User 的模型结构体

	// 使用 GORM 查询用户信息
	if err := model.Conn.First(&user, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch user data",
		})
		return
	}

	// 返回查询到的用户信息给前端
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetProfile(c *gin.Context) {
	// 从请求中获取传入的个人资料 ID
	id := c.Query("id")

	var profile model.User // 假设您有名为 Profile 的模型结构体

	// 使用 GORM 查询个人资料信息
	if err := model.Conn.First(&profile, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch profile data",
		})
		return
	}

	// 返回查询到的个人资料信息给前端
	c.JSON(200, gin.H{
		"data": profile,
	})
}

func UpdateRoom(c *gin.Context) {
	// 从请求中获取房间ID
	id := c.PostForm("id")
	fmt.Println("id:", id)

	// 构建房间对象
	var room model.House
	if err := model.Conn.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "房间未找到"})
		return
	}

	// 从请求表单中获取要更新的数据
	room.Name = c.PostForm("name")                              // 房间名称
	room.Num, _ = strconv.Atoi(c.PostForm("num"))               // 库存数量
	room.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64) // 价格
	// 如果需要更新其他字段，也从请求表单中获取并更新相应的字段
	// 从请求中获取图片文件
	file, err := c.FormFile("image_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传房间图片"})
		return
	}

	// 生成唯一的文件名
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	// 保存更新后的房间信息
	// 保存文件到指定目录
	if err := c.SaveUploadedFile(file, "app/images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}
	// 将文件名存储到房间信息中
	room.ImageUrl = filename
	if err := model.Conn.Save(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新房间信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "房间信息更新成功", "room": room})
}
func UpdateActivity(c *gin.Context) {
	// 从请求中获取活动ID
	id := c.Param("id")

	// 构建活动对象
	var activity model.Bulletin
	if err := model.Conn.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "活动未找到"})
		return
	}

	activity.Title = c.PostForm("title")
	activity.Content = c.PostForm("content")
	activity.Place = c.PostForm("location")
	activity.DateTime = c.PostForm("time")

	if activity.Title == "" || activity.Content == "" || activity.Place == "" || activity.DateTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不能为空"})
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

	activity.ImageUrl = filename
	fmt.Println("activity:", activity)

	if err := model.Conn.Save(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新活动信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "活动信息更新成功", "activity": activity})
}
func UpdateProduct(c *gin.Context) {
	// 从请求中获取产品ID
	id := c.Param("id")

	// 构建产品对象
	var product model.Products
	if err := model.Conn.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "产品未找到"})
		return
	}

	product.Name = c.PostForm("name")
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.Num, _ = strconv.Atoi(c.PostForm("num"))

	if product.Name == "" || product.Price == 0 || product.Num == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不能为空"})

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

	product.ImageUrl = filename

	if err := model.Conn.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新产品信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "产品信息更新成功", "product": product})
}
func UpdateUser(c *gin.Context) {
	// 从请求中获取用户ID
	id := c.Param("id")

	// 构建用户对象
	var user model.User
	if err := model.Conn.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Telephone = c.PostForm("phone")
	user.City = c.PostForm("address")

	if user.Username == "" || user.Password == "" || user.Email == "" || user.Telephone == "" || user.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不能为空"})
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

	user.ImageUrl = filename

	if err := model.Conn.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功", "user": user})
}
