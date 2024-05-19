package outAdapter

type DummyAdapter struct {
}

func NewDummyAdapter() *DummyAdapter {
	return &DummyAdapter{}
}

func (d *DummyAdapter) GetDummy() string {
	return "Hello Hoang Ngu"
}
