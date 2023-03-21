package models

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id             int    `gorm:"column:id;type:int" json:"id,omitempty"`
	Name           string `gorm:"column:name;type:varchar(10);" json:"name,omitempty"`
	Password       string `gorm:"column:password;type:varchar(10);" json:"password"`
	FollowCount    int    `gorm:"column:follow_count;type:int" json:"follow_count,omitempty"`
	FollowerCount  int    `gorm:"column:follower_count;type:int" json:"follower_count,omitempty"`
	Token          string `json:"token"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited int    `json:"total_favorited"`
	FavoriteCount  int    `json:"favorite_count"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

//注册
func (*UserDao) Register(user *User) error {
	err := db.Select("name", "password").Create(&user).Error
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

//根据userId查找用户
//上层调用：视频填充作者、用户个人主页、token授权、评论用户填充
func (*UserDao) QueryUserById(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("未查询用户id")
	}
	return &user, nil
}

//根据id列表查找用户列表
//上层调用：关注列表、粉丝列表
func (*UserDao) MQueryUserById(ids []int) []User {
	var users []User
	err := db.Where("id in ?", ids).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return nil
	}
	return users
}

//注册时根据用户名判重
func (*UserDao) QueryUserByName(name string) (int, error) {
	var count int64
	db.Model(&User{}).Where("name = ?", name).Count(&count)
	return int(count), nil
}

//登录
func (*UserDao) Login(name string, password string) (int, error) {
	var user User
	err := db.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
