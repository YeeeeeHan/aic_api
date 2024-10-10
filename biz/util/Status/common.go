package Status

var Success = &Status{
	code: 200,
	msg:  "success",
}

var Error = &Status{
	code: 400,
	msg:  "error",
}

var InvalidParams = &Status{
	code: 400,
	msg:  "invalid params",
}

var AuthorisationError = &Status{
	code: 401,
	msg:  "not authorised",
}

var EventPasscodeExistsError = &Status{
	code: 1001,
	msg:  "Event passcode already exists",
}

var EventAlreadyJoinedError = &Status{
	code: 1002,
	msg:  "Event already joined",
}

var EventAlreadyFullError = &Status{
	code: 1003,
	msg:  "Event is already full",
}

var EventPasscodeInvalidError = &Status{
	code: 1004,
	msg:  "Event passcode is invalid",
}
