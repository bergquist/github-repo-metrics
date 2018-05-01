package archive

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsingDatetime(t *testing.T) {
	tc := []struct {
		Date     string
		ArchFile *ArchiveFile
	}{
		{Date: "2018-01-01T20:38:03.000Z", ArchFile: NewArchiveFile(2018, 1, 1, 20)},
		{Date: "2015-01-01T20:38:03.000Z", ArchFile: NewArchiveFile(2015, 1, 1, 20)},
	}

	for _, test := range tc {
		date, err := time.Parse("2006-01-02T15:04:05.000Z", test.Date)
		if err != nil {
			t.Fatalf("failed to parse date")
		}

		arch := GetArchDateFrom(date)
		assert.Equal(t, test.ArchFile, arch, "not same arch")
	}
}
