package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string `json:"name" gorm:"not null;size:200;index" binding:"required"`
	SKU              string `json:"sku" gorm:"uniqueIndex;not null;size:50" binding:"required"`
	Description      string `json:"description" gorm:"type:text" binding:"required"`
	ShortDescription string `json:"short_description" gorm:"size:500"`

	CategoryID           uint     `json:"category_id" gorm:"not null;index" binding:"required"`
	Brand                string   `json:"brand" gorm:"not null;size:100;index" binding:"required"`
	Cost                 float64  `json:"cost" gorm:"type:decimal(12,2);not null" binding:"required"`
	Price                float64  `json:"price" gorm:"type:decimal(12,2);not null" binding:"required"`
	CompareAtPrice       *float64 `json:"compare_at_price" gorm:"type:decimal(12,2)"`
	Currency             string   `json:"currency" gorm:"size:3;default:'USD'"`
	Quantity             int      `json:"quantity" gorm:"not null;default:0" binding:"required"`
	LowStockThreshold    *int     `json:"low_stock_threshold" gorm:"default:10"`
	TrackQuantity        bool     `json:"track_quantity" gorm:"default:true"`
	StockStatus          string   `json:"stock_status" gorm:"size:20;default:'in_stock';check:stock_status IN ('in_stock', 'low_stock', 'out_of_stock', 'discontinued')"`
	Weight               *float64 `json:"weight" gorm:"type:decimal(8,3)"` // in kg
	Length               *float64 `json:"length" gorm:"type:decimal(8,2)"` // in cm
	Width                *float64 `json:"width" gorm:"type:decimal(8,2)"`  // in cm
	Height               *float64 `json:"height" gorm:"type:decimal(8,2)"` // in cm
	Material             string   `json:"material" gorm:"size:100"`
	Model                string   `json:"model" gorm:"size:100"`
	Colors               string   `json:"colors" gorm:"type:text"`
	Sizes                string   `json:"sizes" gorm:"type:text"`
	Status               string   `json:"status" gorm:"size:20;default:'active';check:status IN ('active', 'draft', 'archived')" binding:"required"`
	Visibility           string   `json:"visibility" gorm:"size:20;default:'visible';check:visibility IN ('visible', 'hidden')"`
	Featured             bool     `json:"featured" gorm:"default:false"`
	AvailableOnline      bool     `json:"available_online" gorm:"default:true"`
	SupplierID           uint     `json:"supplier_id" gorm:"not null;index" binding:"required"`
	SupplierSKU          string   `json:"supplier_sku" gorm:"size:100"`
	LeadTime             *int     `json:"lead_time"` // in days
	MinimumOrderQuantity *int     `json:"minimum_order_quantity" gorm:"default:1"`

	SEOTitle       string `json:"seo_title" gorm:"size:200"`
	SEODescription string `json:"seo_description" gorm:"size:500"`
	Tags           string `json:"tags" gorm:"type:text"` // JSON array of tags

	Barcode         string `json:"barcode" gorm:"size:50;index"`
	WarrantyPeriod  *int   `json:"warranty_period"` // in months
	CountryOfOrigin string `json:"country_of_origin" gorm:"size:100"`
	HSCode          string `json:"hs_code" gorm:"size:20"`

	CreatedBy uint           `json:"created_by" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Category     ProductCategory        `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Supplier     Supplier               `json:"supplier,omitempty" gorm:"foreignKey:SupplierID"`
	Creator      *User                  `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Images       []ProductImage         `json:"images,omitempty" gorm:"foreignKey:ProductID"`
	Variants     []ProductVariant       `json:"variants,omitempty" gorm:"foreignKey:ProductID"`
	Inventory    []InventoryTransaction `json:"inventory,omitempty" gorm:"foreignKey:ProductID"`
	PriceHistory []ProductPriceHistory  `json:"price_history,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	Code        string    `json:"code" gorm:"uniqueIndex;not null;size:20" binding:"required"`
	Description string    `json:"description" gorm:"size:500"`
	ParentID    *uint     `json:"parent_id" gorm:"index"`
	SortOrder   int       `json:"sort_order" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Parent   *ProductCategory  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []ProductCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products []Product         `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

type ProductImage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductID uint      `json:"product_id" gorm:"not null;index"`
	URL       string    `json:"url" gorm:"not null;size:500" binding:"required"`
	AltText   string    `json:"alt_text" gorm:"size:200"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	IsMain    bool      `json:"is_main" gorm:"default:false"`
	FileSize  *int64    `json:"file_size"` // in bytes
	MimeType  string    `json:"mime_type" gorm:"size:50"`
	CreatedAt time.Time `json:"created_at"`

	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductVariant struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ProductID  uint      `json:"product_id" gorm:"not null;index"`
	Name       string    `json:"name" gorm:"not null;size:100" binding:"required"`
	SKU        string    `json:"sku" gorm:"uniqueIndex;not null;size:50" binding:"required"`
	Price      *float64  `json:"price" gorm:"type:decimal(12,2)"`
	Quantity   int       `json:"quantity" gorm:"default:0"`
	Attributes string    `json:"attributes" gorm:"type:text"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type InventoryTransaction struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	ProductID        uint      `json:"product_id" gorm:"not null;index"`
	ProductVariantID *uint     `json:"product_variant_id" gorm:"index"`
	Type             string    `json:"type" gorm:"not null;size:20;check:type IN ('purchase', 'sale', 'adjustment', 'return', 'transfer', 'damaged', 'expired')"`
	Quantity         int       `json:"quantity" gorm:"not null"`
	UnitCost         *float64  `json:"unit_cost" gorm:"type:decimal(12,2)"`
	TotalCost        *float64  `json:"total_cost" gorm:"type:decimal(12,2)"`
	ReferenceType    string    `json:"reference_type" gorm:"size:50"` // order, purchase_order, adjustment, etc.
	ReferenceID      *uint     `json:"reference_id" gorm:"index"`
	Notes            string    `json:"notes" gorm:"size:500"`
	PerformedBy      uint      `json:"performed_by" gorm:"not null;index"`
	CreatedAt        time.Time `json:"created_at"`

	Product         Product         `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductVariant  *ProductVariant `json:"product_variant,omitempty" gorm:"foreignKey:ProductVariantID"`
	PerformedByUser User            `json:"performed_by_user,omitempty" gorm:"foreignKey:PerformedBy"`
}

type ProductPriceHistory struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	ProductID     uint      `json:"product_id" gorm:"not null;index"`
	OldPrice      *float64  `json:"old_price" gorm:"type:decimal(12,2)"`
	NewPrice      float64   `json:"new_price" gorm:"type:decimal(12,2);not null"`
	OldCost       *float64  `json:"old_cost" gorm:"type:decimal(12,2)"`
	NewCost       *float64  `json:"new_cost" gorm:"type:decimal(12,2)"`
	Reason        string    `json:"reason" gorm:"size:200"`
	ChangedBy     uint      `json:"changed_by" gorm:"not null;index"`
	CreatedAt     time.Time `json:"created_at"`
	Product       Product   `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ChangedByUser User      `json:"changed_by_user,omitempty" gorm:"foreignKey:ChangedBy"`
}

type Supplier struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Name              string         `json:"name" gorm:"not null;size:200;index" binding:"required"`
	Code              string         `json:"code" gorm:"uniqueIndex;not null;size:20" binding:"required"`
	ContactPerson     string         `json:"contact_person" gorm:"size:100"`
	Email             string         `json:"email" gorm:"size:100;index"`
	Phone             string         `json:"phone" gorm:"size:20"`
	Website           string         `json:"website" gorm:"size:255"`
	Address           string         `json:"address" gorm:"size:500"`
	City              string         `json:"city" gorm:"size:100"`
	State             string         `json:"state" gorm:"size:100"`
	ZipCode           string         `json:"zip_code" gorm:"size:20"`
	Country           string         `json:"country" gorm:"size:100"`
	TaxID             string         `json:"tax_id" gorm:"size:50"`
	PaymentTerms      string         `json:"payment_terms" gorm:"size:100;default:'Net 30'"`
	Currency          string         `json:"currency" gorm:"size:3;default:'USD'"`
	MinimumOrderValue *float64       `json:"minimum_order_value" gorm:"type:decimal(12,2)"`
	Status            string         `json:"status" gorm:"size:20;default:'active';check:status IN ('active', 'inactive', 'suspended')"`
	Rating            *float64       `json:"rating" gorm:"type:decimal(3,2);check:rating >= 0 AND rating <= 5"`
	Notes             string         `json:"notes" gorm:"type:text"`
	CreatedBy         uint           `json:"created_by" gorm:"index"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Products []Product `json:"products,omitempty" gorm:"foreignKey:SupplierID"`
	Creator  *User     `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

type ProductReview struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	ProductID      uint       `json:"product_id" gorm:"not null;index"`
	CustomerID     uint       `json:"customer_id" gorm:"not null;index"`
	Rating         int        `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title          string     `json:"title" gorm:"size:200"`
	Comment        string     `json:"comment" gorm:"type:text"`
	IsApproved     bool       `json:"is_approved" gorm:"default:false"`
	ApprovedBy     *uint      `json:"approved_by" gorm:"index"`
	ApprovedAt     *time.Time `json:"approved_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Product        Product    `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Customer       User       `json:"customer,omitempty" gorm:"foreignKey:CustomerID"` // Assuming customers are users
	ApprovedByUser *User      `json:"approved_by_user,omitempty" gorm:"foreignKey:ApprovedBy"`
}

// Model hooks and methods

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	// Update stock status based on quantity
	p.updateStockStatus()
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	p.updateStockStatus()
	return nil
}

func (p *Product) AfterUpdate(tx *gorm.DB) error {
	return p.trackPriceChange(tx)
}

func (p *Product) updateStockStatus() {
	if !p.TrackQuantity {
		return
	}

	if p.Quantity == 0 {
		p.StockStatus = "out_of_stock"
	} else if p.LowStockThreshold != nil && p.Quantity <= *p.LowStockThreshold {
		p.StockStatus = "low_stock"
	} else {
		p.StockStatus = "in_stock"
	}
}

func (p *Product) trackPriceChange(tx *gorm.DB) error {
	return nil
}

func (p *Product) AddInventoryTransaction(tx *gorm.DB, transactionType string, quantity int, unitCost *float64, referenceType string, referenceID *uint, notes string, performedBy uint) error {
	transaction := InventoryTransaction{
		ProductID:     p.ID,
		Type:          transactionType,
		Quantity:      quantity,
		UnitCost:      unitCost,
		ReferenceType: referenceType,
		ReferenceID:   referenceID,
		Notes:         notes,
		PerformedBy:   performedBy,
	}

	if unitCost != nil {
		totalCost := *unitCost * float64(quantity)
		transaction.TotalCost = &totalCost
	}

	if err := tx.Create(&transaction).Error; err != nil {
		return err
	}

	switch transactionType {
	case "purchase", "return", "adjustment":
		if quantity > 0 {
			p.Quantity += quantity
		} else {
			p.Quantity += quantity
		}
	case "sale", "damaged", "expired":
		p.Quantity -= quantity
	}

	if p.Quantity < 0 {
		p.Quantity = 0
	}
	p.updateStockStatus()

	return tx.Save(p).Error
}

func (p *Product) GetCurrentStock() int {
	return p.Quantity
}

func (p *Product) IsInStock() bool {
	return p.TrackQuantity && p.Quantity > 0
}

func (p *Product) IsLowStock() bool {
	if !p.TrackQuantity || p.LowStockThreshold == nil {
		return false
	}
	return p.Quantity <= *p.LowStockThreshold && p.Quantity > 0
}

func (p *Product) CalculateProfit() float64 {
	if p.Cost == 0 {
		return 0
	}
	return p.Price - p.Cost
}

func (p *Product) CalculateProfitMargin() float64 {
	if p.Cost == 0 {
		return 0
	}
	return ((p.Price - p.Cost) / p.Cost) * 100
}

func (pi *ProductImage) BeforeCreate(tx *gorm.DB) error {
	if pi.IsMain {
		tx.Model(&ProductImage{}).Where("product_id = ? AND is_main = ?", pi.ProductID, true).Update("is_main", false)
	}
	return nil
}

func (pi *ProductImage) BeforeUpdate(tx *gorm.DB) error {
	if pi.IsMain {
		tx.Model(&ProductImage{}).Where("product_id = ? AND id != ? AND is_main = ?", pi.ProductID, pi.ID, true).Update("is_main", false)
	}
	return nil
}

func GetDefaultCategories() []ProductCategory {
	return []ProductCategory{
		{Name: "Electronics", Code: "ELEC", Description: "Electronic devices and accessories"},
		{Name: "Clothing", Code: "CLOTH", Description: "Apparel and fashion items"},
		{Name: "Home & Garden", Code: "HOME", Description: "Home improvement and garden supplies"},
		{Name: "Sports & Outdoors", Code: "SPORT", Description: "Sports equipment and outdoor gear"},
		{Name: "Books & Media", Code: "BOOK", Description: "Books, movies, and media"},
		{Name: "Automotive", Code: "AUTO", Description: "Car parts and automotive accessories"},
		{Name: "Beauty & Health", Code: "BEAUTY", Description: "Beauty products and health items"},
		{Name: "Toys & Games", Code: "TOY", Description: "Toys, games, and children's items"},
	}
}

func GetDefaultSuppliers() []Supplier {
	return []Supplier{
		{
			Name:          "Global Electronics Inc.",
			Code:          "GEI",
			ContactPerson: "John Smith",
			Email:         "contact@globalelectronics.com",
			Phone:         "+1-555-0001",
			Country:       "United States",
			PaymentTerms:  "Net 30",
			Status:        "active",
		},
		{
			Name:          "Fashion Forward Ltd.",
			Code:          "FFL",
			ContactPerson: "Jane Doe",
			Email:         "orders@fashionforward.com",
			Phone:         "+1-555-0002",
			Country:       "United States",
			PaymentTerms:  "Net 15",
			Status:        "active",
		},
		{
			Name:          "Tech Solutions Corp.",
			Code:          "TSC",
			ContactPerson: "Mike Johnson",
			Email:         "sales@techsolutions.com",
			Phone:         "+1-555-0003",
			Country:       "United States",
			PaymentTerms:  "Net 30",
			Status:        "active",
		},
	}
}
