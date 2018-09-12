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

func (m Matriz) V() [][]float64 {
	v := make([][]float64, 0)
	for _, linha := range m.elementos {
		copia := make([]float64, len(linha))
		copy(copia, linha)
		v = append(v, copia)
	}
	return v
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
			out += fmt.Sprintf("%5.4f", m.elementos[i][j]) + " "
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

func (m Matriz) Transposta() Matriz {
	elementos := make([][]float64, 0)
	linhas, colunas := m.Dim()
	for j := 0; j < colunas; j++ {
		linha := make([]float64, 0)
		for i := 0; i < linhas; i++ {
			linha = append(linha, m.elementos[i][j])
		}
		elementos = append(elementos, linha)
	}
	return Matriz{elementos}
}

func (m Matriz) TrocaLinhas(l1, l2 int) {
	aux := m.elementos[l1-1]
	m.elementos[l1-1] = m.elementos[l2-1]
	m.elementos[l2-1] = aux
}

func (m Matriz) MultiplicaLinha(l int, c float64) {
	_, colunas := m.Dim()
	for j := 0; j < colunas; j++ {
		m.elementos[l-1][j] *= c
	}
}

func (m Matriz) SomaLinhas(l, l1, l2 int) {
	_, colunas := m.Dim()
	for j := 0; j < colunas; j++ {
		m.elementos[l-1][j] = m.elementos[l1-1][j] + m.elementos[l2-1][j]
	}
}

func (m Matriz) Indentidade() (Matriz, error) {
	M, N := m.Dim()
	if M != N {
		return Matriz{}, errors.New("matriz identidade não pode ser calculada por não ser quadrada")
	}
	linhas := make([][]float64, 0, M)
	for i := 0; i < M; i++ {
		linha := make([]float64, 0, N)
		for j := 0; j < N; j++ {
			valor := 0.0
			if i == j {
				valor = 1.0
			}
			linha = append(linha, valor)
		}
		linhas = append(linhas, linha)
	}
	return Matriz{linhas}, nil
}

//Calcula a matriz inversa.
// Aprendi mesmo aqui: http://www.gregthatcher.com/Mathematics/GaussJordan.aspx
// Outra excelente fonte é https://www.intmath.com/matrices-determinants/inverse-matrix-gauss-jordan-elimination.php
func (m Matriz) Inversa() (Matriz, error) {
	ide, err := m.Indentidade()
	if err != nil {
		return Matriz{}, err
	}
	e1 := m.V()
	e2 := ide.elementos
	dim, _ := m.Dim()
	for i := 0; i < dim; i++ {
		iMax := i
		abs := 0.0
		for k := i + 1; k < dim; k++ {
			valor := math.Abs(e1[k][i])
			if math.Abs(valor) > abs {
				abs = valor
				iMax = k
			}
		}
		if iMax != i {
			aux1 := e1[i]
			e1[i] = e1[iMax]
			e1[iMax] = aux1
			aux2 := e2[i]
			e2[i] = e2[iMax]
			e2[iMax] = aux2
		}
		if e1[i][i] != 1 {
			divisor := e1[i][i]
			for j := 0; j < dim; j++ {
				e1[i][j] /= divisor
				e2[i][j] /= divisor
			}
		}
		for k := i + 1; k < dim; k++ {
			multiplicador := -(e1[k][i] / e1[i][i])
			linhaZerada := true
			for j := 0; j < dim; j++ {
				e1[k][j] += multiplicador * e1[i][j]
				e2[k][j] += multiplicador * e2[i][j]
				if e1[k][j] != 0.0 {
					linhaZerada = false
				}
			}
			if linhaZerada {
				return Matriz{}, errors.New("matriz não é inversível")
			}
		}
	}
	for i := dim - 1; i >= 0; i-- {
		for k := i - 1; k >= 0; k-- {
			multiplicador := -e1[k][i]
			linhaZerada := true
			for j := 0; j < dim; j++ {
				e1[k][j] += multiplicador * e1[i][j]
				e2[k][j] += multiplicador * e2[i][j]
				if e1[k][j] != 0.0 {
					linhaZerada = false
				}
			}
			if linhaZerada {
				return Matriz{}, errors.New("matriz não é inversível")
			}
		}
	}
	return ide, nil
}
