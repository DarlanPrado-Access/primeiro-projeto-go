package model

import (
	"gorm.io/gorm"
)

//criamos o arquivo da model e o struct com o mesmo nome

//aqui é onde mapeamos cada coluna que há em nossa tabela do banco de dados
type ModelExemple struct {
	//informamos que é um model que sera utilizado pelo gorm
	gorm.Model

	//colocamos o nome do campo e seu tipo
	Camp1 string
	Camp2 int
	Camp3 bool
}
