// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}

type CreateResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Desc   string `json:"desc,optional"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount,optional"`
	Status int64  `json:"status,optional"`
}

type UpdateRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Desc   string `json:"desc,optioanl"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount,optional"`
	Status int64  `json:"status,optional"`
}

type UpdateResponse struct {
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}
