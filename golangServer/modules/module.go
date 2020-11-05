package modules

import (
	"google.golang.org/grpc"
)

type ModuleContext struct {
	Reference string
}

type IModule interface {
	RegisterAsGrpcService(server grpc.ServiceRegistrar)
	GetContext() ModuleContext
	setAsActive()
	isActive() bool
}

////////////////////
// Error Handling //
////////////////////
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString{text}
}