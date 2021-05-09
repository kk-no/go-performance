package slice

func noMakeAppend(b []byte) {
	var s []byte
	for _, v := range b {
		s = append(s, v)
	}
}

func useMakeAppend(b []byte) {
	s := make([]byte, 0, len(b))
	for _, v := range b {
		s = append(s, v)
	}
}

func useMakeIndex(b []byte) {
	s := make([]byte, len(b))
	for i, v := range b {
		s[i] = v
	}
}