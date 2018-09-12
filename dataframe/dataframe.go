package dataframe

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/luisfernandogaido/estudosml/matriz"
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

func (d DataFrameFloat64) Divide(proporcoes ...float64) []DataFrameFloat64 {
	dfs := make([]DataFrameFloat64, 0, len(proporcoes))
	somaProporcoes := 0.0
	var n int
	for k := range d.Series {
		serie := d.Series[k]
		n = len(serie.Valores)
		break
	}
	for _, p := range proporcoes {
		somaProporcoes += p
	}
	p, q := 0, 0
	for _, proporcao := range proporcoes {
		q = p + int(math.Round(float64(n)*proporcao/somaProporcoes))
		df := NewDataFrameFloat64()
		for k := range d.Series {
			serie := d.Series[k]
			df.NewSerie(k, serie.Valores[p:q])
		}
		dfs = append(dfs, df)
		p = q
	}
	return dfs
}

type RegressaoLinear struct {
	A float64
	B float64
}

func (d DataFrameFloat64) NewRegressaoLinear(x, y string) (RegressaoLinear, error) {
	serieX, ok := d.Series[x]
	if !ok {
		return RegressaoLinear{}, fmt.Errorf("série %v não existe", x)
	}
	serieY, ok := d.Series[y]
	if !ok {
		return RegressaoLinear{}, fmt.Errorf("série %v não existe", y)
	}
	n := len(serieX.Valores)
	if n != len(serieY.Valores) {
		return RegressaoLinear{}, fmt.Errorf("séries x e y não possuem a mesma quantidade de elementos")
	}
	mediaX := 0.0
	mediaY := 0.0
	for i := 0; i < n; i++ {
		mediaX += serieX.Valores[i]
		mediaY += serieY.Valores[i]
	}
	mediaX /= float64(n)
	mediaY /= float64(n)
	numeradorB := 0.0
	denominadorB := 0.0
	for i := 0; i < n; i++ {
		numeradorB += (serieX.Valores[i] - mediaX) * (serieY.Valores[i] - mediaY)
		denominadorB += math.Pow(serieX.Valores[i]-mediaX, 2)
	}
	a := numeradorB / denominadorB
	b := mediaY - a*mediaX
	return RegressaoLinear{a, b}, nil
}

type RegressaoMultipla []float64

func (d DataFrameFloat64) NewRegressaoMultipla(x []string, y string) (RegressaoMultipla, error) {
	var n int
	for i, k := range x {
		serie, ok := d.Series[k]
		if !ok {
			return nil, errors.New("regressão: rótulo de série não encontrado")
		}
		if i != 0 && len(serie.Valores) != n {
			return nil, errors.New("regressão: séries não possuem a mesma quantidade")
		}
		n = len(serie.Valores)
	}
	serieY, ok := d.Series[y]
	if !ok {
		return nil, errors.New("regressão: rótulo de série não encontrado")
	}
	if len(serieY.Valores) != n {
		return nil, errors.New("regressão: séries não possuem a mesma quantidade")
	}
	vx := make([][]float64, n)
	for i := 0; i < n; i++ {
		linha := make([]float64, len(x)+1)
		linha[0] = 1
		for j, k := range x {
			linha[j+1] = d.Series[k].Valores[i]
		}
		vx[i] = linha
	}
	X := matriz.New(vx)
	vy := make([][]float64, n)
	for i := 0; i < n; i++ {
		vy[i] = []float64{serieY.Valores[i]}
	}
	Y := matriz.New(vy)
	XT := X.Transposta()
	XX, err := XT.Multiplica(X)
	if err != nil {
		return nil, fmt.Errorf("regressão: %v", err)
	}
	XXI, err := XX.Inversa()
	if err != nil {
		return nil, fmt.Errorf("regressão: %v", err)
	}
	XY, err := XT.Multiplica(Y)
	if err != nil {
		return nil, fmt.Errorf("regressão: %v", err)
	}
	fmt.Println(XXI)
	fmt.Println(XY)
	B, err := XXI.Multiplica(XY)
	if err != nil {
		return nil, fmt.Errorf("regressão: %v", err)
	}
	fmt.Println(B)
	v := B.V()
	l, _ := B.Dim()
	ret := make([]float64, l)
	for i := 0; i < l; i++ {
		ret[i] = v[i][0]
	}
	return ret, nil
}

func (r RegressaoLinear) Prediz(x float64) float64 {
	return r.A*x + r.B
}

func (r RegressaoLinear) MAE(valores []float64, observados []float64) (float64, error) {
	n := len(valores)
	if n != len(observados) {
		return 0, fmt.Errorf("número de elementos nos conjuntos de valores e observados diferem")
	}
	mae := 0.0
	for i := 0; i < n; i++ {
		mae += math.Abs(r.Prediz(valores[i]) - observados[i])
	}
	mae /= float64(n)
	return mae, nil
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
