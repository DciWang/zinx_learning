package ziface

/**
In fact, the connection information requested by the client and the data requested by the root are packaged in a request

*/
type IRequest interface {
	//get current connection
	GetConnect() IConnection
	//get data for request
	GetData() []byte
}
