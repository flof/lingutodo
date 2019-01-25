package lingutodo

var list []string

func List() []string {
	return list
}

func Add(item string) {
	list = append(list, item)
}
