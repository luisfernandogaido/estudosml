package main

import (
	"fmt"
	"github.com/luisfernandogaido/estudosml/dataframe"
	"log"
)

func main() {
	df, err := dataframe.NewDataFrameFloat64CSV("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	dfs := df.Divide(4, 1)
	training, test := dfs[0], dfs[1]
	fmt.Println(training)
	fmt.Println(test)
}
