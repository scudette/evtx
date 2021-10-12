// This file tests the WindowsMessageResolver to make sure we can
// properly extract messages from the registry/files as we resolve the
// event.

package evtx

import (
	"os/exec"
	"runtime"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/suite"
)

type EVTXTestSuite struct {
	suite.Suite
	binary string
}

func (self *EVTXTestSuite) SetupTest() {
	self.binary = "./dumpevtx"
	if runtime.GOOS == "windows" {
		self.binary += ".exe"
	}
}

func (self *EVTXTestSuite) TestCollector() {
	cmdline := []string{
		"parse", "--event_id", "4624",
		"--number", "1", "testdata/Security.evtx",
	}
	cmd := exec.Command(self.binary, cmdline...)
	out, err := cmd.CombinedOutput()
	assert.NoError(self.T(), err)

	goldie.Assert(self.T(), "Event4624_"+runtime.GOOS, out)
}

func TestEvtx(t *testing.T) {
	suite.Run(t, &EVTXTestSuite{})
}
