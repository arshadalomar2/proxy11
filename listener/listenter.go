package listener

import (
	"fmt"

	"github.com/sipt/shuttle/conn"
)

type HandleFunc func(conn conn.ICtxConn)
type NewFunc func(addr string) (func(HandleFunc) error, error)

var creator = make(map[string]NewFunc)

// Register: register {key: NewFunc}
func Register(key string, f NewFunc) {
	creator[key] = f
}

// Get: get listener by key
func Get(typ, addr string) (func(HandleFunc) error, error) {
	f, ok := creator[typ]
	if !ok {
		return nil, fmt.Errorf("inbound not support: %s", typ)
	}
	return f(addr)
}