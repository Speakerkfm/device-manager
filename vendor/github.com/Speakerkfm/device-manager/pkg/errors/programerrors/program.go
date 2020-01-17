package programerrors

type EmailTaken struct {
	msg string
}

func (e EmailTaken) Error() string {
	return e.msg
}

func NewEmailTaken(msg string) EmailTaken {
	return EmailTaken{msg: msg}
}

type ObjectNotFound struct {
	msg string
}

func (e ObjectNotFound) Error() string {
	return e.msg
}

func NewObjectNotFound(msg string) ObjectNotFound {
	return ObjectNotFound{msg: msg}
}