package matriz

import (
	"errors"
	"fmt"
	"math"
)

type Matriz struct {
	elementos [][]float64
}

func New(v [][]float64) Matriz {
	return Matriz{v}
}

func (m Matriz) Dim() (int, int) {
	return len(m.elementos), len(m.elementos[0])
}

func (m Matriz) Em(i, j int) float64 {
	return m.elementos[i-1][j-1]
}

func (m Matriz) String() string {
	var out string
	for i := 0; i < len(m.elementos); i++ {
		out += "|"
		for j := 0; j < len(m.elementos[i]); j++ {
			out += fmt.Sprintf("%5.2f", m.elementos[i][j]) + " "
		}
		out += "|\n"
	}
	return out
}

func (m Matriz) Soma(m2 Matriz) (Matriz, error) {
	if len(m.elementos) != len(m2.elementos) || len(m.elementos[0]) != len(m2.elementos[0]) {
		return Matriz{}, errors.New("matrizes não possuem a mesma ordem para serem somadas")
	}
	M, N := len(m.elementos), len(m.elementos[0])
	valores := make([][]float64, 0)
	for i := 0; i < M; i++ {
		linha := make([]float64, 0)
		for j := 0; j < N; j++ {
			linha = append(linha, m.elementos[i][j]+m2.elementos[i][j])
		}
		valores = append(valores, linha)
	}
	return Matriz{valores}, nil
}

func (m Matriz) Subtrai(m2 Matriz) (Matriz, error) {
	if len(m.elementos) != len(m2.elementos) || len(m.elementos[0]) != len(m2.elementos[0]) {
		return Matriz{}, errors.New("matrizes não possuem a mesma ordem para serem subtraídas")
	}
	M, N := len(m.elementos), len(m.elementos[0])
	valores := make([][]float64, 0)
	for i := 0; i < M; i++ {
		linha := make([]float64, 0)
		for j := 0; j < N; j++ {
			linha = append(linha, m.elementos[i][j]-m2.elementos[i][j])
		}
		valores = append(valores, linha)
	}
	return Matriz{valores}, nil
}

func (m Matriz) Multiplica(m2 Matriz) (Matriz, error) {
	K := len(m.elementos[0])
	if K != len(m2.elementos) {
		return Matriz{}, errors.New("matrizes com ordens incompatíveis para multiplicação")
	}
	M, N := len(m.elementos), len(m2.elementos[0])
	valores := make([][]float64, 0)
	for i := 0; i < M; i++ {
		linha := make([]float64, 0)
		for j := 0; j < N; j++ {
			elemento := 0.0
			for k := 0; k < K; k++ {
				elemento += m.elementos[i][k] * m2.elementos[k][j]
			}
			linha = append(linha, elemento)
		}
		valores = append(valores, linha)
	}
	return Matriz{valores}, nil
}

func (m Matriz) Det() (float64, error) {
	M, N := m.Dim()
	if M != N {
		return 0, errors.New("matriz precisa ser quadrada para seu determinante ser calculado")
	}
	if N == 2 {
		d := m.elementos[0][0]*m.elementos[1][1] - m.elementos[0][1]*m.elementos[1][0]
		return d, nil
	}
	d := 0.0
	for k := 0; k < N; k++ {
		valoresMenor := make([][]float64, 0)
		for i := 0; i < N; i++ {
			if i == k {
				continue
			}
			linha := make([]float64, 0)
			for j := 1; j < N; j++ {
				linha = append(linha, m.elementos[i][j])
			}
			valoresMenor = append(valoresMenor, linha)
		}
		menor := Matriz{valoresMenor}
		dMenor, err := menor.Det()
		if err != nil {
			return 0, err
		}
		d += math.Pow(-1, float64(k)+2) * m.elementos[k][0] * dMenor
	}
	return d, nil
}
