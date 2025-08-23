package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoramAddressRepo struct {
	db *gorm.DB
}

func NewGoramAddressRepo(db *gorm.DB) *GoramAddressRepo {
	return &GoramAddressRepo{db: db}
}

func (r *GoramAddressRepo) CreateAddress(userID string, address *entities.Address) (*entities.Address, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	address.AddressBookID = uid
	if err := r.db.Create(address).Error; err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	return address, nil
}

func (r *GoramAddressRepo) GetAllAddresses(userID string) ([]*entities.Address, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var addresses []*entities.Address
	if err := r.db.Where("address_book_id = ?", uid).Find(&addresses).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch addresses: %w", err)
	}

	return addresses, nil
}

func (r *GoramAddressRepo) GetAddressByID(userID, addressID string) (*entities.Address, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}
	aid, err := uuid.Parse(addressID)
	if err != nil {
		return nil, fmt.Errorf("invalid address id: %w", err)
	}
	var address entities.Address
	err = r.db.First(&address, "id = ? AND address_book_id = ?", aid, uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("address not found or does not belong to user")
		}
		return nil, fmt.Errorf("failed to fetch address: %w", err)
	}

	return &address, nil
}

func (r *GoramAddressRepo) UpdateAddress(userID, addressID string, updated *entities.Address) (*entities.Address, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}
	aid, err := uuid.Parse(addressID)
	if err != nil {
		return nil, fmt.Errorf("invalid address id: %w", err)
	}

	updated.ID = aid
	updated.AddressBookID = uid

	result := r.db.Model(&entities.Address{}).
		Where("id = ? AND address_book_id = ?", aid, uid).
		Updates(updated)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update address: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Fetch updated address
	var saved entities.Address
	if err := r.db.First(&saved, "id = ? AND address_book_id = ?", aid, uid).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated address: %w", err)
	}

	return &saved, nil
}

func (r *GoramAddressRepo) DeleteAddress(userID, addressID string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}
	aid, err := uuid.Parse(addressID)
	if err != nil {
		return fmt.Errorf("invalid address id: %w", err)
	}

	result := r.db.Delete(&entities.Address{}, "id = ? AND address_book_id = ?", aid, uid)
	if result.Error != nil {
		return fmt.Errorf("failed to delete address: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("address not found or does not belong to user")
	}
	return nil

}
