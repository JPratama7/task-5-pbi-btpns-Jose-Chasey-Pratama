package app

type User struct {
	ID        int32  `json:"ID" valid:"numeric,required"`
	Usernames string `json:"usernames" valid:"alphanum,required"`
	Email     string `json:"email" valid:"email,required"`
	Password  string `json:"password" valid:"minstringlength(6),alphanum,required" `
}

type Picture struct {
	ID       int32  `json:"ID"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type ReturnResponse[T any] struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
	Data    T
}

type Token struct {
	Token string `json:"token"`
}
