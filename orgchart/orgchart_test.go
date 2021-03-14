package orgchart

import (
	"sort"
	"testing"
)

type userInRoleTestCase struct {
	roleId        int
	expectedUsers []User
}

func TestRoleTree(t *testing.T) {
	roles := []Role{
		Role{Id: 1, Name: "doh", Parent: 0},
		Role{Id: 2, Name: "ray", Parent: 1},
		Role{Id: 3, Name: "me", Parent: 1},
		Role{Id: 4, Name: "fah", Parent: 2},
		Role{Id: 5, Name: "fah", Parent: 6},
	}

	o := NewOrganisation()
	o.SetRoles(roles)

	if subroles, ok := o.roleTree[1]; !ok {
		t.Fatalf("Could not resolve subroles for Role# %d", 1)
	} else {

		if len(subroles) != 2 {
			t.Errorf("Incorrect number of subroles for Role# %d, expected %d, got %d", 1, 2, len(subroles))
		}

		sort.Ints(subroles)
		if subroles[0] != 2 && subroles[1] != 3 {
			t.Errorf("Incorrect subroles, expected %+v, got %+v", []int{2, 3}, subroles)
		}
	}

	if subroles, ok := o.roleTree[2]; !ok {
		t.Fatalf("Could not resolve subroles for Role# %d", 2)
	} else {

		if len(subroles) != 1 {
			t.Errorf("Incorrect number of subroles for Role# %d, expected %d, got %d", 2, 1, len(subroles))
		}

		if subroles[0] != 4 {
			t.Errorf("Incorrect subroles, expected %+v, got %+v", []int{2, 3}, subroles)
		}
	}
}
func TestSetRoles(t *testing.T) {

	tests := [][]Role{
		[]Role{Role{Id: 1, Name: "foo", Parent: 0}},
		[]Role{
			Role{Id: 1, Name: "bar", Parent: 0},
			Role{Id: 2, Name: "baz", Parent: 1},
		},
	}

	for _, roles := range tests {
		o := NewOrganisation()
		o.SetRoles(roles)

		if len(o.roles) != len(roles) {
			t.Errorf("Incorrect number of roles in organisation, expected %d, got %d", len(roles), len(o.roles))
		}

		if len(o.usersInRole) != len(roles) {
			t.Errorf("Incorrect number of roles in user to role mapping in organisation, expected %d, got %d", 1, len(o.roles))
		}

		for _, role := range roles {
			r, ok := o.roles[role.Id]
			if !ok {
				t.Fatalf("Could not resolve Role# %d in organisation roles", role.Id)
			}

			if r != role {
				t.Errorf("Role does not match, expected %+v, got %+v", roles[0], r)
			}

			expectedUsers, ok := o.usersInRole[role.Id]
			if !ok {
				t.Fatalf("Could not resolve Role# %d in user to role mapping", role.Id)
			}
			if len(expectedUsers) != 0 {
				t.Errorf("Users in role not empty for Role#%d", role.Id)
			}
		}
	}
}

func TestSetRolesWithExistingUsers(t *testing.T) {
	userData := []User{
		User{Id: 1, Name: "doh", Role: 1},
		User{Id: 2, Name: "ray", Role: 2},
		User{Id: 3, Name: "me", Role: 2},
	}
	roles := []Role{
		Role{Id: 1, Name: "foo", Parent: 0},
		Role{Id: 2, Name: "bar", Parent: 1},
		Role{Id: 3, Name: "baz", Parent: 1},
	}

	tests := []userInRoleTestCase{
		{
			roleId:        1,
			expectedUsers: userData[0:1],
		},
		{
			roleId:        2,
			expectedUsers: userData[1:3],
		},
		{
			roleId:        3,
			expectedUsers: []User{},
		},
	}

	o := NewOrganisation()
	o.SetUsers(userData)
	o.SetRoles(roles)

	for _, tt := range tests {
		users, ok := o.usersInRole[tt.roleId]
		if !ok {
			t.Errorf("Could not resolve users in Role #%d", tt.roleId)
		}

		testExpectedUsers(t, users, tt.expectedUsers)
	}
}

func TestSetUsers(t *testing.T) {
	tests := [][]User{
		[]User{User{Id: 1, Name: "foo", Role: 1}},
		[]User{
			User{Id: 1, Name: "bar", Role: 1},
			User{Id: 2, Name: "baz", Role: 2},
		},
	}

	for _, expectedUsers := range tests {
		o := NewOrganisation()
		o.SetUsers(expectedUsers)

		if len(o.users) != len(expectedUsers) {
			t.Errorf("Incorrect number of expectedUsers in organisation, expected %d, got %d", len(expectedUsers), len(o.users))
		}

		for _, user := range expectedUsers {
			u, ok := o.users[user.Id]
			if !ok {
				t.Fatalf("Could not resolve User# %d in organisation expectedUsers", user.Id)
			}

			if u != user {
				t.Errorf("User does not match, expected %+v, got %+v", expectedUsers[0], u)
			}
		}
	}
}

func TestSetUsersWithExistingRoles(t *testing.T) {
	userData := []User{
		User{Id: 1, Name: "doh", Role: 1},
		User{Id: 2, Name: "ray", Role: 2},
		User{Id: 3, Name: "me", Role: 2},
	}
	roles := []Role{
		Role{Id: 1, Name: "foo", Parent: 0},
		Role{Id: 2, Name: "bar", Parent: 1},
		Role{Id: 3, Name: "baz", Parent: 1},
	}

	tests := []userInRoleTestCase{
		{
			roleId:        1,
			expectedUsers: userData[0:1],
		},
		{
			roleId:        2,
			expectedUsers: userData[1:3],
		},
		{
			roleId:        3,
			expectedUsers: []User{},
		},
	}

	o := NewOrganisation()
	o.SetRoles(roles)
	o.SetUsers(userData)

	for _, tt := range tests {
		users, ok := o.usersInRole[tt.roleId]
		if !ok {
			t.Errorf("Could not resolve expectedUsers in Role #%d", tt.roleId)
		}

		testExpectedUsers(t, users, tt.expectedUsers)
	}
}

func TestGetSubordinateUsers(t *testing.T) {
	userData := []User{
		User{Id: 1, Name: "doh", Role: 1},
		User{Id: 2, Name: "ray", Role: 4},
		User{Id: 3, Name: "me", Role: 3},
		User{Id: 4, Name: "fah", Role: 2},
		User{Id: 5, Name: "sol", Role: 5},
	}
	roles := []Role{
		Role{Id: 1, Name: "foo", Parent: 0},
		Role{Id: 2, Name: "bar", Parent: 1},
		Role{Id: 3, Name: "baz", Parent: 2},
		Role{Id: 4, Name: "qux", Parent: 3},
		Role{Id: 5, Name: "quux", Parent: 3},
	}

	tests := []struct {
		userId        int
		expectedUsers []User
	}{
		{3, []User{userData[1], userData[4]}},
		{1, userData[1:]},
	}

	o := NewOrganisation()
	o.SetRoles(roles)
	o.SetUsers(userData)

	for _, tt := range tests {
		users, err := o.GetSubordinates(tt.userId)

		if err != nil {
			t.Errorf("Received error: %s", err)
		}
		testExpectedUsers(t, users, tt.expectedUsers)
	}
}

func containsUser(expectedUsers []User, user User) bool {
	for _, u := range expectedUsers {
		if u == user {
			return true
		}
	}
	return false
}

func testExpectedUsers(t *testing.T, users []User, expectedUsers []User) {
	if len(users) != len(expectedUsers) {
		t.Errorf("Incorrect number of users, expected %d, got %d",
			len(expectedUsers),
			len(users))
	}

	for _, u := range expectedUsers {
		if !containsUser(users, u) {
			t.Errorf("User not found, %+v", u)
		}
	}
}
