package bnrequest

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
