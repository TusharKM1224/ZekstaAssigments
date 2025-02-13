package Service

import (
	Repository "github.com/TusharKM1224/Repository/MysqlRepo"
	Type "github.com/TusharKM1224/Type/MysqlType"
)

type repocon struct {
	repo Repository.DBOperations
}

type Services interface{
	CreateUser(u*Type.TableSchema) error
}

func NewService (Repo Repository.DBOperations) Services{
	return &repocon{repo: Repo}
}

func (R*repocon) CreateUser(u*Type.TableSchema) error{
	if err:=R.repo.Create(u);err!=nil{
		return err
	}
	return nil
}