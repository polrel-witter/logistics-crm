package main

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
	ID            int
	Name          string
	Domain        string
	CgCode        *string
	Lead          *string
	Note          *string
	Revenue       *int
	Locations     []string
	EmployeeCount *int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Port struct {
	ID   int
	Name string
	Kind PortType
}

type TradeData struct {
	ID             int
	CompanyID      int
	PortID         int
	Direction      DirectionType
	Volume         int
	Period         Period
	TopCommodities []string
	Trend          VolumeTrend
	CreatedAt      time.Time
	UpdatedAt      time.Time
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
