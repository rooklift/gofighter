package gofighter

func IntFromNumeric(i interface{}) int {
	switch i.(type) {
		case int:
			return int(i.(int))
		case uint:
			return int(i.(uint))
		case int8:
			return int(i.(int8))
		case uint8:
			return int(i.(uint8))
		case int16:
			return int(i.(int16))
		case uint16:
			return int(i.(uint16))
		case int32:
			return int(i.(int32))
		case uint32:
			return int(i.(uint32))
		case int64:
			return int(i.(int64))
		case uint64:
			return int(i.(uint64))
		case float32:
			return int(i.(float32))
		case float64:
			return int(i.(float64))
	}
	return 0
}
