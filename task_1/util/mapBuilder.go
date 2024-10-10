package util

// "os"

// BuildNestedMap builds a nested map from jsonArray based on keys in argvs
func BuildNestedMap(jsonArray []map[string]interface{}, argvs []string) map[string]interface{} {
	nestedMap := make(map[string]interface{})
	for _, entry := range jsonArray {
		strippedEntry := make(map[string]interface{})
		for key, value := range entry {
			if key != argvs[0] {
				strippedEntry[key] = value
			}
		}
		value := keyConv(entry[argvs[0]])
		if nestedMap[value] == nil {
			nestedMap[value] = []map[string]interface{}{strippedEntry}
		}
		nestedMap[value] = append(nestedMap[value].([]map[string]interface{}), strippedEntry)
	}
	if len(argvs) <= 1 {
		return nestedMap
	}
	for key, value := range nestedMap {
		nestedMap[key] = BuildNestedMap(value.([]map[string]interface{}), argvs[1:])
	}
	return nestedMap
}
