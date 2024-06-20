package modes

type Mode int
type PID int

const (
	ModeCurrent            Mode = 0x01
	ModeFreezeFrame             = 0x02
	ModeStoredDTC               = 0x03
	ModeClearDTC                = 0x04
	ModeTestOxygen              = 0x05
	ModeTestOthers              = 0x06
	MModePendingDTC             = 0x07
	ModeControlOperation        = 0x08
	ModeVehicleInformation      = 0x09
	ModePermanentDTC            = 0x0A
)
