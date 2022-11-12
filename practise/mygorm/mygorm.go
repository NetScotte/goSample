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
		fmt.Println(err)
		return nil
	}
	if debug {
		db = db.Debug()
	}
	return db
}

type User struct {
	Id        uint `gorm:"primaryKey"`
	Username  string
	Phone     string
	IsActive  bool
	Password  string
	CreatedAt []uint8
	UpdatedAt []uint8
}

type People struct {
	Id       uint
	Username string
	Phone    string
}

type Orders struct {
	Id        uint64
	Username  string
	Name      string
	Price     uint64
	Count     uint64
	Total     uint64
	CreatedAt []uint8
	UpdatedAt []uint8
}

func (o *Orders) TableName() string {
	return "orders"
}

func (p *People) String() string {
	return fmt.Sprintf("Id: %v, Username: %v, Phone: %v\n", p.Id, p.Username, p.Phone)
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) String() string {
	return fmt.Sprintf("id: %v, username: %v, phone: %v, isActive: %v, createdAt: %v\n", u.Id, u.Username, u.Phone, u.IsActive, string(u.CreatedAt))
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
	if err := db.Select([]string{"username", "password", "created_at"}).Take(user).Error; err != nil {
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

// GetPeople 从数据库中获取数据并放到另一个对象中
func GetPeople() {
	var people []*People
	result := GetDB(false).Model(&User{}).Scan(&people)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(people)
}

func GetUserOrder() {
	type UserOrder struct {
		Username  string
		Phone     string
		Name      string
		Price     uint64
		Count     uint64
		Total     uint64
		CreatedAt []uint8
	}

	var userOrders []*UserOrder

	// select a.username, a.phone, b.name, b.price, b.count, b.total, b.created_at
	//	from user a, orders b
	// where a.username = b.username
	// and a.is_active = 1
	DB := GetDB(true)
	err := DB.Raw("select a.username, a.phone, b.name, b.price, b.count, b.total, b.created_at "+
		"from user a, orders b "+
		"where "+
		"a.username = b.username "+
		"and a.is_active = ?", 1).Scan(&userOrders).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, userOrder := range userOrders {
		fmt.Printf("username: %v, name: %v, price: %v, createdAt: %s\n", userOrder.Username, userOrder.Name, userOrder.Price, userOrder.CreatedAt)
	}

}

func Random() {
	DB := GetDB(true)
	records := &User{
		Username: "liufy47",
		Phone:    "13222734512",
	}
	result := DB.Model(&User{}).Where(records).Update("is_active", false)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Printf("affected row: %v\n", result.RowsAffected)
}
