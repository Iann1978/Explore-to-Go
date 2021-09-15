package data

type LoginResp2 struct {
	ErrorString string
	//Username    string
	//user        *user.User
}

type ErrorCode int32

const (
	NoError        ErrorCode = 0
	UnknownError   ErrorCode = 1
	ParameterError ErrorCode = 2
	UserExist                = 3
	UserNotExist             = 4
)

func (e ErrorCode) String() string {
	switch e {
	case NoError:
		return "No Error"
	case UnknownError:
		return "Unknown Error"
	case ParameterError:
		return "Parameter Error"
	case UserExist:
		return "User Already Exist"
	case UserNotExist:
		return "User Not Exist"
	}
	return "unknown"
}
