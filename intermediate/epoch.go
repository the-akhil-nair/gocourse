package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// 00:00:00 UTC on Jan 1, 1970
	// Before epoch have negative unix time
	t1 := time.Date(1969, time.January, 31, 23, 59, 59, 0, time.UTC)
	fmt.Println("Unix Time for", t1, "is", t1.Unix())

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)
	// Second argument is nanoseconds precision
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02"))

}
