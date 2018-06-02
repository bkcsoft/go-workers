package workers

import (
	"testing"
	"os"

	"github.com/customerio/gospec"
)

// You will need to list every spec in a TestXxx method like this,
// so that gotest can be used to run the specs. Later GoSpec might
// get its own command line tool similar to gotest, but for now this
// is the way to go. This shouldn't require too much typing, because
// there will be typically only one top-level spec per class/feature.

var RedisHost = "localhost:6379"

func TestAllSpecs(t *testing.T) {
	tmp := os.Getenv("REDIS_HOST")
	if len(tmp) > 0 {
		RedisHost = tmp
	}
	
	r := gospec.NewRunner()

	r.Parallel = false

	r.BeforeEach = func() {
		Configure(map[string]string{
			"server":   RedisHost,
			"process":  "1",
			"database": "15",
			"pool":     "1",
		})

		conn := Config.Pool.Get()
		conn.Do("flushdb")
		conn.Close()
	}

	// List all specs here
	r.AddSpec(WorkersSpec)
	r.AddSpec(ConfigSpec)
	r.AddSpec(MsgSpec)
	r.AddSpec(FetchSpec)
	r.AddSpec(WorkerSpec)
	r.AddSpec(ManagerSpec)
	r.AddSpec(ScheduledSpec)
	r.AddSpec(EnqueueSpec)
	r.AddSpec(MiddlewareSpec)
	r.AddSpec(MiddlewareRetrySpec)
	r.AddSpec(MiddlewareStatsSpec)

	// Run GoSpec and report any errors to gotest's `testing.T` instance
	gospec.MainGoTest(r, t)
}
