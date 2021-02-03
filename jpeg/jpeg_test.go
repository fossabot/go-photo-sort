package jpeg_test

import (
	"testing"

	"github.com/sebfoucault/go-photo-sort/jpeg"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestGetExifData(t *testing.T) {

	data, _ := jpeg.GetExifData("../test-files/scotland-nicolas-boulesteix.jpg")

	then.AssertThat(t, data, is.Not(is.Nil()))
}

func TestGetExifDataWithNonExistingFile(t *testing.T) {

	data, err := jpeg.GetExifData("../test-files/does-not-exist.jpg")

	then.AssertThat(t, data, is.Nil())
	then.AssertThat(t, err, is.Not(is.Nil()))
}