package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"house/app/model"
	"house/config"
	"log"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func AdminAddProduct(c *gin.Context) {
	var product model.Products
	product.Name = c.PostForm("name")
	product.Num, _ = strconv.Atoi(c.PostForm("num"))
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	if product.Name == "" || product.Num == 0 || product.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请填写完整的产品信息"})
		return
	} // 从请求中解析产品信息

	file, err := c.FormFile("image_url")
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
	product.ImageUrl = filename

	// 调用模型层方法添加产品数据
	if err := model.AdminAddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

func AdminDeleteProduct(c *gin.Context) {
	// 从请求参数中获取要删除的产品ID
	productID := c.PostForm("Name")
	fmt.Println("productID:", productID)

	// 调用模型层方法删除产品数据
	if err := model.AdminDeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func AdminUpdateProduct(c *gin.Context) {
	// 从请求中解析更新后的产品信息
	var updatedProduct model.Products
	if err := c.Bind(&updatedProduct); err != nil {
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
	updatedProduct.ImageUrl = filename

	// 调用模型层方法更新产品数据
	if err := model.AdminUpdateProduct(&updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// GetShoppingList 获取产品列表
func GetShoppingList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // 从查询参数中获取页码，默认为第一页
	if err != nil || page < 1 {
		page = 1 // 如果页码不合法，则默认为第一页
	}

	limit := 6 // 每页显示的数量

	products, totalProductsCount, err := model.GetProductsByPage(page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(totalProductsCount) / float64(limit))) // 计算总页数

	c.JSON(200, gin.H{
		"status":      "success",
		"message":     "Products retrieved successfully",
		"data":        products,
		"currentPage": page,
		"totalPages":  totalPages,
	})
}

func SearchProduct(c *gin.Context) {
	// 从请求参数中获取关键字
	keyword := c.Request.FormValue("keyword")
	fmt.Println("keyword:", keyword)
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少关键字"})
		return
	}

	// 调用 model 层函数执行数据库查询
	products, err := model.SearchProduct(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		return
	}

	// 返回搜索结果
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "productinfo.tmpl", nil)
}

func GetProductInfoDetail(c *gin.Context) {
	// 获取 URL 参数 roomId
	ProductIdStr := c.Query("productId")
	fmt.Printf("rommIdStr:%s\n", ProductIdStr)
	ProductId, err := strconv.Atoi(ProductIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	// 查询房间详情
	product, err := model.GetProductByID(ProductId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func BuyProduct(c *gin.Context) {
	// 获取表单数据
	UserName := c.PostForm("username")
	ProductId := c.PostForm("productId")
	ProductName := c.PostForm("product")
	Phone := c.PostForm("phone")
	Num := c.PostForm("quantity")
	PriceStr := c.PostForm("unit-price")
	TotalPriceStr := c.PostForm("total-price")
	name := c.PostForm("name")
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
	bm.Set("timeout_express", "5m")
	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "支付链接生成失败"})
		return
	}

	fmt.Println("ProductId:", ProductId)
	fmt.Println("PriceStr:", PriceStr)
	// 调用模型的函数来处理订单信息
	err = model.BuyProduct(name, UserName, ProductId, ProductName, Phone, Num, PriceStr, TotalPriceStr, CreatTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	// 返回购买成功的 JSON 响应，并返回支付链接
	c.JSON(http.StatusOK, gin.H{"message": "恭喜您，买书成功", "payUrl": payUrl})

}
