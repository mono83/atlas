package mysql

import (
	"fmt"
	"github.com/mono83/atlas/mysql/preset"
	"github.com/mono83/atlas/sql"
	"github.com/mono83/table"
	"github.com/mono83/table/cells"
	"github.com/mono83/table/reflect"
	"github.com/spf13/cobra"
	"sort"
	"strconv"
	"strings"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"st", "stat", "stats"},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var conn sql.Modifier

		// Establishing connection
		conn, err = connect(cmd)
		if err != nil {
			return
		}

		// Reading data
		var varsList []struct {
			Name  string `db:"Variable_name"`
			Value string `db:"Value"`
		}
		if err = conn.Invoke(sql.StringStatement("SHOW GLOBAL VARIABLES"), &varsList); err != nil {
			return
		}
		// Building data struct
		var status struct {
			BaseDir       string
			DataDir       string
			Bind          string
			CharsetServer string
			Version       string
			VersionInnoDB string

			BufferPoolChunk uint64
			BufferPoolSize  uint64
			BufferPoolCount int
		}
		for _, v := range varsList {
			switch strings.ToLower(v.Name) {
			case "basedir":
				status.BaseDir = v.Value
			case "datadir":
				status.DataDir = v.Value
			case "bind_address":
				status.Bind = v.Value
			case "character_set_server":
				status.CharsetServer = v.Value
			case "innodb_buffer_pool_chunk_size":
				status.BufferPoolChunk, _ = strconv.ParseUint(v.Value, 10, 64)
			case "innodb_buffer_pool_size":
				status.BufferPoolSize, _ = strconv.ParseUint(v.Value, 10, 64)
			case "innodb_buffer_pool_instances":
				status.BufferPoolCount, _ = strconv.Atoi(v.Value)
			case "version":
				status.Version = v.Value
			case "innodb_version":
				status.VersionInnoDB = v.Value
			}
		}

		if t, err := reflect.StructToTable(status); err == nil {
			fmt.Println("General variables")
			table.PrintStandard(t)
		}

		// Reading processlist
		var pl preset.ProcessList
		if err = sql.SelectFill(conn, &pl); err != nil {
			return
		}
		pl.SortByTimeAsc()

		// Per user utilization
		puu := pl.GetUtilizationByUser()
		sort.Slice(puu, func(i, j int) bool {
			ik := puu[i].Host + puu[i].User
			jk := puu[j].Host + puu[j].User
			return ik < jk
		})
		fmt.Println()
		fmt.Println("Per-user utilization")
		table.PrintStandard(utilizationTable(puu))

		// Printing processlist
		fmt.Println()
		if pl.Len() > 0 {
			fmt.Println("Not IDLE processes")
			pl = pl.NotIDLE()
		} else {
			fmt.Println("Full processes list")
		}
		table.PrintStandard(processesTable(pl.All()))

		return nil
	},
}

type utilizationTable []preset.Utilization

func (utilizationTable) Headers() []string {
	return []string{"User", "Host", "Procs", "Time"}
}

func (u utilizationTable) EachRow(f func(...table.Cell)) {
	for _, row := range u {
		f(
			cells.String(row.User),
			cells.String(row.Host),
			cells.Int(row.Processes),
			cells.Duration(row.Total),
		)
	}
}

type processesTable []preset.Process

func (processesTable) Headers() []string {
	return []string{
		"Id",
		"Who",
		"Database",
		"State",
		"Time",
		"Sent",
		"Command",
	}
}

func (p processesTable) EachRow(f func(...table.Cell)) {
	for _, row := range p {
		var state table.Cell = cells.Empty{}
		row.State.IfNotEmpty(func(s string) {
			state = cells.String(s)
		})
		var db table.Cell = cells.Empty{}
		row.Database.IfNotEmpty(func(s string) {
			db = cells.String(s)
		})
		var who table.Cell
		if row.Database.IsNotEmpty() {
			who = cells.Sprintf("%s@%s/%s", row.User, row.Host, row.Database.OrElse(""))
		} else {
			who = cells.Sprintf("%s@%s", row.User, row.Host)
		}
		if row.User == "root" {
			who = cells.ColoredRedHi(who)
		}

		f(
			cells.Int64(row.ID),
			who,
			db,
			state,
			cells.Duration(row.Elapsed.Duration),
			cells.Int(row.RowsSent),
			cells.String(row.Command),
		)
	}
}
