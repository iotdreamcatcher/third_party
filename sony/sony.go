package sony

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func NextId() string {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := sf.NextID()
	return fmt.Sprintf("%d", id)
}
