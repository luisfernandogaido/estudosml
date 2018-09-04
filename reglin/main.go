package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type Propaganda struct {
	Tv     float64
	Radio  float64
	Jornal float64
	Vendas float64
}

type ModeloPropaganda struct {
	dad []Propaganda
	Med Propaganda
	DeP Propaganda
	Min Propaganda
	Q1  Propaganda
	Q2  Propaganda
	Q3  Propaganda
	Max Propaganda
}

func NewModeloPropaganda(arq string) (ModeloPropaganda, error) {
	f, err := os.Open(arq)
	if err != nil {
		return ModeloPropaganda{}, err
	}
	r := csv.NewReader(f)
	recs, err := r.ReadAll()
	if err != nil {
		return ModeloPropaganda{}, err
	}
	recs = recs[1:]
	modelo := ModeloPropaganda{
		Min: Propaganda{
			Tv:     math.MaxFloat64,
			Radio:  math.MaxFloat64,
			Jornal: math.MaxFloat64,
			Vendas: math.MaxFloat64,
		},
	}
	modelo.dad = make([]Propaganda, 0)
	for _, rec := range recs {
		tv, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {
			return ModeloPropaganda{}, err
		}
		radio, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return ModeloPropaganda{}, err
		}
		jornal, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			return ModeloPropaganda{}, err
		}
		vendas, err := strconv.ParseFloat(rec[3], 64)
		if err != nil {
			return ModeloPropaganda{}, err
		}
		propaganda := Propaganda{tv, radio, jornal, vendas}
		modelo.dad = append(modelo.dad, propaganda)
	}
	n := len(modelo.dad)
	for _, d := range modelo.dad {
		modelo.Med.Tv += d.Tv
		modelo.Med.Radio += d.Radio
		modelo.Med.Jornal += d.Jornal
		modelo.Med.Vendas += d.Vendas
		modelo.Min.Tv = math.Min(modelo.Min.Tv, d.Tv)
		modelo.Min.Radio = math.Min(modelo.Min.Radio, d.Radio)
		modelo.Min.Jornal = math.Min(modelo.Min.Jornal, d.Jornal)
		modelo.Min.Vendas = math.Min(modelo.Min.Vendas, d.Vendas)
		modelo.Max.Tv = math.Max(modelo.Max.Tv, d.Tv)
		modelo.Max.Radio = math.Max(modelo.Max.Radio, d.Radio)
		modelo.Max.Jornal = math.Max(modelo.Max.Jornal, d.Jornal)
		modelo.Max.Vendas = math.Max(modelo.Max.Vendas, d.Vendas)
	}
	modelo.Med.Tv /= float64(n)
	modelo.Med.Radio /= float64(n)
	modelo.Med.Jornal /= float64(n)
	modelo.Med.Vendas /= float64(n)
	for _, d := range modelo.dad {
		modelo.DeP.Tv += math.Pow(d.Tv-modelo.Med.Tv, 2)
		modelo.DeP.Radio += math.Pow(d.Radio-modelo.Med.Radio, 2)
		modelo.DeP.Jornal += math.Pow(d.Jornal-modelo.Med.Jornal, 2)
		modelo.DeP.Vendas += math.Pow(d.Vendas-modelo.Med.Vendas, 2)
	}
	modelo.DeP.Tv = math.Pow(modelo.DeP.Tv/float64(n-1), 0.5)
	modelo.DeP.Radio = math.Pow(modelo.DeP.Radio/float64(n-1), 0.5)
	modelo.DeP.Jornal = math.Pow(modelo.DeP.Jornal/float64(n-1), 0.5)
	modelo.DeP.Vendas = math.Pow(modelo.DeP.Vendas/float64(n-1), 0.5)
	vetorTv := make([]float64, 0, n)
	vetorRadio := make([]float64, 0, n)
	vetorJornal := make([]float64, 0, n)
	vetorVendas := make([]float64, 0, n)
	for _, d := range modelo.dad {
		vetorTv = append(vetorTv, d.Tv)
		vetorRadio = append(vetorRadio, d.Radio)
		vetorJornal = append(vetorJornal, d.Jornal)
		vetorVendas = append(vetorVendas, d.Vendas)
	}
	sort.Slice(vetorTv, func(i, j int) bool {
		return vetorTv[i] < vetorTv[j]
	})
	sort.Slice(vetorRadio, func(i, j int) bool {
		return vetorRadio[i] < vetorRadio[j]
	})
	sort.Slice(vetorJornal, func(i, j int) bool {
		return vetorJornal[i] < vetorJornal[j]
	})
	sort.Slice(vetorVendas, func(i, j int) bool {
		return vetorVendas[i] < vetorVendas[j]
	})
	q1 := int(math.Round(0.25 * float64(n+1)))
	q2 := int(math.Round(0.50 * float64(n+1)))
	q3 := int(math.Round(0.75 * float64(n+1)))
	modelo.Q1.Tv = vetorTv[q1]
	modelo.Q1.Radio = vetorRadio[q1-1]
	modelo.Q1.Jornal = vetorJornal[q1-1]
	modelo.Q1.Vendas = vetorVendas[q1-1]
	if n%2 == 0 {
		modelo.Q2.Tv = (vetorTv[(n/2)-1] + vetorTv[(n/2)-1]) / 2
		modelo.Q2.Radio = (vetorRadio[(n/2)-1] + vetorRadio[(n/2)-1]) / 2
		modelo.Q2.Jornal = (vetorJornal[(n/2)-1] + vetorJornal[(n/2)-1]) / 2
		modelo.Q2.Vendas = (vetorVendas[(n/2)-1] + vetorVendas[(n/2)-1]) / 2
	} else {
		modelo.Q2.Tv = vetorTv[q2-1]
		modelo.Q2.Radio = vetorRadio[q2-1]
		modelo.Q2.Jornal = vetorJornal[q2-1]
		modelo.Q2.Vendas = vetorVendas[q2-1]
	}
	modelo.Q3.Tv = vetorTv[q3-1]
	modelo.Q3.Radio = vetorRadio[q3-1]
	modelo.Q3.Jornal = vetorJornal[q3-1]
	modelo.Q3.Vendas = vetorVendas[q3-1]
	return modelo, nil
}

func (m ModeloPropaganda) CovVendas() Propaganda {
	propaganda := Propaganda{}
	for i := range m.dad {
		propaganda.Tv += (m.dad[i].Tv - m.Med.Tv) * (m.dad[i].Vendas - m.Med.Vendas)
		propaganda.Radio += (m.dad[i].Radio - m.Med.Radio) * (m.dad[i].Vendas - m.Med.Vendas)
		propaganda.Jornal += (m.dad[i].Jornal - m.Med.Jornal) * (m.dad[i].Vendas - m.Med.Vendas)
		propaganda.Vendas += (m.dad[i].Vendas - m.Med.Vendas) * (m.dad[i].Vendas - m.Med.Vendas)
	}
	n := len(m.dad)
	propaganda.Tv /= float64(n)
	propaganda.Radio /= float64(n)
	propaganda.Jornal /= float64(n)
	propaganda.Vendas /= float64(n)
	return propaganda
}

func (m ModeloPropaganda) RegressaoLinear() {

}

func main() {
	modelo, err := NewModeloPropaganda("./Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", modelo.CovVendas())
}
