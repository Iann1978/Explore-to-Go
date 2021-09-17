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
	UserExist      ErrorCode = 3
	UserNotExist   ErrorCode = 4
	UserOffline    ErrorCode = 5
	SessionError   ErrorCode = 6
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
	case UserOffline:
		return "User Offline"
	case SessionError:
		return "Session Error"
	}
	return "unknown"
}
