// zmap

/**************************************************************************/
/************************** Storing using MAP ****************************/
/************************************************************************/

package zmap

import (
	"DAT320/lab5/zaplab/SortMap"
	"DAT320/lab5/zaplab/Store_Zap_Events"
)

type Zmap struct {
	V map[string]int
}

type SortedMap struct {
	m map[int]string
}

type Resultback struct {
	Channels string
	Viewers  int
}

func NewZmapStore() *Zmap {
	zs := Zmap{make(map[string]int, 0)}
	return &zs
}

//If key already exists, decrease or increase. If key doesn't exists, create a map
//and deacrease or increase that map's value, so each channel initial value will either be -1 or 1
func (zs *Zmap) StoreZmap(z Store_Zap_Events.StoreZapEvents) {
	zs.V[z.FromCh]--
	zs.V[z.ToCh]++
}

//Method for given the channel's name, return the value associate with with it
func (zs *Zmap) ComputeViewers(chName string) int {
	return zs.V[chName]
}

//Method that sorts the MAPs key by the order of the its value, in descending order.
//This actually contain two sorting methodes.
func (zs *Zmap) ComputeTop10() []Resultback {
	var keys = make([]string, 0, len(zs.V))
	var values []int

	for k, v := range zs.V {
		values = append(values, v)
		keys = append(keys, k)
	}
	SortMap.SortM(len(keys),
		func(i, j int) bool { return (zs.V[keys[i]] > zs.V[keys[j]]) },
		func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	SortedValues := SortMap.Mapvalues(values)

	Result := make([]Resultback, 0)
	for i, _ := range keys {
		Result = append(Result, Resultback{Channels: keys[i], Viewers: SortedValues[i]})

		if i >= 9 {
			break
		}
	}
	return Result

	//For-loop for printing Channels and Viewers, in a more classy way.
	/*
		for i := 0; i < 10; i++ {
			fmt.Printf("\n%s : %d \n", keys[i], SortedValues[i])
		}
	*/
}
