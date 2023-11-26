// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"
)

type Cart struct {
	ID         int64  `json:"id"`
	ProductID  int64  `json:"product_id"`
	QtyBought  int32  `json:"qty_bought"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_price"`
	// to know which user has a cart
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Invoice struct {
	ID           int64     `json:"id"`
	SessionID    int64     `json:"session_id"`
	OrderCost    string    `json:"order_cost"`
	ShippingCost string    `json:"shipping_cost"`
	InvoiceNo    int64     `json:"invoice_no"`
	UserID       int64     `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Order struct {
	ID         int64  `json:"id"`
	ProductID  int64  `json:"product_id"`
	QtyBought  int32  `json:"qty_bought"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_price"`
	// to know which user has an order
	UserID int64 `json:"user_id"`
	// to track all orders
	SessionID int64     `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// description of the item
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Image         string    `json:"image"`
	QtyAval       int32     `json:"qty_aval"`
	ShopID        int64     `json:"shop_id"`
	CategoryID    int64     `json:"category_id"`
	SubCategoryID int64     `json:"sub_category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Shipping struct {
	ID          int64     `json:"id"`
	InvoiceID   int64     `json:"invoice_id"`
	CourierName string    `json:"courier_name"`
	Eta         int32     `json:"eta"`
	TimeLeft    time.Time `json:"time_left"`
	TimeArrive  time.Time `json:"time_arrive"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Shop struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubCategory struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
}

type User struct {
	ID             int64     `json:"id"`
	Lastname       string    `json:"lastname"`
	Firstname      string    `json:"firstname"`
	HashedPassword string    `json:"hashed_password"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Email          string    `json:"email"`
	IsAdmin        bool      `json:"is_admin"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
