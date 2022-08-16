package bean

type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Type uint8`json:"type"`
}


func NewUser (name string) *User{
	return &User{
		Name: name,
		Type: uint8(1),
	}
}

func (user *User) setName(n string) {
	user.Name = n
}

func (user *User) setType(t uint8)  {
	user.Type = t
}






