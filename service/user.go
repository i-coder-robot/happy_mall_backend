package service

import (
	"errors"
	"fmt"

	"github.com/i-coder-robot/gin-demo/config"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
	"github.com/i-coder-robot/gin-demo/utils"
	uuid "github.com/satori/go.uuid"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	if req.PageSize < 1 {
		req.PageSize = config.PAGE_SIZE
	}
	return srv.Repo.List(req)
}
func (srv *UserService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}
func (srv *UserService) Exist(user model.User) *model.User {
	return srv.Repo.Exist(user)
}

func (srv *UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserID(id)
}

func (srv *UserService) Add(user model.User) (*model.User, error) {
	//根据手机号判断是否存在用户
	result := srv.Repo.ExistByMobile(user.Mobile)
	if result != nil {
		return nil, errors.New("用户已经存在")
	}
	user.UserId = uuid.NewV4().String()
	if user.Password == "" {
		user.Password = utils.Md5("123456")
	}
	user.IsDeleted = false
	user.IsLocked = false
	return srv.Repo.Add(user)
}
func (srv *UserService) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, fmt.Errorf("参数错误")
	}

	exist := srv.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("参数错误")
	}
	exist.NickName = user.NickName
	exist.Mobile = user.Mobile
	exist.Address = user.Address
	return srv.Repo.Edit(*exist)
}
func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}

	user := srv.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	user.IsDeleted = !user.IsDeleted
	return srv.Repo.Delete(*user)
}
