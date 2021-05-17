package services

func FindMapInterfaceByKey(list map[string]interface{}, search string) interface{} {
	for k, v := range list {
		if k == search {
			return v
		}
	}
	return nil
}