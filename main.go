package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)


const serverURL = "http://localhost:8080"

type Job struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

func submitJob(jobType, payload, priority string) (Job, error) {
	requestBody  := map[string]interface{} {
		"type" : jobType,
		"payload": payload,
		"priority": priority,
		"max_attempts" : 3,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return Job{}, err
	}

	resp, err := http.Post(serverURL+"/jobs", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return Job{}, err
	}
	defer resp.Body.Close()

	var job Job
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Job{}, err
	}
	if err := json.Unmarshal(body, &job); err != nil {
		return Job{}, err
	}
	return job, nil
}

func getJobStatus(id string) (Job, error) {
	resp, err := http.Get(serverURL + "/jobs/" + id)
	
	if err != nil {
		return Job{}, err
	}
	defer resp.Body.Close()
	var job Job
	body, err := io.ReadAll(resp.Body)
	fmt.Println("서버 응답 원문:", string(body)) 
	if err != nil {
		return Job{}, err
	}
	if err := json.Unmarshal(body, &job); err != nil {
		return Job{}, err
	}
	return job, nil
}

func loadTest(count int) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, err := submitJob("email:send", "test@example.com", "high")
			if err != nil {
				fmt.Print("에러")
			}
		} (i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%d개 작업 등록에 걸린 시간: %v\n", count, elapsed)
}

func main() {
	// job, err := submitJob("email:send", "test@example.com", "high")
	// if err != nil {
	// 	fmt.Println("에러:", err)
	// 	return
	// }
	// fmt.Printf("등록됨: %+v\n", job)

	// job, err = getJobStatus(job.ID) 
	// if err != nil {
	// 	fmt.Println("에러:", err)
	// 	return
	// }
	// fmt.Printf("job 상태: %+v\n", job.Status)

	loadTest(50)


}