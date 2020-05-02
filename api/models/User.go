package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
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

	err = db.Debug().Create(&u).Error
	if	err != nil {
		return nil, err
	}

	return u, nil
}
