# Goroutine concurrency queue process
## Builds
tba
## Example
```
  import "github.com/muhfajar/queue"
  ...
	ALLOC := 6
	q := NewQueue(&Worker{
		Thread: 3,
		Alloc:  ALLOC,
		Set: Callback{
			TaskDone: func(result interface{}) {
				rs := result.(string)
				fmt.Println(rs)
			},
			QueueDone: func() {
				fmt.Println("done")
			},
		},
	})

	strings := []string{"hello", "world", "fajar", "muhfajar", "queue", "github"}
	for i, s := range strings {
		index, value := i, s
		task := func() interface{} {
			time.Sleep(time.Duration(1) * time.Second)
			return fmt.Sprintf("pid: %d :: %s", index, value)
		}

		q.Append(task)
	}

	q.Start()
```
