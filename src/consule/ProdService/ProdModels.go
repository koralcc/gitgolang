package ProdService

import "strconv"

type ProdModel struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:"pName"`
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		id,
		pname,
	}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "Product"+strconv.Itoa(i)))
	}
	return ret
}
