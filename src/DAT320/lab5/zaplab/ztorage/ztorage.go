// ztorage

package ztorage

import (
	"DAT320/lab5/zaplab/Store_Zap_Events"
)

type Zaps []Store_Zap_Events.StoreZapEvents

func NewZapStore() *Zaps {
	zs := make(Zaps, 0)
	return &zs
}

func (zs *Zaps) StoreZap(z Store_Zap_Events.StoreZapEvents) {
	*zs = append(*zs, z)
}

func (zs *Zaps) ComputeViewers(chName string) int {
	viewers := 0
	for _, v := range *zs {
		if v.ToCh == chName {
			viewers++
		}
		if v.FromCh == chName {
			viewers--
		}
	}

	return viewers
}
