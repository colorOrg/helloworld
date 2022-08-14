package bean

import "encoding/json"

type Response struct {
	Code  uint16 `json:"code"`
	Error string `json:"error"`
	Msg string `json:"msg"`
}

type RelationshipResponse struct {
	 UserId int64 `json:"userId"`
	 State string `json:"state"`
	 Type string `json:"type"`
}

type UserResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (userRes * UserResponse) String() string {
	s, _ := json.Marshal(userRes)
	return string(s);
}

func (rsRes * RelationshipResponse) String() string {
	s, _ := json.Marshal(rsRes)
	return string(s);
}
