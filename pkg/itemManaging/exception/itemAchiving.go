package exception

import "fmt"

type ItemAchiving struct{
	ItemID uint64
}

func (e * ItemAchiving)Error()string{
	return fmt.Sprintf("archving item id: %d failed",e.ItemID)
}