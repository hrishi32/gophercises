package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task is struct for key vaues of buckets
type Task struct {
	Key   int
	Value string
}

// Init initializes the bolt database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask creates a task with a passed parameter as string. Returns id, nil on
// successful creation and returns -1, error on error occured
func CreateTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		id64, _ := bucket.NextSequence()
		id = int(id64)

		return bucket.Put(itob(id), []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func ListAllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		// get the bucket
		bucket := tx.Bucket(taskBucket)
		cursor := bucket.Cursor()

		//loop over elements in bucket using cursor

		for key, val := cursor.First(); key != nil; key, val = cursor.Next() {
			tasks = append(tasks, Task{Key: btoi(key), Value: string(val)})
		}

		return nil
	})

	return tasks, err
}

//DeleteTask deletes the task with assocoated key
func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		// get the bucket
		bucket := tx.Bucket(taskBucket)

		return bucket.Delete(itob(key))
	})

	return err
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))

	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
