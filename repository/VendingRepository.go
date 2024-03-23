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

	Create(domain.Vending) (domain.Vending, error)
	Update(domain.Vending) (domain.Vending, error)
	Delete(uint, domain.Base) error
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

	keys := "vendings"
	query := r.DB.Model(&vendings).
		Where("is_deleted=?", false).
		Find(&vendings)

	// Get All Vending
	config.Query(r.Redis, query, keys)

	return vendings
}

func (r *VendingRepository) GetById(id uint) domain.Vending {

	var vending domain.Vending

	keys := "vending_id_" + strconv.FormatUint(uint64(id), 10)
	query := r.DB.Model(&vending).
		Where("is_deleted=?", false).
		Where("id=?", id).
		Find(&vending)

	// Get Vending By Id
	config.Query(r.Redis, query, keys)

	return vending
}

func (r *VendingRepository) GetByName(name string) domain.Vending {

	var vending domain.Vending

	keys := "vending_name_" + name
	query := r.DB.Model(&vending).
		Where("is_deleted=?", false).
		Where("name=?", name).
		Find(&vending)

	config.Query(r.Redis, query, keys)

	return vending
}

func (r *VendingRepository) Create(vending domain.Vending) (domain.Vending, error) {

	// Create Vending
	err := r.DB.Create(&vending).Error

	return vending, err
}

func (r *VendingRepository) Update(vending domain.Vending) (domain.Vending, error) {

	// Update Vending by id
	err := r.DB.Model(&vending).Update(&vending).Error

	return vending, err
}

func (r *VendingRepository) Delete(id uint, base domain.Base) error {

	var vending domain.Vending

	// Soft Delete
	return r.DB.Model(&vending).Where("id=?", id).Updates(map[string]interface{}{
		"is_actived": false,
		"is_deleted": true,
		"updated_at": base.UpdatedAt,
		"updated_by": base.UpdatedBy,
	}).Error
}