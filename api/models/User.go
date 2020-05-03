package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Nickname	string		`gorm:"size:255;not null;unique" json:"nickname"`
	Email		string		`gorm:"size:255;not null;unique" json:"email"`
	Password	string		`gorm:"size:255;not null" json:"password"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) Prepare()  {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if	err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "login":
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = u.BeforeSave()
	if err != nil {
		return nil, err
	}

	err = db.Debug().Create(&u).Error
	if	err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error)  {
	var err error

	users := []User{}
	err = db.Debug().Model(&u).Limit(100).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *User) FindByUserID(db *gorm.DB, uid uint32) (*User, error) {
	var err error

	err = db.Debug().Model(&u).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}