package cache

import (
	"bytes"
	"compress/gzip"
	"context"

	json "github.com/bytedance/sonic"
	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/constants"
)

const (
	defaultCompressionLevel = gzip.DefaultCompression
)

// CompressStruct converts a struct to JSON and compresses it according to chosen compression library.
func CompressStruct(ctx context.Context, s any, compressionLibrary constants.CompressionLibraryType) ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal struct")
	}
	switch compressionLibrary {
	case constants.GzipCompressionType:
		return gzipCompression(b)
	case constants.SnappyCompressionType:
		return snappyCompression(b)
	default:
		return b, nil
	}
}

func gzipCompression(data []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := gzip.NewWriterLevel(&buffer, defaultCompressionLevel)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new gzip writer")
	}

	_, err = writer.Write(data)
	if err != nil {
		return nil, errors.Wrap(err, "cannot write to gzip")
	}

	err = writer.Close()
	if err != nil {
		return nil, errors.Wrap(err, "cannot close writer")
	}

	return buffer.Bytes(), nil
}

func snappyCompression(data []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer := snappy.NewBufferedWriter(&buffer)

	_, err := writer.Write(data)

	if err != nil {
		return nil, errors.Wrap(err, "cannot write to snappy")
	}

	err = writer.Close()
	if err != nil {
		return nil, errors.Wrap(err, "cannot close writer")
	}

	return buffer.Bytes(), nil
}
