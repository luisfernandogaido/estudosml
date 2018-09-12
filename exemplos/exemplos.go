package exemplos

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func TvRadio() {
	df, err := dataframe.NewDataFrameFloat64CSV("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	partes := df.Divide(4, 1)
	treino, teste := partes[0], partes[1]
	rm, err := treino.NewRegressaoMultipla([]string{"TV", "radio"}, "sales")
	if err != nil {
		log.Fatal(err)
	}
	valores := make([][]float64, 0)
	tv := teste.Series["TV"].Valores
	radio := teste.Series["radio"].Valores
	sales := teste.Series["sales"]
	for i := range sales.Valores {
		valor := []float64{tv[i], radio[i]}
		valores = append(valores, valor)
	}
	fmt.Printf("Fórmula:\n%s\n\n", rm.Formula(2))
	mae, err := rm.MAE(valores, sales.Valores)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MAE = %.2f\n", mae)
}
