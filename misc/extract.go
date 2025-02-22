package miscutils

import (
	"archive/zip"
	"fmt"
	"io"
)

func ExtractPayloadFromZip(
	zipfilePath string,
	filename string,
) ([]byte, error) {
	zr, err := zip.OpenReader(zipfilePath)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	for _, f := range zr.File {
		if f.Name != filename {
			continue
		}
		zf, ferr := f.Open()
		if ferr != nil {
			return nil, ferr
		}
		defer zf.Close()
		return io.ReadAll(zf)
	}
	return nil, fmt.Errorf("file not found in zip")
}
