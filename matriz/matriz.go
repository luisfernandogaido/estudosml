package matriz

import (
	"errors"
	"fmt"
)

type Matriz struct {
	elementos [][]float64
}

func New(v [][]float64) Matriz {
	return Matriz{v}
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
