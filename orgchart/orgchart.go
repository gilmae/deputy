
package orgchart

import (
	"fmt"
	"sort"
)

type Role struct {
	Id     int
	Name   string
	Parent int
}

type User struct {
	Id   int
	Name string
	Role int
}

type Organisation struct {
	Roles    map[int]Role
	roleTree map[int][]int

	Users map[int]User

	usersInRole map[int][]User
}

func NewOrganisation() *Organisation {
	return &Organisation{
		Roles:       make(map[int]Role),
		roleTree:    make(map[int][]int),
		Users:       make(map[int]User),
		usersInRole: make(map[int][]User),
	}
}

func (o *Organisation) GetSubordinates(userId int) ([]User, error) {
	found := make(map[int]User)
	o.mapSubordinates(userId, found)

	subordinates := []User{}

	for _, user := range found {
		subordinates = append(subordinates, user)
	}

	sort.Slice(subordinates, func(i, j int) bool {
		return subordinates[i].Id < subordinates[j].Id
	})

	return subordinates, nil
}

func (o *Organisation) SetRoles(roles []Role) {
	o.Roles = make(map[int]Role)
	o.roleTree = make(map[int][]int)

	o.usersInRole = make(map[int][]User)
	for _, role := range roles {
		o.Roles[role.Id] = role

		// Forgo referential integrity check for now
		if _, ok := o.roleTree[role.Parent]; !ok {
			o.roleTree[role.Parent] = []int{role.Id}
		} else {
			o.roleTree[role.Parent] = append(o.roleTree[role.Parent], role.Id)
		}

		o.usersInRole[role.Id] = make([]User, 0)
	}
	o.mapUsersToRoles()
}

func (o *Organisation) SetUsers(users []User) {
	o.Users = make(map[int]User)
	for _, user := range users {
		o.Users[user.Id] = user
	}
	o.mapUsersToRoles()
}

func (o *Organisation) mapSubordinates(userId int, found map[int]User) (error) {
	user, ok := o.Users[userId]
	
	if !ok {
		return fmt.Errorf("User not found")
	}

	subRoles, ok := o.roleTree[user.Role]
	if !ok {
		return nil
	}

	for _, roleId := range subRoles {
		if users, ok := o.usersInRole[roleId]; ok {
			for _, u := range users {
				if _, ok := found[u.Id]; !ok {
					found[u.Id] = u
					o.mapSubordinates(u.Id, found)
				}
			}
		}
	}

	return nil
}

func (o *Organisation) mapUsersToRoles() {
	for _, user := range o.Users {
		if users, ok := o.usersInRole[user.Role]; ok {
			o.usersInRole[user.Role] = append(users, user)
		}
	}
}

