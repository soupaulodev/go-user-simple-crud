package usecase

import "user-simple-crud/domain"

type UserUsecase interface {
	GetByID(id int) (*domain.User, error)
	GetAll() ([]domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int) error
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetByID(id int) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) GetAll() ([]domain.User, error) {
	return u.repo.GetAll()
}

func (u *userUsecase) Create(user *domain.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
