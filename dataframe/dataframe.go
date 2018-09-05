package dataframe

import (
	"encoding/csv"
	"math"
	"os"
	"sort"
	"strconv"
)

type SerieFloat64 struct {
	Indice       int
	Nome         string
	Valores      []float64
	Media        float64
	DesvioPadrao float64
	Maximo       float64
	Minimo       float64
	Q1           float64
	Q2           float64
	Q3           float64
}

func (s *SerieFloat64) AdicionaValor(valor float64) {
	if s.Valores == nil {
		s.Valores = make([]float64, 0)
	}
	s.Valores = append(s.Valores, valor)
}

func (s *SerieFloat64) AdicionaValores(valores []float64) {
	if s.Valores == nil {
		s.Valores = make([]float64, 0)
	}
	s.Valores = append(s.Valores, valores...)
}

func (s *SerieFloat64) Roda() {
	n := len(s.Valores)
	soma := 0.0
	s.Minimo = math.MaxFloat64
	s.Maximo = -s.Minimo
	for _, v := range s.Valores {
		soma += v
		s.Minimo = math.Min(v, s.Minimo)
		s.Maximo = math.Max(v, s.Maximo)
	}
	s.Media = soma / float64(len(s.Valores))
	s.DesvioPadrao = 0.0
	sorteados := make([]float64, len(s.Valores))
	copy(sorteados, s.Valores)
	sort.Float64s(sorteados)
	q1 := int(math.Round(0.25 * float64(n+1)))
	s.Q1 = sorteados[q1-1]
	if n%2 == 0 {
		s.Q2 = (sorteados[int(n/2)-1] + sorteados[int(n/2)]) / 2
	} else {
		s.Q2 = sorteados[int((n+1)/2)-1]
	}
	q3 := int(math.Round(0.75 * float64(n+1)))
	s.Q3 = sorteados[q3-1]
	for _, v := range s.Valores {
		s.DesvioPadrao += math.Pow(v-s.Media, 2)
	}
	s.DesvioPadrao = math.Pow(s.DesvioPadrao/(float64(len(s.Valores))-1), 0.5)
}

type DataFrameFloat64 struct {
	Series map[string]SerieFloat64
}

func NewDataFrameFloat64() DataFrameFloat64 {
	df := DataFrameFloat64{}
	df.Series = make(map[string]SerieFloat64)
	return df
}

func (d DataFrameFloat64) NewSerie(nome string, valores []float64) {
	d.Series[nome] = SerieFloat64{Nome: nome, Valores: valores}
}

func (d *DataFrameFloat64) Roda() {
	for k := range d.Series {
		serie := d.Series[k]
		serie.Roda()
		d.Series[k] = serie
	}
}

func NewDataFrameFloat64CSV(arq string) (DataFrameFloat64, error) {
	f, err := os.Open(arq)
	if err != nil {
		return DataFrameFloat64{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	recs, err := r.ReadAll()
	if err != nil {
		return DataFrameFloat64{}, err
	}
	df := NewDataFrameFloat64()
	titulos := recs[:1][0]
	linhas := recs[1:]
	for _, titulo := range titulos {
		df.NewSerie(titulo, nil)
	}
	for _, linha := range linhas {
		for i, coluna := range linha {
			valor, err := strconv.ParseFloat(coluna, 64)
			if err != nil {
				return DataFrameFloat64{}, err
			}
			serie := df.Series[titulos[i]]
			serie.Valores = append(serie.Valores, valor)
			df.Series[titulos[i]] = serie
		}
	}
	return df, err
}