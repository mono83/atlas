package mysql

func escape(key string) string {
	return "`" + key + "`"
}
