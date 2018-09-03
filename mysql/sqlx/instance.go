package sqlx

import (
	"errors"

	"github.com/mono83/xray/args"
	"github.com/mono83/xray/std"

	"github.com/mono83/xray"

	"github.com/jmoiron/sqlx"
	"github.com/mono83/atlas/sql"
)

// Instance is wrapper over SQLx
type Instance struct {
	*sqlx.DB
}

var foo sql.ModifierX = &Instance{}

// Invoke is sql.Invoker interface implementation
func (i *Instance) Invoke(stmt sql.Statement, target interface{}) error {
	return i.InvokeX(xray.ROOT, stmt, target)
}

// InvokeX is sql.InvokerX interface implementation
func (i *Instance) InvokeX(ray xray.Ray, stmt sql.Statement, target interface{}) error {
	ray = ray.WithLogger("sqlx").With(args.SQL(stmt.GetSQL()), args.Type("invoke"))
	timer := std.Timer(ray, "mysql")
	defer timer.Stop()

	if i.DB == nil {
		return errors.New("DB not configured")
	}

	err := i.DB.Select(target, stmt.GetSQL(), stmt.GetPlaceholders()...)
	if err == nil {
		ray.Inc("mysql.success")
	} else {
		ray.Inc("mysql.fail")
	}

	return err
}

// Modify is sql.Modifier interface implementation
func (i *Instance) Modify(stmt sql.Statement) (int64, error) {
	return i.ModifyX(xray.ROOT, stmt)
}

// ModifyX is sql.ModifierX interface implementation
func (i *Instance) ModifyX(ray xray.Ray, stmt sql.Statement) (int64, error) {
	ray = ray.WithLogger("sqlx").With(args.SQL(stmt.GetSQL()), args.Type("invoke"))
	timer := std.Timer(ray, "mysql")
	defer timer.Stop()

	if i.DB == nil {
		return 0, errors.New("DB not configured")
	}
	if stmt == nil {
		return 0, errors.New("empty statement")
	}

	res, err := i.DB.Exec(stmt.GetSQL(), stmt.GetPlaceholders()...)
	if err != nil {
		ray.Inc("mysql.fail")
		return 0, err
	}
	ray.Inc("mysql.success")

	if id, e := res.LastInsertId(); e == nil && id != 0 {
		return id, nil
	}

	rc, _ := res.RowsAffected()
	return rc, nil
}
