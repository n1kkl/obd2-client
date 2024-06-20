package modes

type ModeCurrentPID PID

const (
	PidEngineCoolantTemp ModeCurrentPID = 0x05
	PidEngineSpeed       ModeCurrentPID = 0x0C
	PidVehicleSpeed      ModeCurrentPID = 0x0D
	PidMassAirflow       ModeCurrentPID = 0x10
	PidThrottlePosition  ModeCurrentPID = 0x11
)
