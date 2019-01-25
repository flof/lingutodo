package lingutodo

var list []string

func List() ([]string, error) {
	return list, nil
}

func Add(item string) {
	list = append(list, "Item: "+item)
}
