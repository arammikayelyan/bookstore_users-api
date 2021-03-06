package users

import (
	"fmt"

	"github.com/arammikayelyan/bookstore_users-api/utils/date_utils"
	"github.com/arammikayelyan/bookstore_users-api/utils/errors"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {

	result := userDB[user.Id]
	if result == nil {
		return errors.NotFound(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save user in the database
func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}
