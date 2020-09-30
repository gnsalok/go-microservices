package domain

import(
	"errors"
	"fmt"
)

var (
	users := map[int64]*User{
		123: {Id: 123, FirstName:"Alok", LastName: "Tripathi", Email:"gns123@gmail.com"},
	}
)

func GetUser(int64)(*User, error){
	if user := users[userId]; user!=nil{
		return user, nil
	}
	return nil, errors.new(fmt.Sprintf("User %v not found",userId))

}