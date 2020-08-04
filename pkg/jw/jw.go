/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package jw

import (
	"time"

	"github.com/arnumina/failure"
)

type (
	// Priority AFAIRE.
	Priority int
)

const (
	// None AFAIRE.
	None Priority = 0
	// Low AFAIRE.
	Low = 20
	// Medium AFAIRE.
	Medium = 50
	// High AFAIRE.
	High = 80
	// Critical AFAIRE.
	Critical = 100
)

type (
	// Status AFAIRE.
	Status string
)

const (
	// Todo AFAIRE.
	Todo Status = "todo"
	// Running AFAIRE.
	Running = "running"
	// Pending AFAIRE.
	Pending = "pending"
	// Succeeded AFAIRE.
	Succeeded = "succeeded"
	// Failed AFAIRE.
	Failed = "failed"
	// Aborted AFAIRE.
	Aborted = "aborted"
)

type (
	// Job AFAIRE.
	Job struct {
		ID        string
		Name      string
		Plugin    string
		Type      string
		Origin    string
		Priority  Priority
		Key       *string
		Emails    *string
		Data      map[string]interface{}
		Status    Status
		Attempt   int
		WfID      *string
		WfFailed  *bool
		Reference time.Time
		CreatedAt time.Time
		RunAfter  time.Time
		Result    *string
		NextStep  *string
	}

	// Step AFAIRE.
	Step struct {
		Plugin string
		Type   string
		Config map[string]interface{}
		Next   map[string]interface{}
	}

	// Workflow AFAIRE.
	Workflow struct {
		ID         string
		Name       string
		Title      string
		Origin     string
		Priority   Priority
		FirstStep  string
		Steps      map[string]*Step
		Emails     *string
		Data       map[string]interface{}
		Status     Status
		CreatedAt  time.Time
		FinishedAt *time.Time
	}

	// Result AFAIRE.
	Result struct {
		Failure *failure.Failure
		Delay   time.Duration
	}

	// Runner AFAIRE.
	Runner interface {
		Run() *Result
	}
)

// NewJob AFAIRE.
func NewJob(id, n, p, t, o string, pr Priority, k, e *string) *Job {
	return &Job{
		ID:        id,
		Name:      n,
		Plugin:    p,
		Type:      t,
		Origin:    o,
		Priority:  pr,
		Key:       k,
		Emails:    e,
		Data:      make(map[string]interface{}),
		Status:    Todo,
		Reference: time.Now(),
		CreatedAt: time.Now(),
		RunAfter:  time.Now(),
	}
}

// SetResult AFAIRE.
func (j *Job) SetResult(s string) {
	*j.Result = s
}

// SetNextStep AFAIRE.
func (j *Job) SetNextStep(s string) {
	*j.NextStep = s
}

// NewWorkflow AFAIRE.
func NewWorkflow(id, n, t, o string, pr Priority, fs string, s map[string]*Step, e *string) *Workflow {
	return &Workflow{
		ID:        id,
		Name:      n,
		Title:     t,
		Origin:    o,
		Priority:  pr,
		FirstStep: fs,
		Steps:     s,
		Emails:    e,
		Data:      make(map[string]interface{}),
		Status:    Running,
		CreatedAt: time.Now(),
	}
}

// NewResult AFAIRE.
func NewResult(failure *failure.Failure, delay time.Duration) *Result {
	return &Result{
		Failure: failure,
		Delay:   delay,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
