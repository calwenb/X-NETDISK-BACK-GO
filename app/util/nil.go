package util

func HasNil(objs ...any) bool {
	for i := range objs {
		if objs[i] == nil || objs[i] == "" {
			return true
		}
	}
	return false
}
func IsNil(obj any) bool {
	return obj == nil
}
