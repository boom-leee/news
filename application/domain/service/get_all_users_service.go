package service

import (
	inport "news-api/application/port/in"
	outport "news-api/application/port/out"
)

type GetAllUsersService struct {
	getAllUserPort outport.GetAllUsers
}

func NewGetAllUsersService(getAllUserPort outport.GetAllUsers) *GetAllUsersService {
	return &GetAllUsersService{getAllUserPort: getAllUserPort}
}

func (g *GetAllUsersService) Get() ([]*inport.User, error) {
	usersList, err := g.getAllUserPort.Get()
	if err != nil {
		return nil, err
	}
	return func() []*inport.User {
		result := make([]*inport.User, len(usersList))
		for i, v := range usersList {
			result[i] = MapUser(v)
		}
		return result
	}(), nil
}
