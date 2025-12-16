package exception

import "fmt"

type ItemQuanitityNotEnough struct{
	ItemID uint64
}

func (e *ItemQuanitityNotEnough) Error() string {
	return fmt.Sprintf("itemID: %d is not enough",e.ItemID)
}

