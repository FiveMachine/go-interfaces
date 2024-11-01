package builders

import (
	"io"

	ci "github.com/taubyte/go-simple-container"
	"github.com/taubyte/go-specs/builders"
)

type Builder interface {
	// Build will build the given working directory as per builder configuration and returns Output
	Build(...ci.ContainerOption) (Output, error)
	// Close cleans up the builder
	Close() error
	// Config returns the builder configuration
	Config() *builders.Config
	// Wd returns the builder working directory
	Wd() builders.Dir
	// Tarball returns the tarball of the image used to build, if any
	Tarball() []byte
}

type Output interface {
	// Compress takes a supported CompressionMethod, compress the files built by the Builder, and returns the ReadSeekCloser of the compressed file
	Compress(CompressionMethod) (io.ReadSeekCloser, error)
	// Logs returns the ReadCloser of the build logs
	Logs() Logs
	// OutDir returns the output directory of the built files, pre zip or compress
	OutDir() string
	// Closes logs
	Close() error
}

// Logs is an interface defining methods for handling log operations.
type Logs interface {
	// Embeds io.ReadCloser to allow reading and closing log data streams.
	io.ReadCloser

	// CopyTo copies log data to a given destination (dst), which implements io.Writer.
	// Returns the number of bytes copied and an error if any occurs.
	CopyTo(dst io.Writer) (int64, error)

	// CopyFrom copies data from a source (src), which implements io.Reader, to the log.
	// Returns the number of bytes copied and an error if any occurs.
	CopyFrom(src io.Reader) (int64, error)

	// FormatErr formats an error message according to the specified format and arguments.
	// Returns an error containing the formatted message.
	FormatErr(format string, args ...any) error
}


// CompressionMethod defines the method used to compress build Output of a Builder
type CompressionMethod uint32
