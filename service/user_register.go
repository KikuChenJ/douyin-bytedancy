package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type UserRegisterFlow struct {
	Username string
	Password string
	User     *models.User
	UserId   int
	Token    string
}

func UserRegister(username string, password string) (*models.User, error) {
	return NewUserRegisterFlow(username, password).Do()
}

func NewUserRegisterFlow(username string, password string) *UserRegisterFlow {
	return &UserRegisterFlow{Username: username, Password: password}
}

func (f *UserRegisterFlow) Do() (*models.User, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.Register(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.User, nil
}

func (f *UserRegisterFlow) checkParam() error {
	if f.Username == "" {
		return errors.New("用户名为空")
	}
	if len(f.Username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if f.Password == "" {
		return errors.New("密码为空")
	}
	if len(f.Password) > MaxPasswordLength || len(f.Password) < MinPasswordLength {
		return errors.New("密码长度应为6-20个字符")
	}
	return nil
}

func (f *UserRegisterFlow) Register() error {
	userDao := models.NewUserDaoInstance()

	if count, err := userDao.QueryUserByName(f.Username); err == nil && count > 0 {
		return errors.New("用户名已存在")
	}
	user := &models.User{
		Name:     f.Username,
		Password: f.Password,
	}
	err := userDao.Register(user)
	if err != nil {
		return err
	}
	token, err := jwt.GenToken(user.Id)
	if err != nil {
		return err
	}
	f.Token = token
	f.UserId = user.Id
	return nil
}

func (f *UserRegisterFlow) packData() error {
	f.User = &models.User{
		Id:    f.UserId,
		Token: f.Token,
	}
	return nil
}
