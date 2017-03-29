package reader

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func newTmpFile(data string) (tmpFile *os.File, err error) {
	tmpFile, err = ioutil.TempFile("", "")
	if err != nil {
		return
	}

	_, err = io.WriteString(tmpFile, data)
	if err != nil {
		return
	}

	_, err = tmpFile.Seek(0, os.SEEK_SET)
	if err != nil {
		return
	}

	return
}

func TestGetLocationPoints(t *testing.T) {
	// Mock the data file.
	csvData1 := `"id","lat","lng"
1,1.12,3.45
2,6.78,9.10`

	tmpFile1, err := newTmpFile(csvData1)
	if err != nil {
		t.Error(err)
	}

	r1 := &Reader{
		file: tmpFile1,
	}

	// Retrieve and validate the data.
	var count int
	err = r1.GetLocationPoints(func(id int, lat, lng float64) error {
		switch count {
		case 0:
			if id != 1 || lat != 1.12 || lng != 3.45 {
				t.Errorf("a record of if %d contains invalid data - lat: %f, lng: %f", id, lat, lng)
			}
		case 1:
			if id != 2 || lat != 6.78 || lng != 9.10 {
				t.Errorf("a record of if %d contains invalid data - lat: %f, lng: %f", id, lat, lng)
			}
		default:
			t.Error("received parsed data is bigger than expected")
		}
		count++
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("failed to loop through all the records of a valid file", count)
	}

	tmpFile1.Close()
	os.Remove(tmpFile1.Name())

	// Mock invalid file contents (wrong type).
	csvData2 := `"id","lat","lng"
1,abc,3.45`

	tmpFile2, err := newTmpFile(csvData2)
	if err != nil {
		t.Error(err)
	}

	r2 := &Reader{
		file: tmpFile2,
	}

	err = r2.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if err == nil {
		t.Error("failed to determine invalid latitude record in the CSV file")
	}

	tmpFile2.Close()
	os.Remove(tmpFile2.Name())

	// Mock invalid file contents (wrong type).
	csvData3 := `"id","lat","lng"
1,3.45,abc`

	tmpFile3, err := newTmpFile(csvData3)
	if err != nil {
		t.Error(err)
	}

	r3 := &Reader{
		file: tmpFile3,
	}

	err = r3.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if err == nil {
		t.Error("failed to determine invalid longitude record in the CSV file")
	}

	tmpFile3.Close()
	os.Remove(tmpFile3.Name())

	// Mock invalid file contents (wrong type).
	csvData4 := `"id","lat","lng"
abc,3.45,3.45`

	tmpFile4, err := newTmpFile(csvData4)
	if err != nil {
		t.Error(err)
	}

	r4 := &Reader{
		file: tmpFile4,
	}

	err = r4.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if err == nil {
		t.Error("failed to determine invalid id record in the CSV file")
	}

	tmpFile4.Close()
	os.Remove(tmpFile4.Name())

	// Mock invalid file contents (insufficient amount of columns).
	csvData5 := `"id","lat","lng"
1,1.12`

	tmpFile5, err := newTmpFile(csvData5)
	if err != nil {
		t.Error(err)
	}

	r5 := &Reader{
		file: tmpFile5,
	}

	parsedInvalidData2Err := r5.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if parsedInvalidData2Err == nil {
		t.Error("failed to determine insufficient amount of columns in the CSV file")
	}

	tmpFile5.Close()
	os.Remove(tmpFile5.Name())
}
