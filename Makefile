include $(GOROOT)/src/Make.inc

TARG=goirc
GOFILES=\
	src/goirc.go\
	src/messages/*.go

include $(GOROOT)/src/Make.cmd
