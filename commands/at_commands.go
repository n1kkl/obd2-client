package commands

// reference: https://www.elmelectronics.com/ELM327/AT_Commands.pdf

type AtCommand string

const (
	AtWarmStart         AtCommand = "ATWS"
	AtResetAll                    = "ATZ"
	AtEchoOff                     = "ATE0"
	AtEchoOn                      = "ATE1"
	AtLineFeedsOff                = "ATL0"
	AtLineFeedsOn                 = "ATL1"
	AtHeadersOff                  = "ATH0"
	AtHeadersOn                   = "ATH1"
	AtIdentify                    = "ATI"
	AtDeviceDescription           = "AT@1"
	AtDeviceIdentifier            = "AT@2"
	AtSetProtocol                 = "ATSP"
	AtDescribeProtocol            = "ATDP"
)
