package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"house/app/logic"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func New() {
	r := gin.Default()

	r.LoadHTMLGlob("app/view/*")

	r.Static("/static", "app/images")

	r.GET("/api/getData", logic.GetData)
	r.GET("/api/searchData", logic.SearchData)
	r.POST("/api/deleteItem", logic.DeleteItem)
	r.GET("/api/getItemData", logic.GetItemData)
	r.POST("/api/editData", logic.EditData)
	r.POST("/api/addRoom", logic.AdminAddRoom)
	r.POST("/api/addActivity", logic.AdminAddActivity)
	r.POST("/api/addProduct", logic.AdminAddProduct)
	r.POST("/api/addUser", logic.AdminAddUser)
	r.GET("/api/getRoom", logic.GetRoom)
	r.GET("/api/getActivity", logic.GetActivity)
	r.GET("/api/getProduct", logic.GetProduct)
	r.GET("/api/getUser", logic.GetUser)
	r.GET("/api/getProfile", logic.GetProfile)
	r.POST("/api/updateRoom", logic.UpdateRoom)
	r.POST("/api/updateActivity", logic.UpdateActivity)
	r.POST("/api/updateProduct", logic.UpdateProduct)
	r.POST("/api/updateUser", logic.UpdateUser)
	r.GET("/order", logic.Order)
	r.GET("/chatroom", logic.Messages)
	r.GET("/profile", logic.GeRen)

	r.GET("/api/rooms", logic.GetOrderRooms)
	r.GET("/api/products", logic.GetOrderProducts)
	r.POST("/api/orders/:orderId/deliver", logic.MarkAsDelivered)
	r.GET("/api/messages", logic.GetMessages)
	r.POST("/api/messages/:messageId/reply", logic.ReplyMessage)
	//r.GET("/api/editData", logic.EditData)
	//r.GET("/api/deleteData", logic.DeleteData)

	// 用户相关路由
	r.GET("/captcha", logic.GetCaptcha)
	r.GET("/admin/captcha", logic.GetCaptchaAdmin)
	user := r.Group("/user")
	user.GET("/login", logic.UserLogin)
	user.POST("/login", logic.DoUserLogin)
	user.POST("/register", logic.DoUserRegister)
	user.POST("/update/profile", logic.UpdateProfile)
	user.GET("/profile", logic.Profile)
	user.GET("/profile/detail", logic.Loadprofile)
	user.POST("/order/house", logic.OrderHouse)
	user.POST("/order/product", logic.OrderProduct)
	user.GET("/logout", logic.UserLogout)
	user.POST("/orders/room", logic.GetUserOrder)
	user.POST("/orders/product", logic.GetUserProductOrder)
	//user.POST("/sendmail", logic.SendVerificationCode)
	//user.POST("/email-login", logic.DoUserLoginByEmail)
	//user.POST("/sends", logic.SendVerificationcodeBySms)
	//
	user.GET("/index", logic.UserIndex)
	user.GET("/announce", logic.UserAnnounce)
	user.GET("/announce/detail", logic.GetUserAnnounce)
	user.GET("/shopping", logic.GetShopping)
	user.GET("/shopping/detail", logic.GetShoppingList)
	user.GET("/chat", logic.GetChat)
	user.GET("/room/price", logic.GetBookPrice)
	user.POST("/buy/room", logic.BuyRoom)
	user.GET("/product/search", logic.SearchProduct)
	user.GET("/product/info", logic.GetProductDetail)
	user.GET("/product/info/detail", logic.GetProductInfoDetail)
	user.POST("/buy/product", logic.BuyProduct)
	user.GET("/messages", logic.GetMessage)
	user.POST("/messages", logic.SendMessage)

	////user.Use(middleware.CheckUser)
	//user.POST("/borrow", logic.BorrowBook)
	//user.POST("/return", logic.ReturnBook)
	//
	//user.POST("/updatability", logic.UploadFileHandlerAvatar)
	//user.POST("/updatable", logic.UploadFileHandlerBook)
	//
	// 管理员相关路由
	admin := r.Group("/admin")
	admin.POST("/login", logic.DoAdminLogin)
	admin.GET("/index", logic.AdminIndex)
	admin.GET("/room_list", logic.AdminGetRooms)
	admin.GET("/room_list1", logic.AdminGetRooms1)

	admin.GET("/activity_list", logic.AdminGetActivities)

	admin.GET("/product_list", logic.AdminGetProducts)
	admin.GET("/user_list", logic.GetUsers)
	admin.GET("/order-house-list", logic.GetHouseOrder)
	admin.GET("/order-product-list", logic.GetProductOrder)

	//admin.GET("/logout", logic.AdminLogout)
	admin.POST("/announce/add", logic.AddAnnounce)
	admin.POST("/announce", logic.AdminAnnounce)
	admin.POST("/add_room", logic.AdminAddRoom)
	admin.POST("/delete_room", logic.AdminDeleteRoom)
	admin.POST("/update_room", logic.AdminUpdateRoom)
	admin.POST("/add_activity", logic.AdminAddActivity)
	admin.POST("/remove_activity", logic.AdminDeleteActivity)
	admin.POST("/edit_activity", logic.AdminUpdateActivity)
	admin.POST("/add_product", logic.AdminAddProduct)
	admin.POST("/remove_product", logic.AdminDeleteProduct)
	admin.POST("/edit_product", logic.AdminUpdateProduct)
	admin.POST("/add_user", logic.AdminAddUser)
	admin.POST("/remove_user", logic.AdminDeleteUser)
	admin.POST("/edit_user", logic.AdminUpdateUser)
	admin.GET("/chatroom-list", logic.AdminGetAllMessages)
	admin.GET("/search-chatroom", logic.AdminSearchMessage)
	admin.GET("/get-profile-data", logic.AdminProfile)
	admin.POST("/update-profile", logic.AdminUpdateProfile)
	admin.GET("/logout", logic.AdminLogout)
	//// 图书相关路由
	room := r.Group("/room")
	room.GET("/list", logic.GetRooms)
	room.GET("/info", logic.RoonInfo)
	room.GET("/info/detail", logic.GetRoomDetail)
	room.GET("/search", logic.GetRoomBySearch)
	//book.GET("/admin/list", logic.GetBooksFromRedis)
	//book.POST("/add", logic.AddBook)
	//book.POST("/update", logic.SaveBook)
	//book.POST("/delete", logic.DelBook)

	// 创建 HTTP 服务器实例
	srv := &http.Server{
		Addr:    ":8086", // 设置服务器监听地址和端口号
		Handler: r,       // 设置服务器处理请求的 Handler
	}

	// 异步启动 HTTP 服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号，例如 Ctrl+C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // 监听中断信号
	<-quit                            // 等待中断信号的到来

	fmt.Println("Shutdown Server ...")

	// 创建一个超时上下文，设置等待时间为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("HTTP server shutdown: %s\n", err)
	}
	fmt.Println("Server exiting")
}
