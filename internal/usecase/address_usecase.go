package usecase

import (
	"fmt"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"

	"gorm.io/gorm"
)

type AddressUseCase struct {
	repo repository.AddressRepository
}

func NewAddressUseCase(repo repository.AddressRepository) *AddressUseCase {
	return &AddressUseCase{
		repo: repo,
	}
}

// Create a new address for a user
func (uc *AddressUseCase) CreateAddressUseCase(userID string, address *entities.Address) (*entities.Address, error) {
	created, err := uc.repo.CreateAddress(userID, address)
	if err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}
	return created, nil
}

// Get all addresses belonging to a user
func (uc *AddressUseCase) GetAllAddressesUseCase(userID string) ([]*entities.Address, error) {
	addresses, err := uc.repo.GetAllAddresses(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*entities.Address{}, nil // return empty list instead of error
		}
		return nil, fmt.Errorf("failed to fetch addresses: %w", err)
	}
	return addresses, nil
}

// Get a specific address (scoped to user)
func (uc *AddressUseCase) GetAddressByIDUseCase(userID, addressID string) (*entities.Address, error) {
	address, err := uc.repo.GetAddressByID(userID, addressID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch address: %w", err)
	}
	return address, nil
}

// Update a user’s address
func (uc *AddressUseCase) UpdateAddressUseCase(userID, addressID string, updated *entities.Address) (*entities.Address, error) {
	address, err := uc.repo.UpdateAddress(userID, addressID, updated)
	if err != nil {
		return nil, fmt.Errorf("failed to update address: %w", err)
	}
	return address, nil
}

// Delete a user’s address
func (uc *AddressUseCase) DeleteAddressUseCase(userID, addressID string) error {
	err := uc.repo.DeleteAddress(userID, addressID)
	if err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}
	return nil
}
