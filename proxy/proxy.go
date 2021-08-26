package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found", string(id))
}

type UserListProxy struct {
	SomDatabase               UserList
	StackCache                UserList
	StackCapacity             int
	DidDidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {

	user, err := u.StackCache.FindUser(id)

	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidDidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomDatabase.FindUser(id)

	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.DidDidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}

}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
