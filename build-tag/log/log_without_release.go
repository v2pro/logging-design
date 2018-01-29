//+build !release

package log

import "fmt"

func Trace(event string) {
	fmt.Println(event)
}