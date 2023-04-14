package counters

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
