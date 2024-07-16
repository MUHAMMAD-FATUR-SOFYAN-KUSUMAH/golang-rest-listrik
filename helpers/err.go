package helpers

func Err(err error, message ...string) {
	if len(message) > 0 {
		panic(message[0])
	} else {
		panic(err)
	}

}
