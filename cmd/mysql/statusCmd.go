package mysql

import (
	"fmt"
	"github.com/mono83/atlas/impl/mysql/exec"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

type status struct {
	BaseDir       string
	DataDir       string
	Bind          string
	CharsetServer string
	Version       string
	VersionInnoDB string

	BufferPoolChunk uint64
	BufferPoolSize  uint64
	BufferPoolCount int

	Running []exec.Process
}

func (s status) getSlow() []exec.Process {
	response := []exec.Process{}
	for _, p := range s.Running {
		if p.Elapsed.Seconds() > 0.1 && "Sleep" != p.Command {
			response = append(response, p)
		}
	}

	return response
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"st"},
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := getConnection(cmd)
		if err != nil {
			return err
		}

		var status status

		// Reading global variables
		rows, err := db.Query("SHOW GLOBAL VARIABLES")
		if err != nil {
			return err
		}

		var name, value string
		for rows.Next() {
			if err = rows.Scan(&name, &value); err != nil {
				return err
			}

			switch strings.ToLower(name) {
			case "basedir":
				status.BaseDir = value
			case "datadir":
				status.DataDir = value
			case "bind_address":
				status.Bind = value
			case "character_set_server":
				status.CharsetServer = value
			case "innodb_buffer_pool_chunk_size":
				status.BufferPoolChunk, _ = strconv.ParseUint(value, 10, 64)
			case "innodb_buffer_pool_size":
				status.BufferPoolSize, _ = strconv.ParseUint(value, 10, 64)
			case "innodb_buffer_pool_instances":
				status.BufferPoolCount, _ = strconv.Atoi(value)
			case "version":
				status.Version = value
			case "innodb_version":
				status.VersionInnoDB = value
			}
		}

		rows.Close()

		// Reading slow queries
		status.Running, err = exec.ShowFullProcesslist(db)
		if err != nil {
			return err
		}

		// Printing status
		fmt.Println("Version:      ", status.Version, "(InnoDB:", status.VersionInnoDB+")")
		fmt.Println("Base dir:     ", status.BaseDir)
		fmt.Println("Data dir:     ", status.BaseDir)
		fmt.Println("Bind address: ", status.Bind, " Charset: ", status.CharsetServer)
		fmt.Printf(
			"Buffer pool:   %d chunks, %.1fMb each, %.1fMb total\n",
			status.BufferPoolCount,
			float32(status.BufferPoolChunk)/1024./1024.,
			float32(status.BufferPoolSize)/1024./1024.,
		)

		// Printing slow queries
		if slow := status.getSlow(); len(slow) > 0 {
			fmt.Println()
			fmt.Println("Slow queries:")
			for _, q := range slow {
				fmt.Println(q.ID, q.Elapsed, q.Info)
			}
		} else {
			fmt.Println("No slow queries")
		}

		return nil
	},
}
