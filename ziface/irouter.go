package ziface

/**
router interface
*/
type IRouter interface {
	//a hook of before handle business
	PreHandle(requeest IRequest)
	//a hook of  handle business
	Handle(requeest IRequest)
	//a hook of after handle business
	PostHandle(request IRequest)
}
