package scheduling

func Optimize(tasks []Task) {
	pending := make(map[string]bool, len(tasks))
	for _, t := range tasks {
		pending[t.NameSource] = true
	}

	i := 0
	for i < len(tasks) {
		task := &tasks[i]
		if !pending[task.NameTarget] {
			// Can be renamed therefore can remain in its position.
			pending[task.NameSource] = false
			i++
		} else {
			// Name collision therefore must first rename the other one.
			j := i + 1 + indexByNameSource(tasks[i+1:], task.NameTarget)
			sliceShiftRight(tasks[i : j+1])
		}
	}
}

func sliceShiftRight[T any](items []T) {
	if len(items) < 2 {
		return
	}
	temp := items[len(items)-1]
	for j := len(items) - 1; j > 0; j-- {
		items[j] = items[j-1]
	}
	items[0] = temp
}

func indexByNameSource(items []Task, needle string) int {
	for i, v := range items {
		if v.NameSource == needle {
			return i
		}
	}
	return -1
}
