package uictx

func Get(value interface{}, key string) (interface{}, error) {
	return nil, nil
}

type Variables map[string]interface{}

func (self *Variables) Get(key string) (interface{}, error) {
	return nil, nil
}

func (self *Variables) Set(key string, value interface{}) {
}
