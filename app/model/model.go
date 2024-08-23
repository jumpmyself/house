package model

import "time"

// Admin undefined
type Admin struct {
	ID       int64  `json:"id" gorm:"id"`
	Username string `json:"Username" gorm:"Username"`
	Email    string `json:"Email" gorm:"Email"`
	Password string `json:"Password" gorm:"Password"`
}

// TableName 表名称
func (*Admin) TableName() string {
	return "admin"
}

// Houseinfo undefined
type Houseinfo struct {
	ID          int64   `json:"id" gorm:"id"`
	Name        string  `json:"name" gorm:"name"`
	Price       float64 `json:"price" gorm:"price"`
	HouseID     int64   `json:"house_id" gorm:"house_id"`
	Num         int64   `json:"num" gorm:"num"`
	Description string  `json:"description" gorm:"description"`
	Area        string  `json:"area" gorm:"area"`
	Type        string  `json:"type" gorm:"type"`
	Facility    string  `json:"facility" gorm:"facility"`
	Policy      string  `json:"policy" gorm:"policy"`
	ImageUrl    string  `json:"image_url" gorm:"image_url"`
}

// TableName 表名称
func (*Houseinfo) TableName() string {
	return "houseinfo"
}

// User undefined
type User struct {
	ID          int64     `json:"id" gorm:"id"`
	Username    string    `json:"username" gorm:"username"`
	Email       string    `json:"email" gorm:"email"`
	Password    string    `json:"password" gorm:"password"`
	City        string    `json:"city" gorm:"city"`
	Age         int       `json:"age" gorm:"age"`
	Uid         int64     `json:"uid" gorm:"uid"`
	Telephone   string    `json:"telephone" gorm:"telephone"`
	ImageUrl    string    `json:"image_url" gorm:"image_url"`
	Createdtime time.Time `json:"created_time" gorm:"created_time"`
	Updatedtime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}

// UserHouse undefined
type UserHouse struct {
	UUID        int64     `json:"uuid" gorm:"uuid"`
	UserName    string    `json:"user_name" gorm:"user_name"`
	Name        string    `json:"name" gorm:"name"`
	HouseId     int64     `json:"house_id" gorm:"house_id"`
	HouseName   string    `json:"house_name" gorm:"house_name"`
	Phone       string    `json:"phone" gorm:"phone"`
	Notes       string    `json:"notes" gorm:"notes"`
	Num         int64     `json:"num" gorm:"num"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
	Price       float64   `json:"price" gorm:"price"`
	TotalPrice  float64   `json:"total_price" gorm:"total_price"`
}

// TableName 表名称
func (*UserHouse) TableName() string {
	return "user_house"
}

// Bulletin undefined
type Bulletin struct {
	ID       int64  `json:"id" gorm:"id"`
	Title    string `json:"title" gorm:"title"`
	Content  string `json:"content" gorm:"content"`
	ImageUrl string `json:"image_url" gorm:"image_url"`
	DateTime string `json:"datetime" gorm:"datetime"`
	Place    string `json:"place" gorm:"place"`
}

// TableName 表名称
func (*Bulletin) TableName() string {
	return "bulletin"
}

// Products undefined
type Products struct {
	ID          int64   `json:"id" gorm:"id"`
	Name        string  `json:"name" gorm:"name"`
	Price       float64 `json:"price" gorm:"price"`
	Description string  `json:"description" gorm:"description"`
	ImageUrl    string  `json:"image_url" gorm:"image_url"`
	Weight      string  `json:"weight" gorm:"weight"`
	Num         int     `json:"num" gorm:"num"`
}

// TableName 表名称
func (*Products) TableName() string {
	return "products"
}

// UserProduct undefined
type UserProduct struct {
	ID          int64     `json:"id" gorm:"id"`
	Uuid        string    `json:"uuid" gorm:"uuid"`
	Username    string    `json:"user_name" gorm:"user_name"`
	Name        string    `json:"name" gorm:"name"`
	ProductId   int64     `json:"product_id" gorm:"product_id"`
	ProductName string    `json:"product_name" gorm:"product_name"`
	Quantity    int64     `json:"quantity" gorm:"quantity"`
	Address     string    `json:"address" gorm:"address"`
	Telephone   string    `json:"telephone" gorm:"telephone"`
	Price       float64   `json:"price" gorm:"price"`
	TotalPrice  float64   `json:"total_price" gorm:"total_price"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	Status      string    `json:"status" gorm:"status"`
}

// TableName 表名称
func (*UserProduct) TableName() string {
	return "user_product"
}

// House undefined
type House struct {
	ID          int64   `json:"id" gorm:"id"`
	Name        string  `json:"name" gorm:"name"`
	Price       float64 `json:"price" gorm:"price"`
	HouseId     int64   `json:"house_id" gorm:"house_id"`
	Num         int     `json:"num" gorm:"num"`
	Description string  `json:"description" gorm:"description"`
	Area        string  `json:"area" gorm:"area"`
	Type        string  `json:"type" gorm:"type"`
	Facility    string  `json:"facility" gorm:"facility"`
	Policy      string  `json:"policy" gorm:"policy"`
	ImageUrl    string  `json:"image_url" gorm:"image_url"`
}

// TableName 表名称
func (*House) TableName() string {
	return "house"
}

// Messages undefined
type Messages struct {
	ID        int64     `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	Content   string    `json:"content" gorm:"content"`
	Timestamp time.Time `json:"timestamp" gorm:"timestamp"`
	Message   string    `json:"message" gorm:"message"`
}

// TableName 表名称
func (*Messages) TableName() string {
	return "messages"
}
