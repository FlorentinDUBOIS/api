package db

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User structure
type User struct {
	UID       string     `gorm:"column:uid;type:uuid;primary_key;not null"      json:"uid,omitempty"`
	FirstName string     `gorm:"column:first_name;type:varchar(64);not null"    json:"first_name,omitempty"`
	LastName  string     `gorm:"column:last_name;type:varchar(64);not null"     json:"last_name,omitempty"`
	Password  string     `gorm:"column:password;type:varchar(256);not null"     json:"-"`
	Email     string     `gorm:"column:email;type:varchar(256);unique;not null" json:"email,omitempty"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null"      json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;not null"      json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp"               json:"deleted_at,omitempty"`
}

// BeforeCreate user in database
func (pUser *User) BeforeCreate(pScope *gorm.Scope) error {
	return pScope.SetColumn("UID", uuid.NewV4().String())
}

// Save user
func (pUser *User) Save(pDB *gorm.DB) error {
	if pDB.NewRecord(pUser) {
		if err := pDB.Create(pUser).Error; err != nil {
			return err
		}
	} else {
		if err := pDB.Save(pUser).Error; err != nil {
			return err
		}
	}

	return nil
}

// SetUID of user
func (pUser *User) SetUID(pUID string) {
	pUser.UID = pUID
}

// GetUID of user
func (pUser User) GetUID() string {
	return pUser.UID
}

// SetFirstName of user
func (pUser *User) SetFirstName(pFirstName string) {
	pUser.FirstName = pFirstName
}

// GetFirstName of user
func (pUser User) GetFirstName() string {
	return pUser.FirstName
}

// SetLastName of user
func (pUser *User) SetLastName(pLastName string) {
	pUser.LastName = pLastName
}

// GetLastName of user
func (pUser User) GetLastName() string {
	return pUser.LastName
}

// SetPassword of user
func (pUser *User) SetPassword(pPassword string) {
	pUser.Password = pPassword
}

// GetPassword of user
func (pUser User) GetPassword() string {
	return pUser.Password
}

// SetEmail of user
func (pUser *User) SetEmail(pEmail string) {
	pUser.Email = pEmail
}

// GetEmail of user
func (pUser User) GetEmail() string {
	return pUser.Email
}

// SetCreatedAt of user
func (pUser *User) SetCreatedAt(pCreatedAt time.Time) {
	pUser.CreatedAt = pCreatedAt
}

// GetCreatedAt of user
func (pUser User) GetCreatedAt() time.Time {
	return pUser.CreatedAt
}

// SetUpdatedAt of user
func (pUser *User) SetUpdatedAt(pUpdatedAt time.Time) {
	pUser.UpdatedAt = pUpdatedAt
}

// GetUpdatedAt of user
func (pUser User) GetUpdatedAt() time.Time {
	return pUser.UpdatedAt
}

// SetDeletedAt of user
func (pUser *User) SetDeletedAt(pDeletedAt time.Time) {
	pUser.DeletedAt = &pDeletedAt
}

// GetDeletedAt of user
func (pUser User) GetDeletedAt() *time.Time {
	return pUser.DeletedAt
}
