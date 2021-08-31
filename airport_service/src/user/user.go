package user


type User struct {
	userid int32
	name string
	x float64
	y float64
}

type UserDatabase struct {

}

type UserSet interface {
	HasUser(username string) (int32, error)
}

func (users UserDatabase) HasUser(username string) (int32, error) {
	return 0, nil
}
