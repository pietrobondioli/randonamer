package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// FileExists checks if a file exists, regardless of its extension.
func FileExists(path string) bool {
	matches, err := filepath.Glob(path + ".*")
	DebugLog("Checking if file exists with pattern: %s.*, found matches: %v, error: %v", path, matches, err)
	if err != nil {
		return false
	}
	return len(matches) > 0
}

var (
	ErrNoMatch  = errors.New("prefixpath: no matches")
	ErrMultiple = errors.New("prefixpath: more than one match")
)

func ReadFirst(prefix string) ([]byte, string, error) {
	DebugLog("Reading first file with prefix: %s", prefix)
	paths, err := filepath.Glob(prefix + ".*")
	if err != nil {
		return nil, "", err
	}
	if len(paths) == 0 {
		return nil, "", ErrNoMatch
	}
	if len(paths) > 1 {
		return nil, "", ErrMultiple
	}
	content, err := os.ReadFile(paths[0])
	if err != nil {
		return nil, "", err
	}
	DebugLog("Read file: %s, extension: %s", paths[0], filepath.Ext(paths[0]))
	return content, filepath.Ext(paths[0]), nil
}

func AgnosticUnmarshall[T any](path string, e *T) error {
	DebugLog("Unmarshalling file in path: %s", path)
	content, ext, err := ReadFirst(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, e)
	case ".json":
		err = json.Unmarshal(content, e)
	default:
		return errors.New("unsupported file extension")
	}

	if err != nil {
		return fmt.Errorf("failed to unmarshal file: %w", err)
	}

	DebugLog("Unmarshalled file: %s, result: %+v", path, e)
	return nil
}

func CopyDir(src, dest string) error {
	DebugLog("Copying directory from %s to %s", src, dest)
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(destPath, entry.Type()); err != nil {
				return err
			}
			if err := CopyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dest string) error {
	DebugLog("Copying file from %s to %s", src, dest)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
