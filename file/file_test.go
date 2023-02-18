package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsFileExist(t *testing.T) {
	os.Mkdir("./tmp/", 0644)
	defer os.RemoveAll("./tmp/")
	os.Create("./tmp/1.log")
	assert.Equal(t, false, IsFileExist("./tmp/2.log"))
	assert.Equal(t, false, IsFileExist("./tmp/1.log/"))
	assert.Equal(t, true, IsFileExist("./tmp/1.log"))
}
