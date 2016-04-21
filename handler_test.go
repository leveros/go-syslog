package syslog

import (
	"github.com/leveros/syslogparser"
	. "gopkg.in/check.v1"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestHandle(c *C) {
	logPart := syslogparser.LogParts{"tag": "foo"}

	channel := make(LogPartsChannel, 1)
	handler := NewChannelHandler(channel)
	handler.Handle(logPart, 10, nil)

	fromChan := <-channel
	c.Check(fromChan["tag"], Equals, logPart["tag"])
}
