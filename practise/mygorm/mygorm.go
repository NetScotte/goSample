package mygorm

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBHost   = "127.0.0.1"
	Port     = 3306
	UserName = "root"
	Password = "123456"
	DBName   = "test"
)

func GetDB(debug bool) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", UserName, Password, DBHost, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if debug {
		db = db.Debug()
	}
	return db
}

type User struct {
	Id        uint
	Username  string
	Phone     string
	IsActive  bool
	Password  string
	CreatedAt []uint8
	UpdatedAt []uint8
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) String() string {
	return fmt.Sprintf("id: %v, username: %v, phone: %v, isActive: %v", u.Id, u.Username, u.Phone, u.IsActive)
}

func Migrate() error {
	db := GetDB(true)
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Create 创建用户
func Create() error {
	db := GetDB(true)
	user := &User{
		Username: "zhangly97",
		Phone:    "13245734512",
		IsActive: true,
	}
	if err := db.Create(user).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// GetById 查询单个用户
func GetById(id uint) (*User, error) {
	db := GetDB(true)
	user := &User{
		Id: id,
	}
	if err := db.First(user).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("%v", user)
	return user, nil
}

// GetList 查询用户列表
func GetList(page int, size int, username string) ([]*User, error) {
	userList := []*User{}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	offset := (page - 1) * size
	db := GetDB(true)
	query := db.Offset(offset).Limit(size).Where("is_active = ?", true)
	if username != "" {
		query = query.Where("username like ?", fmt.Sprintf("%%%s%%", username))
	}
	err := query.Find(&userList).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, user := range userList {
		log.Info(user)
	}
	return userList, nil

}

// DisableUser 删除用户
func DisableUser(id uint) error {
	db := GetDB(true)
	err := db.Model(&User{Id: id}).Update("is_active", false).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
