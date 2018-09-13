package exemplos

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func PortalAction() {
	//portal-action.pdf
	df, err := dataframe.NewDataFrameFloat64CSV("./data/portal-action.csv")
	if err != nil {
		log.Fatal(err)
	}
	r, err := df.NewRegressaoMultipla([]string{"x1", "x2"}, "ganho")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Formula(2))
}
