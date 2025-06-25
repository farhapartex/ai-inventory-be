package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     string `json:"order_id" gorm:"uniqueIndex;not null;size:20"`
	OrderNumber string `json:"order_number" gorm:"uniqueIndex;not null;size:50"`

	CustomerID    uint       `json:"customer_id" gorm:"not null;index" binding:"required"`
	CustomerNotes string     `json:"customer_notes" gorm:"type:text"`
	Status        string     `json:"status" gorm:"size:20;default:'pending';check:status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled', 'refunded')" binding:"required"`
	Priority      string     `json:"priority" gorm:"size:20;default:'normal';check:priority IN ('low', 'normal', 'high', 'urgent')"`
	OrderDate     time.Time  `json:"order_date" gorm:"not null" binding:"required"`
	RequiredDate  *time.Time `json:"required_date"`
	ShippedDate   *time.Time `json:"shipped_date"`
	DeliveredDate *time.Time `json:"delivered_date"`

	ShippingName    string `json:"shipping_name" gorm:"not null;size:200" binding:"required"`
	ShippingEmail   string `json:"shipping_email" gorm:"not null;size:150" binding:"required,email"`
	ShippingPhone   string `json:"shipping_phone" gorm:"not null;size:20" binding:"required"`
	ShippingCompany string `json:"shipping_company" gorm:"size:200"`
	ShippingAddress string `json:"shipping_address" gorm:"not null;size:500" binding:"required"`
	ShippingCity    string `json:"shipping_city" gorm:"not null;size:100" binding:"required"`
	ShippingState   string `json:"shipping_state" gorm:"not null;size:100" binding:"required"`
	ShippingZip     string `json:"shipping_zip" gorm:"not null;size:20" binding:"required"`
	ShippingCountry string `json:"shipping_country" gorm:"not null;size:100" binding:"required"`

	BillingName    string `json:"billing_name" gorm:"size:200"`
	BillingEmail   string `json:"billing_email" gorm:"size:150"`
	BillingPhone   string `json:"billing_phone" gorm:"size:20"`
	BillingCompany string `json:"billing_company" gorm:"size:200"`
	BillingAddress string `json:"billing_address" gorm:"size:500"`
	BillingCity    string `json:"billing_city" gorm:"size:100"`
	BillingState   string `json:"billing_state" gorm:"size:100"`
	BillingZip     string `json:"billing_zip" gorm:"size:20"`
	BillingCountry string `json:"billing_country" gorm:"size:100"`

	PaymentStatus    string `json:"payment_status" gorm:"size:20;default:'pending';check:payment_status IN ('pending', 'paid', 'failed', 'refunded', 'partial')" binding:"required"`
	PaymentMethod    string `json:"payment_method" gorm:"not null;size:50" binding:"required"`
	PaymentReference string `json:"payment_reference" gorm:"size:100"`

	Subtotal       float64 `json:"subtotal" gorm:"type:decimal(12,2);not null;default:0"`
	TaxRate        float64 `json:"tax_rate" gorm:"type:decimal(5,4);default:0"`
	TaxAmount      float64 `json:"tax_amount" gorm:"type:decimal(12,2);default:0"`
	ShippingCost   float64 `json:"shipping_cost" gorm:"type:decimal(12,2);default:0"`
	DiscountAmount float64 `json:"discount_amount" gorm:"type:decimal(12,2);default:0"`
	TotalAmount    float64 `json:"total_amount" gorm:"type:decimal(12,2);not null"`
	Currency       string  `json:"currency" gorm:"size:3;default:'USD'"`

	TrackingNumber  string `json:"tracking_number" gorm:"size:100"`
	ShippingMethod  string `json:"shipping_method" gorm:"size:100"`
	ShippingCarrier string `json:"shipping_carrier" gorm:"size:100"`

	OrderNotes           string `json:"order_notes" gorm:"type:text"`
	CustomerInstructions string `json:"customer_instructions" gorm:"type:text"`
	InternalNotes        string `json:"internal_notes" gorm:"type:text"`

	CreatedBy  uint           `json:"created_by" gorm:"index"`
	AssignedTo *uint          `json:"assigned_to" gorm:"index"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Customer     Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Creator      *User    `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	AssignedUser *User    `json:"assigned_user,omitempty" gorm:"foreignKey:AssignedTo"`
	//OrderItems   []OrderItem     `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
	//OrderHistory []OrderHistory  `json:"order_history,omitempty" gorm:"foreignKey:OrderID"`
	//Payments     []OrderPayment  `json:"payments,omitempty" gorm:"foreignKey:OrderID"`
	//Shipments    []OrderShipment `json:"shipments,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID               uint     `json:"id" gorm:"primaryKey"`
	OrderID          uint     `json:"order_id" gorm:"not null;index"`
	ProductID        uint     `json:"product_id" gorm:"not null;index"`
	ProductVariantID *uint    `json:"product_variant_id" gorm:"index"`
	ProductName      string   `json:"product_name" gorm:"not null;size:200"`
	ProductSKU       string   `json:"product_sku" gorm:"not null;size:50"`
	VariantName      string   `json:"variant_name" gorm:"size:100"`
	Quantity         int      `json:"quantity" gorm:"not null" binding:"required,min=1"`
	UnitPrice        float64  `json:"unit_price" gorm:"type:decimal(12,2);not null"`
	UnitCost         *float64 `json:"unit_cost" gorm:"type:decimal(12,2)"` // For profit calculation
	LineTotal        float64  `json:"line_total" gorm:"type:decimal(12,2);not null"`
	DiscountAmount   float64  `json:"discount_amount" gorm:"type:decimal(12,2);default:0"`

	Status           string    `json:"status" gorm:"size:20;default:'pending';check:status IN ('pending', 'confirmed', 'picked', 'packed', 'shipped', 'delivered', 'cancelled', 'returned')"`
	QuantityShipped  int       `json:"quantity_shipped" gorm:"default:0"`
	QuantityReturned int       `json:"quantity_returned" gorm:"default:0"`
	Notes            string    `json:"notes" gorm:"size:500"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	Order          Order           `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product        Product         `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductVariant *ProductVariant `json:"product_variant,omitempty" gorm:"foreignKey:ProductVariantID"`
}

type OrderHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OrderID     uint      `json:"order_id" gorm:"not null;index"`
	Action      string    `json:"action" gorm:"not null;size:50"`
	OldValue    string    `json:"old_value" gorm:"size:200"`
	NewValue    string    `json:"new_value" gorm:"size:200"`
	Description string    `json:"description" gorm:"size:500"`
	PerformedBy uint      `json:"performed_by" gorm:"not null;index"`
	CreatedAt   time.Time `json:"created_at"`

	Order           Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	PerformedByUser User  `json:"performed_by_user,omitempty" gorm:"foreignKey:PerformedBy"`
}

type OrderPayment struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	OrderID         uint       `json:"order_id" gorm:"not null;index"`
	PaymentMethod   string     `json:"payment_method" gorm:"not null;size:50"`
	PaymentProvider string     `json:"payment_provider" gorm:"size:50"` // Stripe, PayPal, etc.
	TransactionID   string     `json:"transaction_id" gorm:"size:100"`
	Amount          float64    `json:"amount" gorm:"type:decimal(12,2);not null"`
	Currency        string     `json:"currency" gorm:"size:3;default:'USD'"`
	Status          string     `json:"status" gorm:"size:20;check:status IN ('pending', 'processing', 'completed', 'failed', 'cancelled', 'refunded')"`
	ProcessedAt     *time.Time `json:"processed_at"`
	FailureReason   string     `json:"failure_reason" gorm:"size:500"`
	Notes           string     `json:"notes" gorm:"size:500"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderShipment struct {
	ID             uint     `json:"id" gorm:"primaryKey"`
	OrderID        uint     `json:"order_id" gorm:"not null;index"`
	TrackingNumber string   `json:"tracking_number" gorm:"size:100"`
	Carrier        string   `json:"carrier" gorm:"size:100"`
	Service        string   `json:"service" gorm:"size:100"` // Express, Standard, etc.
	Cost           *float64 `json:"cost" gorm:"type:decimal(12,2)"`
	Weight         *float64 `json:"weight" gorm:"type:decimal(8,3)"`

	ShippedDate       *time.Time `json:"shipped_date"`
	EstimatedDelivery *time.Time `json:"estimated_delivery"`
	DeliveredDate     *time.Time `json:"delivered_date"`

	Status    string    `json:"status" gorm:"size:20;check:status IN ('pending', 'picked_up', 'in_transit', 'out_for_delivery', 'delivered', 'failed', 'returned')"`
	Notes     string    `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Order         Order               `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	ShipmentItems []OrderShipmentItem `json:"shipment_items,omitempty" gorm:"foreignKey:ShipmentID"`
}

type OrderShipmentItem struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	ShipmentID  uint `json:"shipment_id" gorm:"not null;index"`
	OrderItemID uint `json:"order_item_id" gorm:"not null;index"`
	Quantity    int  `json:"quantity" gorm:"not null"`

	Shipment  OrderShipment `json:"shipment,omitempty" gorm:"foreignKey:ShipmentID"`
	OrderItem OrderItem     `json:"order_item,omitempty" gorm:"foreignKey:OrderItemID"`
}

type Customer struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CustomerID string `json:"customer_id" gorm:"uniqueIndex;not null;size:20"`

	FirstName      string `json:"first_name" gorm:"not null;size:100" binding:"required"`
	LastName       string `json:"last_name" gorm:"not null;size:100" binding:"required"`
	Email          string `json:"email" gorm:"uniqueIndex;not null;size:150" binding:"required,email"`
	Phone          string `json:"phone" gorm:"size:20"`
	SecondaryPhone string `json:"secondary_phone" gorm:"size:20"`
	Company        string `json:"company" gorm:"size:200"`
	Address        string `json:"address" gorm:"size:500"`
	City           string `json:"city" gorm:"size:100"`
	State          string `json:"state" gorm:"size:100"`
	ZipCode        string `json:"zip_code" gorm:"size:20"`
	Country        string `json:"country" gorm:"size:100;default:'United States'"`

	CustomerType string   `json:"customer_type" gorm:"size:20;default:'individual';check:customer_type IN ('individual', 'business', 'wholesale', 'vip')"`
	Status       string   `json:"status" gorm:"size:20;default:'active';check:status IN ('active', 'inactive', 'suspended', 'blocked')"`
	TaxID        string   `json:"tax_id" gorm:"size:50"`
	CreditLimit  *float64 `json:"credit_limit" gorm:"type:decimal(12,2)"`
	PaymentTerms string   `json:"payment_terms" gorm:"size:50;default:'immediate'"`
	Discount     *float64 `json:"discount" gorm:"type:decimal(5,2)"`

	MarketingOptIn         bool   `json:"marketing_opt_in" gorm:"default:true"`
	PreferredCommunication string `json:"preferred_communication" gorm:"size:20;default:'email';check:preferred_communication IN ('email', 'phone', 'sms')"`

	TotalOrders       int        `json:"total_orders" gorm:"default:0"`
	TotalSpent        float64    `json:"total_spent" gorm:"type:decimal(12,2);default:0"`
	AverageOrderValue float64    `json:"average_order_value" gorm:"type:decimal(12,2);default:0"`
	LastOrderDate     *time.Time `json:"last_order_date"`
	Notes             string     `json:"notes" gorm:"type:text"`
	Source            string     `json:"source" gorm:"size:100"` // How they found us
	ReferredBy        *uint      `json:"referred_by" gorm:"index"`

	CreatedBy uint           `json:"created_by" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	//Orders             []Order   `json:"orders,omitempty" gorm:"foreignKey:CustomerID"`
	ReferredByCustomer *Customer `json:"referred_by_customer,omitempty" gorm:"foreignKey:ReferredBy"`
	Creator            *User     `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.OrderID == "" {
		o.OrderID = generateOrderID(tx)
	}

	if o.OrderNumber == "" {
		o.OrderNumber = generateOrderNumber(tx)
	}

	return nil
}

func (o *Order) AfterCreate(tx *gorm.DB) error {
	// Create initial order history entry
	history := OrderHistory{
		OrderID:     o.ID,
		Action:      "order_created",
		NewValue:    o.Status,
		Description: "Order created",
		PerformedBy: o.CreatedBy,
	}
	return tx.Create(&history).Error
}

func (o *Order) AfterUpdate(tx *gorm.DB) error {
	// Update customer metrics when order changes
	return o.updateCustomerMetrics(tx)
}

func (o *Order) CalculateTotals(tx *gorm.DB) error {
	var items []OrderItem
	if err := tx.Where("order_id = ?", o.ID).Find(&items).Error; err != nil {
		return err
	}

	subtotal := 0.0
	for _, item := range items {
		subtotal += item.LineTotal - item.DiscountAmount
	}

	o.Subtotal = subtotal
	o.TaxAmount = subtotal * o.TaxRate
	o.TotalAmount = o.Subtotal + o.TaxAmount + o.ShippingCost - o.DiscountAmount

	return tx.Save(o).Error
}

// UpdateStatus updates order status and creates history entry
func (o *Order) UpdateStatus(tx *gorm.DB, newStatus string, performedBy uint, notes string) error {
	oldStatus := o.Status
	o.Status = newStatus

	// Create history entry
	history := OrderHistory{
		OrderID:     o.ID,
		Action:      "status_change",
		OldValue:    oldStatus,
		NewValue:    newStatus,
		Description: notes,
		PerformedBy: performedBy,
	}

	if err := tx.Create(&history).Error; err != nil {
		return err
	}

	// Update timestamps based on status
	now := time.Now()
	switch newStatus {
	case "shipped":
		o.ShippedDate = &now
	case "delivered":
		o.DeliveredDate = &now
	}

	return tx.Save(o).Error
}

// AddPayment adds a payment record to the order
func (o *Order) AddPayment(tx *gorm.DB, payment OrderPayment) error {
	payment.OrderID = o.ID
	if err := tx.Create(&payment).Error; err != nil {
		return err
	}

	// Update payment status if payment is successful
	if payment.Status == "completed" {
		var totalPaid float64
		tx.Model(&OrderPayment{}).Where("order_id = ? AND status = ?", o.ID, "completed").Select("COALESCE(SUM(amount), 0)").Scan(&totalPaid)

		if totalPaid >= o.TotalAmount {
			o.PaymentStatus = "paid"
		} else if totalPaid > 0 {
			o.PaymentStatus = "partial"
		}

		return tx.Save(o).Error
	}

	return nil
}

// GetOrderSummary returns a summary of the order
func (o *Order) GetOrderSummary() map[string]interface{} {
	return map[string]interface{}{
		"order_id":      o.OrderID,
		"order_number":  o.OrderNumber,
		"status":        o.Status,
		"total_amount":  o.TotalAmount,
		"currency":      o.Currency,
		"order_date":    o.OrderDate,
		"customer_name": o.ShippingName,
	}
}

// updateCustomerMetrics updates customer statistics
func (o *Order) updateCustomerMetrics(tx *gorm.DB) error {
	var customer Customer
	if err := tx.First(&customer, o.CustomerID).Error; err != nil {
		return err
	}

	// Count total orders and calculate total spent
	var totalOrders int64
	var totalSpent float64

	tx.Model(&Order{}).Where("customer_id = ? AND status NOT IN (?)", o.CustomerID, []string{"cancelled", "refunded"}).Count(&totalOrders)
	tx.Model(&Order{}).Where("customer_id = ? AND status NOT IN (?)", o.CustomerID, []string{"cancelled", "refunded"}).Select("COALESCE(SUM(total_amount), 0)").Scan(&totalSpent)

	// Calculate average order value
	averageOrderValue := 0.0
	if totalOrders > 0 {
		averageOrderValue = totalSpent / float64(totalOrders)
	}

	// Update customer metrics
	customer.TotalOrders = int(totalOrders)
	customer.TotalSpent = totalSpent
	customer.AverageOrderValue = averageOrderValue
	customer.LastOrderDate = &o.OrderDate

	return tx.Save(&customer).Error
}

// OrderItem methods

// BeforeCreate hook for OrderItem
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	// Calculate line total
	oi.LineTotal = float64(oi.Quantity) * oi.UnitPrice
	return nil
}

// BeforeUpdate hook for OrderItem
func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) error {
	// Recalculate line total
	oi.LineTotal = float64(oi.Quantity) * oi.UnitPrice
	return nil
}

// AfterCreate and AfterUpdate hooks for OrderItem
func (oi *OrderItem) AfterCreate(tx *gorm.DB) error {
	return oi.updateOrderTotals(tx)
}

func (oi *OrderItem) AfterUpdate(tx *gorm.DB) error {
	return oi.updateOrderTotals(tx)
}

func (oi *OrderItem) AfterDelete(tx *gorm.DB) error {
	return oi.updateOrderTotals(tx)
}

// updateOrderTotals recalculates the order totals when items change
func (oi *OrderItem) updateOrderTotals(tx *gorm.DB) error {
	var order Order
	if err := tx.First(&order, oi.OrderID).Error; err != nil {
		return err
	}
	return order.CalculateTotals(tx)
}

// Helper functions

// generateOrderID generates a unique order ID
func generateOrderID(tx *gorm.DB) string {
	for {
		orderID := fmt.Sprintf("ORD-%s", time.Now().Format("20060102")) + fmt.Sprintf("-%04d", time.Now().Unix()%10000)
		var count int64
		tx.Model(&Order{}).Where("order_id = ?", orderID).Count(&count)
		if count == 0 {
			return orderID
		}
	}
}

// generateOrderNumber generates a human-readable order number
func generateOrderNumber(tx *gorm.DB) string {
	for {
		orderNumber := fmt.Sprintf("ORD-%06d", time.Now().Unix()%1000000)
		var count int64
		tx.Model(&Order{}).Where("order_number = ?", orderNumber).Count(&count)
		if count == 0 {
			return orderNumber
		}
	}
}

// generateCustomerID generates a unique customer ID
func generateCustomerID(tx *gorm.DB) string {
	for {
		customerID := fmt.Sprintf("CUS-%05d", time.Now().Unix()%100000)
		var count int64
		tx.Model(&Customer{}).Where("customer_id = ?", customerID).Count(&count)
		if count == 0 {
			return customerID
		}
	}
}

// BeforeCreate hook for Customer
func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	// Generate CustomerID if not provided
	if c.CustomerID == "" {
		c.CustomerID = generateCustomerID(tx)
	}
	return nil
}

// GetFullName returns customer's full name
func (c *Customer) GetFullName() string {
	return c.FirstName + " " + c.LastName
}

// IsVIP checks if customer is VIP
func (c *Customer) IsVIP() bool {
	return c.CustomerType == "vip"
}

// GetDefaultOrderStatuses returns valid order statuses
func GetDefaultOrderStatuses() []string {
	return []string{"pending", "confirmed", "processing", "shipped", "delivered", "cancelled", "refunded"}
}

// GetDefaultPaymentMethods returns valid payment methods
func GetDefaultPaymentMethods() []string {
	return []string{"credit_card", "debit_card", "paypal", "bank_transfer", "cash", "check", "store_credit"}
}
