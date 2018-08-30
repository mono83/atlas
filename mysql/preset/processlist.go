package preset

import (
	"github.com/mono83/atlas/sql"
	"github.com/mono83/atlas/types"
	"time"
)

// Process contains information about running queries
type Process struct {
	ID       int64                 `db:"Id"`
	User     string                `db:"User"`
	Host     string                `db:"Host"`
	Database types.OptString       `db:"db"`
	Command  string                `db:"Command"`
	Elapsed  types.DurationSeconds `db:"Time"`
	State    types.OptString       `db:"State"`
	Info     types.OptString       `db:"Info"`
}

// ProcessList contains list of processes
type ProcessList []Process

// Len return amount of items
func (p ProcessList) Len() int {
	return len(p)
}

// All returns full list of process information
func (p ProcessList) All() []Process {
	return p
}

// WithElapsed returns only that processes, that lasts more that provided duration (inclusive)
func (p ProcessList) WithElapsed(d time.Duration) (list []Process) {
	for _, process := range p {
		if process.Elapsed.Duration >= d {
			list = append(list, process)
		}
	}
	return
}

// GetStatement returns statement, that could be used to obtain data
// Also it is sql.Filler interface implementation
func (p *ProcessList) GetStatement(args ...interface{}) (sql.Statement, error) {
	return sql.StringStatement("SHOW FULL PROCESSLIST"), nil
}
