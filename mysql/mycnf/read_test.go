package mycnf

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader(`
[client]
user=root1
password=root2
`)
	cc, err := parse(reader)
	assert.NoError(err)
	assert.Len(cc, 1)
	assert.Equal("", cc[0].ConnectionName)
	assert.Equal("root1", cc[0].User)
	assert.Equal("root2", cc[0].Passwd)
	assert.Equal("", cc[0].DBName)
	assert.Equal("localhost:3306", cc[0].Addr)

	reader = strings.NewReader(`
[client]
user=root1
password=root2
database=test
port=4000
`)
	cc, err = parse(reader)
	assert.NoError(err)
	assert.Len(cc, 1)
	assert.Equal("", cc[0].ConnectionName)
	assert.Equal("root1", cc[0].User)
	assert.Equal("root2", cc[0].Passwd)
	assert.Equal("test", cc[0].DBName)
	assert.Equal("localhost:4000", cc[0].Addr)

	reader = strings.NewReader(`
[clientfoo]
user=root1
password=root2
database=test
port=4000
host=mysql


[clientbar]
user=second
password=pwd
port=4002
host=mysql2
`)
	cc, err = parse(reader)
	assert.NoError(err)
	assert.Len(cc, 2)
	assert.Equal("foo", cc[0].ConnectionName)
	assert.Equal("root1", cc[0].User)
	assert.Equal("root2", cc[0].Passwd)
	assert.Equal("test", cc[0].DBName)
	assert.Equal("mysql:4000", cc[0].Addr)
	assert.Equal("bar", cc[1].ConnectionName)
	assert.Equal("second", cc[1].User)
	assert.Equal("pwd", cc[1].Passwd)
	assert.Equal("", cc[1].DBName)
	assert.Equal("mysql2:4002", cc[1].Addr)
}
