package service

type JWT struct {
	secret string
}

func NewJWT() *JWT {
	return &JWT{}
}


func (j *JWT) CreateUserJWT(){

}