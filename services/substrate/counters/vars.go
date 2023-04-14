package counters

import goTime "time"

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

/*
Current counter paths for reference

{pid}/resource type/{rid}/success
{pid}/resource type/{rid}/success/time
{pid}/resource type/{rid}/success/coldStart/time
{pid}/resource type/{rid}/success/execution/time


{pid}/resource type/{rid}/fail/
{pid}/resource type/{rid}/fail/time
{pid}/resource type/{rid}/fail/coldStart/success
{pid}/resource type/{rid}/fail/coldStart/success/time
{pid}/resource type/{rid}/fail/coldStart/fail/
{pid}/resource type/{rid}/fail/coldStart/fail/time
{pid}/resource type/{rid}/fail/execution/
{pid}/resource type/{rid}/fail/execution/time

*/
