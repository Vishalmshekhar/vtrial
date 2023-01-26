package Controllers

type I_Controller interface{
	Routes()
	healthCheck()
	UseSavePage()
	UseComputeResult()
}

