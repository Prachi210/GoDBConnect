package helper

import (
	"fmt"
	"sort"

	config "github.com/Prachi210/GoDBConnect/Consumer/Config"
	cmodel "github.com/Prachi210/GoDBConnect/Consumer/model"
)

var ParentProductSlice []cmodel.ConsumerProduct

func SortProductSlices(products []cmodel.ConsumerProduct) {
	fmt.Println("sorting slice")
	sort.Slice(products, func(i, j int) bool {
		var sortByID bool
		// sort by sold quantity
		sortByID = products[i].ID < products[j].ID
		return sortByID
	})

	//write into DB

}
func ProcessData(products []cmodel.ConsumerProduct, config config.ConfigModel) int {
	// add the data to parent
	var checkStatus int
	checkStatus = 0
	if products[0].ID == 0 {
		fmt.Println("Total Slice:", ParentProductSlice)
		SortProductSlices(ParentProductSlice)
		checkStatus = WriteProductsIntoDB(ParentProductSlice, config)
		ParentProductSlice = nil
	} else {
		ParentProductSlice = append(ParentProductSlice, products...)
	}
	return checkStatus
}
