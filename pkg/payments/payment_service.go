package payments

import (
	"errors"
	"github.com/seuc-frp-utn/api/pkg/application"
	"github.com/seuc-frp-utn/api/pkg/auth"
	"github.com/seuc-frp-utn/api/pkg/users"
	"reflect"
)

type Service struct {
	repository *application.IRepository
}

func NewService(repository *application.IRepository) *application.IService {
	var s application.IService
	s = &Service{
		repository: repository,
	}
	return &s
}

func (s Service) GetRepository() (*application.IRepository, error) {
	if s.repository == nil {
		return nil, errors.New("undefined repository")
	}
	return s.repository, nil
}

func (s Service) SetRepository(repository *application.IRepository) error {
	s.repository = repository
	return nil
}

func (s Service) Create(entity reflect.Value) (interface{}, error) {
	paymentCreate, ok := entity.Interface().(PaymentCreate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	_, err := (*users.UserService).Get(paymentCreate.Payer)
	if err != nil {
		return nil, err
	}

	payment := Payment{
		UUID:        auth.GenerateUUID(),
		Description: paymentCreate.Description,
		Amount:      paymentCreate.Amount,
		Payer:       paymentCreate.Payer,
	}

	result, err := (*s.repository).Create(payment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) Get(uuid string) (interface{}, error) {
	result, err := (*s.repository).Get(uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) GetAll() (interface{}, error) {
	result, err := (*s.repository).GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Remove(uuid string) (interface{}, error) {
	result, err := (*s.repository).Remove(uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Update(uuid string, entity reflect.Value) (interface{}, error) {
	paymentUpdate, ok := entity.Interface().(PaymentUpdate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	found, err := s.Get(uuid)
	payment, ok := found.(Payment)
	if !ok {
		return nil, errors.New("wrong format")
	}

	payment.Description = paymentUpdate.Description
	payment.Amount = paymentUpdate.Amount
	payment.Payer = paymentUpdate.Payer

	result, err := (*s.repository).Update(uuid, payment)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Find(field string, value interface{}) (interface{}, error) {
	result, err := (*s.repository).Find(field, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}
