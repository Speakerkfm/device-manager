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

type ObjectNotFount struct {
	msg string
}

func (e ObjectNotFount) Error() string {
	return e.msg
}

func NewObjectNotFount(msg string) ObjectNotFount {
	return ObjectNotFount{msg: msg}
}