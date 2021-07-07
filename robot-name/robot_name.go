package robotname

type Robot struct{}

func (Robot) Name() (string, error) {
	return "", nil
}

func (Robot) Reset() {}
