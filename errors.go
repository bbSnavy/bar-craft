package main

func GetError(name string, err error) string {
	if *DEBUG {
		return name + "->" + err.Error()
	}
	return name
}
