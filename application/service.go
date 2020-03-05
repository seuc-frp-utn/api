package application

type IService interface {
	GetRepository() (*IRepository, error)
	SetRepository(repository *IRepository) error
	Create(entity interface{}) (interface{}, error)
	Get(uuid string) (interface{}, error)
	GetAll() (interface{}, error)
	Update(uuid string, entity interface{}) (interface{}, error)
	Remove(uuid string) (interface{}, error)
	Find(field string, value interface{}) (interface{}, error)
}