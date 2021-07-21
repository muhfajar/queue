package queue

import (
	"testing"
	"time"
)

type data struct {
	index int
	value string
}

func TestWorker_Append(t *testing.T) {
	NewQueue(&Worker{}).Append(func() interface{} {
		return true
	})
}

func TestWorker_Start(t *testing.T) {
	q := NewQueue(&Worker{
		Set: Callback{
			TaskDone: func(result interface{}) {
				rs := result.(bool)
				t.Logf("Callback: ok :: %v", rs)
			},
		},
	})
	q.Append(func() interface{} {
		return true
	})

	q.Start()
}

func TestNewQueue(t *testing.T) {
	TOTAL := 0
	ALLOC := 6
	q := NewQueue(&Worker{
		Thread: 3,
		Alloc:  ALLOC,
		Set: Callback{
			TaskDone: func(result interface{}) {
				rs := result.(data)
				TOTAL++
				t.Log(rs)
			},
			QueueDone: func() {
				t.Log("done")
			},
		},
	})

	strings := []string{"hello", "world", "fajar", "muhfajar", "queue", "github"}
	for i, s := range strings {
		index, value := i, s
		task := func() interface{} {
			time.Sleep(time.Duration(1) * time.Second)
			return data{
				index: index,
				value: value,
			}
		}

		q.Append(task)
	}

	q.Start()

	if TOTAL != ALLOC {
		t.Errorf("Queue size = %d; want 1", TOTAL)
	}
}
