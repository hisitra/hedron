package internalclient

import (
	"github.com/hisitra/hedron/src/comcn"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var connMap = &connectionMap{
	locker:      sync.RWMutex{},
	connections: map[string]*grpc.ClientConn{},
}

type connectionMap struct {
	locker      sync.RWMutex
	connections map[string]*grpc.ClientConn
}

func (cm *connectionMap) getConn(address string) comcn.InternalClient {
	cm.locker.Lock()
	defer cm.locker.Unlock()

	conn, exists := cm.connections[address]
	if exists {
		return comcn.NewInternalClient(conn)
	}
	conn, err := cm.connectGRPC(address)
	if err != nil {
		return nil
	}
	return comcn.NewInternalClient(conn)
}

func (cm *connectionMap) refreshConn(address string) {
	cm.locker.Lock()
	defer cm.locker.Unlock()

	cm.disconnectGRPC(address)
	_, _ = cm.connectGRPC(address)
}

func (cm *connectionMap) connectGRPC(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Failed to connect to node:", address, "because:", err)
		return nil, err
	}

	cm.connections[address] = conn
	return conn, nil
}

func (cm *connectionMap) disconnectGRPC(address string) {
	conn, exists := cm.connections[address]
	if !exists {
		return
	}
	_ = conn.Close()
}
