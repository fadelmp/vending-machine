package repository

import (
	"strconv"
	"vending-machine/config"
	"vending-machine/domain"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type VendingRepositoryContract interface {
	GetAll() []domain.Vending
	GetById(uint) domain.Vending
	GetByName(string) domain.Vending

	Create(*domain.Vending) error
	Update(*domain.Vending) error
	Delete(*domain.Vending) error
}

// Class
type VendingRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewVendingRepository(DB *gorm.DB, Redis *redis.Client) *VendingRepository {
	return &VendingRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *VendingRepository) GetAll() []domain.Vending {

	var vendings []domain.Vending

	// Get All Data
	query := r.DB.Model(&vendings).
		Unscoped().
		Where("is_deleted=?", false).
		Find(&vendings)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "vendings")

	return vendings
}

func (r *VendingRepository) GetById(id uint) domain.Vending {

	var vending domain.Vending

	// Get Data By Id
	query := r.DB.Model(&vending).
		Unscoped().
		Where("is_deleted=?", false).
		Where("id=?", id).
		Find(&vending)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "vending_id_"+strconv.FormatUint(uint64(id), 10))

	return vending
}

func (r *VendingRepository) GetByName(name string) domain.Vending {

	var vending domain.Vending

	// Get Data By Name
	query := r.DB.Model(&vending).
		Unscoped().
		Where("is_deleted=?", false).
		Where("name=?", name).
		Find(&vending)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "vending_name_"+name)

	return vending
}

func (r *VendingRepository) Create(vending *domain.Vending) error {

	// Flush Vending Cache
	config.FlushData(r.Redis, "vending*")

	// Create Vending
	return r.DB.Create(&vending).Error

}

func (r *VendingRepository) Update(vending *domain.Vending) error {

	// Flush Vending Cache
	config.FlushData(r.Redis, "vending*")

	// Update Vending
	return r.DB.Model(&vending).Unscoped().Update(&vending).Error
}

func (r *VendingRepository) Delete(vending *domain.Vending) error {

	// Flush Vending Cache
	config.FlushData(r.Redis, "vending*")

	// Delete Vending
	return r.DB.Model(&vending).Unscoped().Update(&vending).Error
}
