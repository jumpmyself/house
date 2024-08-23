package model

import (
	"fmt"
	"github.com/go-pay/util/snowflake"
	"strconv"
	"time"
)

func AdminAddProduct(product *Products) error {
	// 连接数据库并添加产品数据
	if err := Conn.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func AdminDeleteProduct(productID string) error {
	// 连接数据库并删除产品数据
	if err := Conn.Where("name = ?", productID).Delete(&Products{}).Error; err != nil {
		return err
	}
	return nil
}

func AdminUpdateProduct(updatedProduct *Products) error {
	// 连接数据库并更新产品数据
	if err := Conn.Save(updatedProduct).Error; err != nil {
		return err
	}
	return nil
}

// GetProductsByPage 根据页码获取产品列表
func GetProductsByPage(page, limit int) ([]Products, int64, error) {
	var products []Products
	var totalProductsCount int64
	offset := (page - 1) * limit
	if err := Conn.Model(&Products{}).Count(&totalProductsCount).Error; err != nil {
		return nil, 0, err
	}
	if err := Conn.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, totalProductsCount, nil
}

// SearchProduct 根据关键字搜索商品
func SearchProduct(keyword string) ([]Products, error) {
	var products []Products
	// 在这里执行数据库查询，这里使用的是 GORM 作为 ORM 库
	result := Conn.Where("name LIKE ?", "%"+keyword+"%").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProductByID(ProductId int) (*Products, error) {
	var product Products
	fmt.Printf("ProductId: %d\n", ProductId)
	if err := Conn.Table("products").Where("ID = ?", ProductId).First(&product).Error; err != nil {
		return nil, err
	}
	fmt.Printf("romm:%+v\n", product)
	return &product, nil
}

func BuyProduct(name, UserName, ProductId, ProductName, Phone, Num, PriceStr, TotalPriceStr, CreatTime string) error {
	num, _ := strconv.Atoi(Num)
	totalPrice, _ := strconv.ParseFloat(TotalPriceStr, 64)
	productId, _ := strconv.Atoi(ProductId)

	// 创建一个新的雪花节点
	node, err := snowflake.NewNode(1) // 这里的参数是节点的 ID
	if err != nil {
		fmt.Println("uuid创建失败:", err)
		return nil
	}

	// 生成一个唯一的 ID
	uuid := node.Generate().Int64()

	// 移除单位部分
	//priceStr := strings.Replace(PriceStr, "元", "", 1)

	// 将价格部分转换为浮点数
	price, err := strconv.ParseFloat(PriceStr, 64)
	if err != nil {
		fmt.Println("解析价格失败:", err)
		return nil
	}

	fmt.Println("价格:", price)

	order := UserProduct{
		Uuid:        strconv.FormatInt(uuid, 10),
		Username:    UserName,
		ProductId:   int64(productId),
		Name:        name,
		ProductName: ProductName,
		Telephone:   Phone,
		Quantity:    int64(num),
		Price:       price,
		TotalPrice:  totalPrice,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 插入数据到数据库
	if err := Conn.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOrderStatus 根据订单 UUID 更新订单状态
func UpdateOrderStatus(uuid, newStatus string) error {
	// 获取全局的数据库连接
	db := Conn

	// 查找对应 UUID 的订单
	var order UserProduct
	if err := db.Where("uuid = ?", uuid).First(&order).Error; err != nil {
		return err
	}

	// 更新订单状态
	order.Status = newStatus

	// 保存更新后的订单状态
	if err := db.Save(&order).Error; err != nil {
		return err
	}

	return nil
}
