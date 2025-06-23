package models

import "time"

// Enums as string types with constants
type PortType string

const (
	PortTypeOcean PortType = "ocean"
	PortTypeRail  PortType = "rail"
)

type DirectionType string

const (
	DirectionTypeImport DirectionType = "import"
	DirectionTypeExport DirectionType = "export"
)

type VolumeTrend string

const (
	VolumeTrendIncreasing VolumeTrend = "increasing"
	VolumeTrendDecreasing VolumeTrend = "decreasing"
	VolumeTrendStable     VolumeTrend = "stable"
)

type Period string

const (
	PeriodMonthly   Period = "monthly"
	PeriodQuarterly Period = "quarterly"
	PeriodYearly    Period = "yearly"
)

// Base structs (database entities)
type Company struct {
	ID        int
	Domain    string
	Name      string
	CgCode    *string
	Note      *string
	Industry  *string
	Revenue   *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Port struct {
	ID   int
	Name string
	Kind PortType
}

type TradeData struct {
	ID        int
	CompanyID int
	PortID    int
	Direction DirectionType
	Volume    int
	Period    Period
	Trend     VolumeTrend
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contact struct {
	ID        int
	CompanyID int
	FirstName string
	LastName  string
	Email     *string
	Phone     *string
	Title     *string
	CreatedAt time.Time
}

// Methods
func (c *Company) ScanFields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.Domain,
		&c.Name,
		&c.CgCode,
		&c.Note,
		&c.Industry,
		&c.Revenue,
		&c.CreatedAt,
		&c.UpdatedAt,
	}
}
