package cache

import (
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"

	json "github.com/bytedance/sonic"
	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/constants"
)

// DecompressStruct decompresses GZIP compressed data and unmarshal it to the target struct. Note: target struct must be a pointer.
func DecompressStruct(ctx context.Context, data []byte, targetStruct any, compressionLibrary constants.CompressionLibraryType) error {
	var err error
	switch compressionLibrary {
	case constants.GzipCompressionType:
		b, err := gzipDecompression(data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, targetStruct)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal for struct decompression")
		}
	case constants.SnappyCompressionType:
		b, err := snappyDecompression(data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, targetStruct)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal for struct decompression")
		}
	default:
		err = json.Unmarshal(data, targetStruct)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal for struct decompression")
		}
	}

	return nil
}

func gzipDecompression(data []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(data)

	reader, err := gzip.NewReader(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new gzip reader")
	}

	decompressedDat, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "read all failed")
	}

	err = reader.Close()
	if err != nil {
		return nil, errors.Wrap(err, "unable to close read all")
	}

	return decompressedDat, nil
}

func snappyDecompression(data []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(data)

	reader := snappy.NewReader(buffer)

	decompressedDat, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "read all failed")
	}

	return decompressedDat, nil
}
