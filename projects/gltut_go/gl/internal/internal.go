package internal

func CallAll(a []func()) {
	if nil != a {
		for _, f := range a {
			f()
		}
	}
}
