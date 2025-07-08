package v2

func safeGetStr(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}
