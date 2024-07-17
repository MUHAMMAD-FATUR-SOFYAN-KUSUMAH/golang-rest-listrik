package helper

func Err(err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			panic(message[0])
		}
		panic(err)
	}
}
