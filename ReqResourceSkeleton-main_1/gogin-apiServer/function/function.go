package function

import "time"

func GetTimeAndFormat(types string) string {

	var ret string

	now := time.Now()

	switch types {
	case "second":
		ret = now.Format("2006-01-02 15:04:05")
	case "millisecond":
		ret = now.Format("2006-01-02 15:04:05.999")

	default:
		ret = now.Format("2006-01-02 15:04:05")
	}

	return ret
}
