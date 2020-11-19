// +build !solution

package fs

import (
	"context"
	"dat320/lab8/grpcfs/proto"
	"errors"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// OpenRead is passed to Open to indicate a file should be opened as read-only
	OpenRead = iota
	// OpenWrite is passed to Open to indicate a file should be opened as write-only
	OpenWrite
	// OpenReadWrite is passed to Open to indicate a file should be opened as read-write
	OpenReadWrite
)

var (
	// ErrOpenDir occurs if the caller tries to open a directory (only files
	// can be opened since none of the File operations fit directories)
	ErrOpenDir = errors.New("cannot open a directory")
	// ErrOpenInvalidFlag occurs if the caller passes an unknown flag to Open
	ErrOpenInvalidFlag = errors.New("tried to open file with an unrecognized flag")
	errNotImplemented  = errors.New("this feature is not yet implemented")
)

// FileSystem acts as a pseudo-filesystem over gRPC. Files can be opened by the
// FS, which then interact with the filesystem of the remote server over gRPC
// for their supported operations.
type FileSystem struct {
	// gRPC client - communicates with the gRPC file system server
	client proto.FileSystemClient
	// TODO(student): add necessary fields here (if any)
}

// NewFileSystem returns a FileSystem with an initialized gRPC client connection
func NewFileSystem(addr string, opts ...grpc.DialOption) (*FileSystem, error) {
	var conn grpc.ClientConnInterface
	var err error

	if len(opts) == 0 {
		// if no options are provided, default to insecure connection
		conn, err = grpc.Dial(addr, grpc.WithInsecure())
	} else {
		// use the options provided by the caller
		conn, err = grpc.Dial(addr, opts...)
	}

	if err != nil {
		return nil, err
	}

	fs := &FileSystem{
		client: proto.NewFileSystemClient(conn),
		// TODO(student): initialize additional fields here (if any)
	}

	return fs, nil
}

// Open opens the file `name` at the remote server.
// `flag` must be one of `OpenRead`, `OpenWrite` or `OpenReadWrite`.
// Files opened with `OpenRead` must already exist.
// Calling `Open` with `OpenWrite` or `OpenReadWrite` creates the file on the
// remote server if it does not already exist.
// Calling `Open` with `OpenRead` on a non-existent file returns an error.
// Directories cannot be opened.
func (fs *FileSystem) Open(name string, flag int) (f *File, err error) {
	// TODO(student): implement this method
	return nil, errNotImplemented
}

// Remove removes the file/directory at the remote server
func (fs *FileSystem) Remove(name string) error {
	_, err := fs.client.Remove(context.Background(), &proto.RemoveRequest{Path: name})
	return translateStatusError(err)
}

// Mkdir makes a directory at the remote server
func (fs *FileSystem) Mkdir(path string, perm os.FileMode) error {
	// TODO(student): implement this method
	return errNotImplemented
}

// Lookup looks up the path at the remote server and returns any information it receives
func (fs *FileSystem) Lookup(path string) (isDir bool, size int, files []string, err error) {
	// TODO(student): implement this method
	return false, 0, nil, errNotImplemented
}

// translateStatusError translates status errors to errors from the `os` package
// by matching the codes to those defined by the server
func translateStatusError(err error) error {
	if _, isStatusError := status.FromError(err); isStatusError {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			err = os.ErrNotExist
		case codes.AlreadyExists:
			err = os.ErrExist
		case codes.FailedPrecondition:
			err = os.ErrClosed
		case codes.PermissionDenied:
			err = os.ErrPermission
		case codes.InvalidArgument:
			err = os.ErrInvalid
		default:
			// Workaround. For some reason these are sometimes not
			// matched at the server-side.
			if err != nil && strings.Contains(err.Error(), os.ErrNotExist.Error()) {
				err = os.ErrNotExist
			} else if (err != nil && strings.Contains(err.Error(), os.ErrExist.Error())) ||
				(err != nil && strings.Contains(err.Error(), "already exists")) {
				err = os.ErrExist
			}
		}
	}

	return err
}
