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

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

var (
	ErrNoMatch  = errors.New("prefixpath: no matches")
	ErrMultiple = errors.New("prefixpath: more than one match")
)

func ReadFirst(prefix string) ([]byte, string, error) {
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
	return content, filepath.Ext(paths[0]), nil
}

func AgnosticUnmarshall[T any](path, file string, e *T) error {
	content, ext, err := ReadFirst(filepath.Join(path, file))
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

	return nil
}

func CopyDir(src, dest string) error {
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
