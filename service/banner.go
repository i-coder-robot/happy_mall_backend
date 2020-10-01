package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
	uuid "github.com/satori/go.uuid"
)

type BannerSrv interface {
	List(req *query.ListQuery) (Banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(Banner model.Banner) (*model.Banner, error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

type BannerService struct {
	Repo repository.BannerRepoInterface
}

func (srv *BannerService) List(req *query.ListQuery) (banners []*model.Banner, err error) {
	return srv.Repo.List(req)
}
func (srv *BannerService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *BannerService) Get(banner model.Banner) (*model.Banner, error) {
	return srv.Repo.Get(banner)
}
func (srv *BannerService) Exist(banner model.Banner) *model.Banner {
	return srv.Repo.Exist(banner)
}

func (srv *BannerService) ExistByBannerID(id string) *model.Banner {
	return srv.Repo.ExistByBannerID(id)
}

func (srv *BannerService) Add(banner model.Banner) (*model.Banner, error) {
	if banner.BannerID == "" {
		banner.BannerID = uuid.NewV4().String()
	}
	return srv.Repo.Add(banner)
}
func (srv *BannerService) Edit(banner model.Banner) (bool, error) {
	return srv.Repo.Edit(banner)
}
func (srv *BannerService) Delete(id string) (bool, error) {
	return srv.Repo.Delete(id)
}
