package lib

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/pijusn/files-rename/lib/scheduling"
)

type Job struct {
	sourceName string
	targetName string
}

func Run(config *Config) error {
	err := config.Validate()
	if err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	filenames, err := allFiles(config.Directory)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}
	slices.Sort(filenames)

	tasks := plan(filenames, config.Name)
	for _, task := range tasks {
		pathOld := path.Join(config.Directory, task.NameSource)
		pathNew := path.Join(config.Directory, task.NameTarget)
		err := os.Rename(pathOld, pathNew)
		if err != nil {
			return fmt.Errorf("failed to rename file %s to %s: %w", pathOld, pathNew, err)
		}
	}

	return nil
}

func allFiles(name string) ([]string, error) {
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		names = append(names, entry.Name())
	}
	return names, nil
}

func plan(originalFilenames []string, namePattern string) []scheduling.Task {
	tasks := makeTasks(originalFilenames, namePattern)
	scheduling.Optimize(tasks)
	return tasks
}

func makeTasks(originalFilenames []string, namePattern string) []scheduling.Task {
	tasks := make([]scheduling.Task, 0, len(originalFilenames))
	for i, nameOld := range originalFilenames {
		nameNew := computeFilename(nameOld, namePattern, i+1)
		if nameOld == nameNew {
			continue
		}
		tasks = append(tasks, scheduling.Task{
			NameSource: nameOld,
			NameTarget: nameNew,
		})
	}
	return tasks
}

func computeFilename(originalName string, newNamePattern string, id int) string {
	extension := strings.ToLower(path.Ext(originalName))
	base := fmt.Sprintf(newNamePattern, id)
	return base + extension
}
