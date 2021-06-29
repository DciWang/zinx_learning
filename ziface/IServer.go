package ziface

/**
define a server's interface
*/
type IServer interface {
	//start  the server
	Start()
	//stop the server
	Stop()
	//run the server
	Server()
}
