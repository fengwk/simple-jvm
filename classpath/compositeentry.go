package classpath

// CompositeEntry 使用组合模式组织Entry列表
type CompositeEntry []Entry

func NewCompositeEntry(es []Entry) *CompositeEntry {
	ce := CompositeEntry(es)
	return &ce
}

func (ce *CompositeEntry) Read(name string) ([]byte, error) {
	for _, e := range *ce {
		if cls, err := e.Read(name); err == nil {
			return cls, nil
		}
	}

	return nil, ErrClassNotFound
}
