package main

import (
	"fmt"

	"github.com/kw510/set"
)

type Permission struct {
	ObjectId string
	Roles    map[string]*set.Set
}

type Resource struct {
	Id   string
	Data map[string]any
}

type User struct {
	Id    string
	Email string
}

// Create a new permission object
func New(objectId string, roles ...string) Permission {
	rs := map[string]*set.Set{}
	for _, role := range roles {
		rs[role] = set.New()
	}

	return Permission{
		ObjectId: objectId,
		Roles:    rs,
	}
}

func main() {
	// Setup users
	alice := User{
		Id:    "id:alice",
		Email: "email:alice",
	}
	bob := User{
		Id:    "id:bob",
		Email: "email:bob",
	}
	charlie := User{
		Id:    "id:charlie",
		Email: "email:charlie",
	}

	// Setup permission object
	permission := New("ObjectId:documentXXX", "reader", "writer", "owner")
	permission.Roles["reader"].Insert(permission.Roles["writer"])
	permission.Roles["writer"].Insert(permission.Roles["owner"])

	// Insert users
	permission.Roles["owner"].Insert(alice)
	permission.Roles["reader"].Insert(bob)
	permission.Roles["writer"].Insert(charlie)

	// Check full set has updated
	fmt.Println("Owners:", permission.Roles["owner"].Flatten())
	fmt.Println("Writers:", permission.Roles["writer"].Flatten())
	fmt.Println("Readers:", permission.Roles["reader"].Flatten())
}
