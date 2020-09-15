package lib

func GetKeys(object map[string]interface{}) []string {
	keys := make([]string, 0, len(object))
	for k := range object {
		keys = append(keys, k)
	}
	return keys
}
