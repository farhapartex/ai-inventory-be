package config

import (
	"fmt"
	"log"
	"os"

	"github.com/farhapartex/ainventory/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Error connecting to the database: %v", dbErr)
	}

	DB = db
	log.Println("Database connection established successfully")
}

func MigrateDB() {
	fmt.Println("Migrating database...")

	dbModels := []interface{}{
		&models.Permission{},
		&models.Role{},
		&models.RolePermission{},
		&models.PermissionModule{},
		&models.PermissionTemplate{},
		&models.TemplatePermission{},
		// &models.RoleHistory{},
		&models.Department{},
		&models.DepartmentRole{},
		&models.DepartmentHierarchy{},
		&models.DepartmentBudget{},
		&models.DepartmentPermission{},
		&models.DepartmentHistory{},
		&models.User{},
		&models.TokenBlacklist{},
		&models.UserHistory{},
		&models.UserPermission{},
		&models.UserProfile{},
		&models.Customer{},
		&models.Organization{},
		// &models.Product{},
		// &models.ProductCategory{},
		// &models.ProductImage{},
		// &models.ProductVariant{},
		// &models.InventoryTransaction{},
		// &models.ProductPriceHistory{},
		// &models.Supplier{},
		// &models.ProductReview{},
		// &models.Order{},
		// &models.OrderItem{},
		// &models.OrderHistory{},
		// &models.OrderPayment{},
		// &models.OrderShipment{},
		// &models.OrderShipmentItem{},
	}

	//DB.Migrator().DropTable(&models.User{})

	for _, model := range dbModels {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Error migrating model %T: %v", model, err)
		} else {
			log.Printf("Model %T migrated successfully", model)
		}

	}

	log.Println("Model migration completed!")
}
