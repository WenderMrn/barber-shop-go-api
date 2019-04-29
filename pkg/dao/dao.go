package dao

//IGenericDAO interface
type IGenericDAO interface {
	Create(obj interface{}) error
	Find(ID uint) error
	Update(obj interface{}) error
	Delete(obj interface{}) error
	Begin() error
	Commit() error
	Rollback() error
}

//DAO struct
type DAO struct {
	IGenericDAO
}
