package sub

import (
	"fmt"
	"time"
)

func subE(stopAllCh chan struct{}) {
	time.Sleep(time.Second * 4)
	fmt.Print("#E#")
}
