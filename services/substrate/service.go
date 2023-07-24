package substrate

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/moody"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
	http "github.com/taubyte/http"

	goPath "path"
	goTime "time"
)

type Service interface {
	services.Service
	// Http Returns the http service attached to the Substrate
	Http() http.Service
	// Vm Returns the  VM service attached to the Substrate
	Vm() vm.Service
	// Logger returns the logger for the Substrate
	Logger() moody.Logger
	// Tns returns the Tns client attached to the Substrate
	Tns() tns.Client
	// Branch returns the branch the Substrate listens to
	Branch() string
	// Counter returns the counter service attached to the Substrate s
	Counter() CounterService
	// SmartOps returns the smartops service attached to the Substrate
	SmartOps() SmartOpsService
	Orbitals() []vm.Plugin

	Dev() bool
	Verbose() bool
	Context() context.Context

	// // P2P returns the p2p service attached to the Substrate
	// P2P() p2p.Service
}

type SmartOpsService interface {
	Service
	Run(caller SmartOpEventCaller, smartOpIds []string) (uint32, error)
}

type SmartOpEventCaller interface {
	Context() context.Context
	Type() uint32
	Application() string
	Project() (cid.Cid, error)
}

// Util is the node utilities used by the smartOps
type Util interface {
	GPU() bool
}

type Instance interface {
	Context() context.Context
	ContextCancel()

	Run(caller SmartOpEventCaller) (uint32, error)
}

type SmartOpsCache interface {
	Close()
	Get(project, application, smartOpId string, ctx context.Context) (instance Instance, ok bool)
	Put(project, application, smartOpId string, ctx context.Context, instance Instance) error
}

type CounterService interface {
	Service
	Context() context.Context
	Push(...*WrappedMetric)
}

type Metric interface {
	Reset()
	Aggregate(Metric) error
	Interface() interface{}
}

type WrappedMetric struct {
	Key    string
	Metric Metric
}

type Path interface {
	ColdStart() path
	Execution() path
	FailColdStartMetricPaths() (successCountPath string, successTimePath string, failCountPath string, failTimePath string)
	FailExecutionMetricPaths() (countPath string, timePath string)
	FailMetricPaths() (countPath string, timePath string)
	Failed() path
	Join(toJoin string) path
	Memory() path
	SmartOp(smartOpId string) path
	String() string
	Success() path
	SuccessColdStartMetricPaths() (countPath string, timePath string)
	SuccessExecutionMetricPaths() (countPath string, timePath string)
	SuccessMetricPaths() (countPath string, timePath string)
	Time() path
}

type path string

func NewPath(basePath string) path {
	return path(basePath)
}

func (c path) String() string {
	return string(c)
}

func join[v string | path](basePath Path, toJoin v) path {
	return NewPath(goPath.Join(basePath.String(), string(toJoin)))
}

func (c path) Join(toJoin string) path {
	return join(c, toJoin)
}

func (c path) Failed() path {
	return join(c, failed)
}

func (c path) Success() path {
	return join(c, success)
}

func (c path) Time() path {
	return join(c, time)
}

func (c path) Memory() path {
	return join(c, memory)
}

func (c path) Execution() path {
	return join(c, execution)
}

func (c path) ColdStart() path {
	return join(c, coldStart)
}

func (c path) SmartOp(smartOpId string) path {
	return c.Join(smartOpId).Join(smartOpId)
}

/***************************** Common Paths For Getting From Database ****************************/

func (c path) SuccessMetricPaths() (countPath, timePath string) {
	return c.Success().String(), c.Success().Time().String()
}

func (c path) SuccessColdStartMetricPaths() (countPath, timePath string) {
	return c.Success().ColdStart().Success().String(), c.Success().ColdStart().Success().Time().String()
}

func (c path) SuccessExecutionMetricPaths() (countPath, timePath string) {
	return c.Success().Execution().String(), c.Success().Execution().Time().String()
}

func (c path) FailMetricPaths() (countPath, timePath string) {
	return c.Failed().String(), c.Failed().Time().String()
}

func (c path) FailColdStartMetricPaths() (successCountPath, successTimePath, failCountPath, failTimePath string) {
	return c.Failed().ColdStart().Success().String(),
		c.Failed().ColdStart().Success().Time().String(),
		c.Failed().ColdStart().Failed().String(),
		c.Failed().ColdStart().Failed().Time().String()
}

func (c path) FailExecutionMetricPaths() (countPath, timePath string) {
	return c.Failed().Execution().String(), c.Failed().Execution().Time().String()
}

var (
	DefaultReportTime = 5 * goTime.Minute
)

const (
	time   = "t"
	memory = "m"

	failed  = "f"
	success = "s"

	coldStart = "cs"
	smartOp   = "so"
	execution = "e"
)
