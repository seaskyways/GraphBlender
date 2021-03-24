package GraphBlender

import (
	"github.com/rocketlaunchr/dataframe-go"
	"strings"
)

type dfType = *dataframe.DataFrame

type DataFrameCollection []DataFrameEntry

type DataFrameEntry struct {
	Name      string
	DataFrame dfType
}

func (dfc DataFrameCollection) Find(name string) *DataFrameEntry {
	for i := range dfc {
		if strings.Contains(dfc[i].Name, name) {
			return &dfc[i]
		}
	}
	return nil
}

func (dfc DataFrameCollection) FindExact(name string) *DataFrameEntry {
	for i := range dfc {
		if dfc[i].Name == name {
			return &dfc[i]
		}
	}
	return nil
}
