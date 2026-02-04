package domain

func (e *Mark) IsMissing() bool {
	if e == nil {
		return true
	}

	if e.Image == nil {
		return true
	}

	if len(e.Image.ImageUri) == 0 {
		return true
	}

	return false
}
