package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct{
	ID uint32  
	Nickname string
	Email string
	Password string
	Created_At time.Time
	Updated_At time.Time
}

// func HashPassword(password string) ([] byte, error){
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPassword(hashedPassword, password string) error{
// 	return  bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }


// func (u *User) BeforeSave() error{
// 	hashPassword, err := HashPassword(u.Password)

// 	if err != nil {
// 		return err
		
// 	}
// 	u.Password = string(hashPassword)
// 	return nil
// }

func (u *User) Prepare(){
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Created_At = time.Now()
	u.Updated_At = time.Now()
}

func (u *User) Validate() error{
	if u.Nickname == "" {
		return errors.New("REQUIRED NAME")

	}
	if u.Email == "" {
		return errors.New("REQUIRED EMAIL")

	}
	if u.Password == "" {
		return errors.New("REQUIRED PASSWORD")

	}
	return nil
	// switch strings.ToLower(action){
	// case "update" :
	// 	if u.Nickname == "" {
	// 		return errors.New("required nickname")
			
	// 	}
	// 	if u.Email == "" {
	// 		return errors.New("required email")
	// 	}
	// 	if u.Password == "" {
	// 		return errors.New("required password")
	// 	}
	// 	if err := checkmail.ValidateFormat(u.Email); err != nil{
	// 		return errors.New("invalid email")
	// 	}

	// 	return nil
	// case "login" :
	// 	if u.Email == "" {
	// 		return errors.New("required email")
	// 	}
	// 	if u.Password == "" {
	// 		return errors.New("required password")
	// 	}
	// 	if err := checkmail.ValidateFormat(u.Email); err != nil{
	// 		return errors.New("invalid email")
	// 	}

	// 	return nil
	// default:
	// 	if u.Nickname == "" {
	// 		return errors.New("required nickname")
	// 	}
	// 	if u.Email == "" {
	// 		return errors.New("required Email")
	// 	}
	// 	if u.Password == "" {
	// 		return errors.New("required Password")
	// 	}
	// 	if err := checkmail.ValidateFormat(u.Email); err != nil{
	// 		return errors.New("invalid email")
	// 	}
	// 	return nil
	// }
}

func (u *User) SaveUser(db *gorm.DB) (*User, error){

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err	
	}
	return u, nil
	
}


func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error){

	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
		
	}

	return &users, nil
}


func (u *User) FindAUser(db *gorm.DB, uid uint32)(*User, error){

	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err){
		return &User{}, errors.New("user not found")
	}
	return u, err

}


func (u *User) UpdateUser(db *gorm.DB, uid uint32)(*User, error){

	var err error
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

	
}


func (u *User) DeleteUser( db *gorm.DB, uid uint32) (int64, error){

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	
	}
	return db.RowsAffected, nil

}