package repos

import (
	"gorm.io/gorm"
	"pgsql/models"
)

type AccountInterface interface {
	FindByUsername(username string) (*models.Account, error)
	Save(acc *models.Account) error
}

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (acc *AccountRepo) FindByUsername(username string) (*models.Account, error) {
	var account models.Account
	err := acc.db.Debug().Where("username = ?", username).First(&account).Error
	return &account, err
}

	
func (acc *AccountRepo) Save(account *models.Account) error {
	err := acc.db.Debug().Create(&account).Error
	return err
}
