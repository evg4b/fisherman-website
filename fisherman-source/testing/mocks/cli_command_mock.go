package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i fisherman/internal.CliCommand -o ./testing/mocks/cli_command_mock.go

import (
	mm_internal "fisherman/internal"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// CliCommandMock implements internal.CliCommand
type CliCommandMock struct {
	t minimock.Tester

	funcDescription          func() (s1 string)
	inspectFuncDescription   func()
	afterDescriptionCounter  uint64
	beforeDescriptionCounter uint64
	DescriptionMock          mCliCommandMockDescription

	funcInit          func(args []string) (err error)
	inspectFuncInit   func(args []string)
	afterInitCounter  uint64
	beforeInitCounter uint64
	InitMock          mCliCommandMockInit

	funcName          func() (s1 string)
	inspectFuncName   func()
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mCliCommandMockName

	funcRun          func(ctx mm_internal.ExecutionContext) (err error)
	inspectFuncRun   func(ctx mm_internal.ExecutionContext)
	afterRunCounter  uint64
	beforeRunCounter uint64
	RunMock          mCliCommandMockRun
}

// NewCliCommandMock returns a mock for internal.CliCommand
func NewCliCommandMock(t minimock.Tester) *CliCommandMock {
	m := &CliCommandMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DescriptionMock = mCliCommandMockDescription{mock: m}

	m.InitMock = mCliCommandMockInit{mock: m}
	m.InitMock.callArgs = []*CliCommandMockInitParams{}

	m.NameMock = mCliCommandMockName{mock: m}

	m.RunMock = mCliCommandMockRun{mock: m}
	m.RunMock.callArgs = []*CliCommandMockRunParams{}

	return m
}

type mCliCommandMockDescription struct {
	mock               *CliCommandMock
	defaultExpectation *CliCommandMockDescriptionExpectation
	expectations       []*CliCommandMockDescriptionExpectation
}

// CliCommandMockDescriptionExpectation specifies expectation struct of the CliCommand.Description
type CliCommandMockDescriptionExpectation struct {
	mock *CliCommandMock

	results *CliCommandMockDescriptionResults
	Counter uint64
}

// CliCommandMockDescriptionResults contains results of the CliCommand.Description
type CliCommandMockDescriptionResults struct {
	s1 string
}

// Expect sets up expected params for CliCommand.Description
func (mmDescription *mCliCommandMockDescription) Expect() *mCliCommandMockDescription {
	if mmDescription.mock.funcDescription != nil {
		mmDescription.mock.t.Fatalf("CliCommandMock.Description mock is already set by Set")
	}

	if mmDescription.defaultExpectation == nil {
		mmDescription.defaultExpectation = &CliCommandMockDescriptionExpectation{}
	}

	return mmDescription
}

// Inspect accepts an inspector function that has same arguments as the CliCommand.Description
func (mmDescription *mCliCommandMockDescription) Inspect(f func()) *mCliCommandMockDescription {
	if mmDescription.mock.inspectFuncDescription != nil {
		mmDescription.mock.t.Fatalf("Inspect function is already set for CliCommandMock.Description")
	}

	mmDescription.mock.inspectFuncDescription = f

	return mmDescription
}

// Return sets up results that will be returned by CliCommand.Description
func (mmDescription *mCliCommandMockDescription) Return(s1 string) *CliCommandMock {
	if mmDescription.mock.funcDescription != nil {
		mmDescription.mock.t.Fatalf("CliCommandMock.Description mock is already set by Set")
	}

	if mmDescription.defaultExpectation == nil {
		mmDescription.defaultExpectation = &CliCommandMockDescriptionExpectation{mock: mmDescription.mock}
	}
	mmDescription.defaultExpectation.results = &CliCommandMockDescriptionResults{s1}
	return mmDescription.mock
}

//Set uses given function f to mock the CliCommand.Description method
func (mmDescription *mCliCommandMockDescription) Set(f func() (s1 string)) *CliCommandMock {
	if mmDescription.defaultExpectation != nil {
		mmDescription.mock.t.Fatalf("Default expectation is already set for the CliCommand.Description method")
	}

	if len(mmDescription.expectations) > 0 {
		mmDescription.mock.t.Fatalf("Some expectations are already set for the CliCommand.Description method")
	}

	mmDescription.mock.funcDescription = f
	return mmDescription.mock
}

// Description implements internal.CliCommand
func (mmDescription *CliCommandMock) Description() (s1 string) {
	mm_atomic.AddUint64(&mmDescription.beforeDescriptionCounter, 1)
	defer mm_atomic.AddUint64(&mmDescription.afterDescriptionCounter, 1)

	if mmDescription.inspectFuncDescription != nil {
		mmDescription.inspectFuncDescription()
	}

	if mmDescription.DescriptionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDescription.DescriptionMock.defaultExpectation.Counter, 1)

		mm_results := mmDescription.DescriptionMock.defaultExpectation.results
		if mm_results == nil {
			mmDescription.t.Fatal("No results are set for the CliCommandMock.Description")
		}
		return (*mm_results).s1
	}
	if mmDescription.funcDescription != nil {
		return mmDescription.funcDescription()
	}
	mmDescription.t.Fatalf("Unexpected call to CliCommandMock.Description.")
	return
}

// DescriptionAfterCounter returns a count of finished CliCommandMock.Description invocations
func (mmDescription *CliCommandMock) DescriptionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDescription.afterDescriptionCounter)
}

// DescriptionBeforeCounter returns a count of CliCommandMock.Description invocations
func (mmDescription *CliCommandMock) DescriptionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDescription.beforeDescriptionCounter)
}

// MinimockDescriptionDone returns true if the count of the Description invocations corresponds
// the number of defined expectations
func (m *CliCommandMock) MinimockDescriptionDone() bool {
	for _, e := range m.DescriptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DescriptionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDescriptionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDescription != nil && mm_atomic.LoadUint64(&m.afterDescriptionCounter) < 1 {
		return false
	}
	return true
}

// MinimockDescriptionInspect logs each unmet expectation
func (m *CliCommandMock) MinimockDescriptionInspect() {
	for _, e := range m.DescriptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CliCommandMock.Description")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DescriptionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDescriptionCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Description")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDescription != nil && mm_atomic.LoadUint64(&m.afterDescriptionCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Description")
	}
}

type mCliCommandMockInit struct {
	mock               *CliCommandMock
	defaultExpectation *CliCommandMockInitExpectation
	expectations       []*CliCommandMockInitExpectation

	callArgs []*CliCommandMockInitParams
	mutex    sync.RWMutex
}

// CliCommandMockInitExpectation specifies expectation struct of the CliCommand.Init
type CliCommandMockInitExpectation struct {
	mock    *CliCommandMock
	params  *CliCommandMockInitParams
	results *CliCommandMockInitResults
	Counter uint64
}

// CliCommandMockInitParams contains parameters of the CliCommand.Init
type CliCommandMockInitParams struct {
	args []string
}

// CliCommandMockInitResults contains results of the CliCommand.Init
type CliCommandMockInitResults struct {
	err error
}

// Expect sets up expected params for CliCommand.Init
func (mmInit *mCliCommandMockInit) Expect(args []string) *mCliCommandMockInit {
	if mmInit.mock.funcInit != nil {
		mmInit.mock.t.Fatalf("CliCommandMock.Init mock is already set by Set")
	}

	if mmInit.defaultExpectation == nil {
		mmInit.defaultExpectation = &CliCommandMockInitExpectation{}
	}

	mmInit.defaultExpectation.params = &CliCommandMockInitParams{args}
	for _, e := range mmInit.expectations {
		if minimock.Equal(e.params, mmInit.defaultExpectation.params) {
			mmInit.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInit.defaultExpectation.params)
		}
	}

	return mmInit
}

// Inspect accepts an inspector function that has same arguments as the CliCommand.Init
func (mmInit *mCliCommandMockInit) Inspect(f func(args []string)) *mCliCommandMockInit {
	if mmInit.mock.inspectFuncInit != nil {
		mmInit.mock.t.Fatalf("Inspect function is already set for CliCommandMock.Init")
	}

	mmInit.mock.inspectFuncInit = f

	return mmInit
}

// Return sets up results that will be returned by CliCommand.Init
func (mmInit *mCliCommandMockInit) Return(err error) *CliCommandMock {
	if mmInit.mock.funcInit != nil {
		mmInit.mock.t.Fatalf("CliCommandMock.Init mock is already set by Set")
	}

	if mmInit.defaultExpectation == nil {
		mmInit.defaultExpectation = &CliCommandMockInitExpectation{mock: mmInit.mock}
	}
	mmInit.defaultExpectation.results = &CliCommandMockInitResults{err}
	return mmInit.mock
}

//Set uses given function f to mock the CliCommand.Init method
func (mmInit *mCliCommandMockInit) Set(f func(args []string) (err error)) *CliCommandMock {
	if mmInit.defaultExpectation != nil {
		mmInit.mock.t.Fatalf("Default expectation is already set for the CliCommand.Init method")
	}

	if len(mmInit.expectations) > 0 {
		mmInit.mock.t.Fatalf("Some expectations are already set for the CliCommand.Init method")
	}

	mmInit.mock.funcInit = f
	return mmInit.mock
}

// When sets expectation for the CliCommand.Init which will trigger the result defined by the following
// Then helper
func (mmInit *mCliCommandMockInit) When(args []string) *CliCommandMockInitExpectation {
	if mmInit.mock.funcInit != nil {
		mmInit.mock.t.Fatalf("CliCommandMock.Init mock is already set by Set")
	}

	expectation := &CliCommandMockInitExpectation{
		mock:   mmInit.mock,
		params: &CliCommandMockInitParams{args},
	}
	mmInit.expectations = append(mmInit.expectations, expectation)
	return expectation
}

// Then sets up CliCommand.Init return parameters for the expectation previously defined by the When method
func (e *CliCommandMockInitExpectation) Then(err error) *CliCommandMock {
	e.results = &CliCommandMockInitResults{err}
	return e.mock
}

// Init implements internal.CliCommand
func (mmInit *CliCommandMock) Init(args []string) (err error) {
	mm_atomic.AddUint64(&mmInit.beforeInitCounter, 1)
	defer mm_atomic.AddUint64(&mmInit.afterInitCounter, 1)

	if mmInit.inspectFuncInit != nil {
		mmInit.inspectFuncInit(args)
	}

	mm_params := &CliCommandMockInitParams{args}

	// Record call args
	mmInit.InitMock.mutex.Lock()
	mmInit.InitMock.callArgs = append(mmInit.InitMock.callArgs, mm_params)
	mmInit.InitMock.mutex.Unlock()

	for _, e := range mmInit.InitMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmInit.InitMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInit.InitMock.defaultExpectation.Counter, 1)
		mm_want := mmInit.InitMock.defaultExpectation.params
		mm_got := CliCommandMockInitParams{args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInit.t.Errorf("CliCommandMock.Init got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInit.InitMock.defaultExpectation.results
		if mm_results == nil {
			mmInit.t.Fatal("No results are set for the CliCommandMock.Init")
		}
		return (*mm_results).err
	}
	if mmInit.funcInit != nil {
		return mmInit.funcInit(args)
	}
	mmInit.t.Fatalf("Unexpected call to CliCommandMock.Init. %v", args)
	return
}

// InitAfterCounter returns a count of finished CliCommandMock.Init invocations
func (mmInit *CliCommandMock) InitAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInit.afterInitCounter)
}

// InitBeforeCounter returns a count of CliCommandMock.Init invocations
func (mmInit *CliCommandMock) InitBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInit.beforeInitCounter)
}

// Calls returns a list of arguments used in each call to CliCommandMock.Init.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInit *mCliCommandMockInit) Calls() []*CliCommandMockInitParams {
	mmInit.mutex.RLock()

	argCopy := make([]*CliCommandMockInitParams, len(mmInit.callArgs))
	copy(argCopy, mmInit.callArgs)

	mmInit.mutex.RUnlock()

	return argCopy
}

// MinimockInitDone returns true if the count of the Init invocations corresponds
// the number of defined expectations
func (m *CliCommandMock) MinimockInitDone() bool {
	for _, e := range m.InitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInitCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInit != nil && mm_atomic.LoadUint64(&m.afterInitCounter) < 1 {
		return false
	}
	return true
}

// MinimockInitInspect logs each unmet expectation
func (m *CliCommandMock) MinimockInitInspect() {
	for _, e := range m.InitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CliCommandMock.Init with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInitCounter) < 1 {
		if m.InitMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CliCommandMock.Init")
		} else {
			m.t.Errorf("Expected call to CliCommandMock.Init with params: %#v", *m.InitMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInit != nil && mm_atomic.LoadUint64(&m.afterInitCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Init")
	}
}

type mCliCommandMockName struct {
	mock               *CliCommandMock
	defaultExpectation *CliCommandMockNameExpectation
	expectations       []*CliCommandMockNameExpectation
}

// CliCommandMockNameExpectation specifies expectation struct of the CliCommand.Name
type CliCommandMockNameExpectation struct {
	mock *CliCommandMock

	results *CliCommandMockNameResults
	Counter uint64
}

// CliCommandMockNameResults contains results of the CliCommand.Name
type CliCommandMockNameResults struct {
	s1 string
}

// Expect sets up expected params for CliCommand.Name
func (mmName *mCliCommandMockName) Expect() *mCliCommandMockName {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("CliCommandMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &CliCommandMockNameExpectation{}
	}

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the CliCommand.Name
func (mmName *mCliCommandMockName) Inspect(f func()) *mCliCommandMockName {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for CliCommandMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by CliCommand.Name
func (mmName *mCliCommandMockName) Return(s1 string) *CliCommandMock {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("CliCommandMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &CliCommandMockNameExpectation{mock: mmName.mock}
	}
	mmName.defaultExpectation.results = &CliCommandMockNameResults{s1}
	return mmName.mock
}

//Set uses given function f to mock the CliCommand.Name method
func (mmName *mCliCommandMockName) Set(f func() (s1 string)) *CliCommandMock {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the CliCommand.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the CliCommand.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Name implements internal.CliCommand
func (mmName *CliCommandMock) Name() (s1 string) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName()
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)

		mm_results := mmName.NameMock.defaultExpectation.results
		if mm_results == nil {
			mmName.t.Fatal("No results are set for the CliCommandMock.Name")
		}
		return (*mm_results).s1
	}
	if mmName.funcName != nil {
		return mmName.funcName()
	}
	mmName.t.Fatalf("Unexpected call to CliCommandMock.Name.")
	return
}

// NameAfterCounter returns a count of finished CliCommandMock.Name invocations
func (mmName *CliCommandMock) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of CliCommandMock.Name invocations
func (mmName *CliCommandMock) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *CliCommandMock) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockNameInspect logs each unmet expectation
func (m *CliCommandMock) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CliCommandMock.Name")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Name")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Name")
	}
}

type mCliCommandMockRun struct {
	mock               *CliCommandMock
	defaultExpectation *CliCommandMockRunExpectation
	expectations       []*CliCommandMockRunExpectation

	callArgs []*CliCommandMockRunParams
	mutex    sync.RWMutex
}

// CliCommandMockRunExpectation specifies expectation struct of the CliCommand.Run
type CliCommandMockRunExpectation struct {
	mock    *CliCommandMock
	params  *CliCommandMockRunParams
	results *CliCommandMockRunResults
	Counter uint64
}

// CliCommandMockRunParams contains parameters of the CliCommand.Run
type CliCommandMockRunParams struct {
	ctx mm_internal.ExecutionContext
}

// CliCommandMockRunResults contains results of the CliCommand.Run
type CliCommandMockRunResults struct {
	err error
}

// Expect sets up expected params for CliCommand.Run
func (mmRun *mCliCommandMockRun) Expect(ctx mm_internal.ExecutionContext) *mCliCommandMockRun {
	if mmRun.mock.funcRun != nil {
		mmRun.mock.t.Fatalf("CliCommandMock.Run mock is already set by Set")
	}

	if mmRun.defaultExpectation == nil {
		mmRun.defaultExpectation = &CliCommandMockRunExpectation{}
	}

	mmRun.defaultExpectation.params = &CliCommandMockRunParams{ctx}
	for _, e := range mmRun.expectations {
		if minimock.Equal(e.params, mmRun.defaultExpectation.params) {
			mmRun.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRun.defaultExpectation.params)
		}
	}

	return mmRun
}

// Inspect accepts an inspector function that has same arguments as the CliCommand.Run
func (mmRun *mCliCommandMockRun) Inspect(f func(ctx mm_internal.ExecutionContext)) *mCliCommandMockRun {
	if mmRun.mock.inspectFuncRun != nil {
		mmRun.mock.t.Fatalf("Inspect function is already set for CliCommandMock.Run")
	}

	mmRun.mock.inspectFuncRun = f

	return mmRun
}

// Return sets up results that will be returned by CliCommand.Run
func (mmRun *mCliCommandMockRun) Return(err error) *CliCommandMock {
	if mmRun.mock.funcRun != nil {
		mmRun.mock.t.Fatalf("CliCommandMock.Run mock is already set by Set")
	}

	if mmRun.defaultExpectation == nil {
		mmRun.defaultExpectation = &CliCommandMockRunExpectation{mock: mmRun.mock}
	}
	mmRun.defaultExpectation.results = &CliCommandMockRunResults{err}
	return mmRun.mock
}

//Set uses given function f to mock the CliCommand.Run method
func (mmRun *mCliCommandMockRun) Set(f func(ctx mm_internal.ExecutionContext) (err error)) *CliCommandMock {
	if mmRun.defaultExpectation != nil {
		mmRun.mock.t.Fatalf("Default expectation is already set for the CliCommand.Run method")
	}

	if len(mmRun.expectations) > 0 {
		mmRun.mock.t.Fatalf("Some expectations are already set for the CliCommand.Run method")
	}

	mmRun.mock.funcRun = f
	return mmRun.mock
}

// When sets expectation for the CliCommand.Run which will trigger the result defined by the following
// Then helper
func (mmRun *mCliCommandMockRun) When(ctx mm_internal.ExecutionContext) *CliCommandMockRunExpectation {
	if mmRun.mock.funcRun != nil {
		mmRun.mock.t.Fatalf("CliCommandMock.Run mock is already set by Set")
	}

	expectation := &CliCommandMockRunExpectation{
		mock:   mmRun.mock,
		params: &CliCommandMockRunParams{ctx},
	}
	mmRun.expectations = append(mmRun.expectations, expectation)
	return expectation
}

// Then sets up CliCommand.Run return parameters for the expectation previously defined by the When method
func (e *CliCommandMockRunExpectation) Then(err error) *CliCommandMock {
	e.results = &CliCommandMockRunResults{err}
	return e.mock
}

// Run implements internal.CliCommand
func (mmRun *CliCommandMock) Run(ctx mm_internal.ExecutionContext) (err error) {
	mm_atomic.AddUint64(&mmRun.beforeRunCounter, 1)
	defer mm_atomic.AddUint64(&mmRun.afterRunCounter, 1)

	if mmRun.inspectFuncRun != nil {
		mmRun.inspectFuncRun(ctx)
	}

	mm_params := &CliCommandMockRunParams{ctx}

	// Record call args
	mmRun.RunMock.mutex.Lock()
	mmRun.RunMock.callArgs = append(mmRun.RunMock.callArgs, mm_params)
	mmRun.RunMock.mutex.Unlock()

	for _, e := range mmRun.RunMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmRun.RunMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRun.RunMock.defaultExpectation.Counter, 1)
		mm_want := mmRun.RunMock.defaultExpectation.params
		mm_got := CliCommandMockRunParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRun.t.Errorf("CliCommandMock.Run got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRun.RunMock.defaultExpectation.results
		if mm_results == nil {
			mmRun.t.Fatal("No results are set for the CliCommandMock.Run")
		}
		return (*mm_results).err
	}
	if mmRun.funcRun != nil {
		return mmRun.funcRun(ctx)
	}
	mmRun.t.Fatalf("Unexpected call to CliCommandMock.Run. %v", ctx)
	return
}

// RunAfterCounter returns a count of finished CliCommandMock.Run invocations
func (mmRun *CliCommandMock) RunAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRun.afterRunCounter)
}

// RunBeforeCounter returns a count of CliCommandMock.Run invocations
func (mmRun *CliCommandMock) RunBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRun.beforeRunCounter)
}

// Calls returns a list of arguments used in each call to CliCommandMock.Run.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRun *mCliCommandMockRun) Calls() []*CliCommandMockRunParams {
	mmRun.mutex.RLock()

	argCopy := make([]*CliCommandMockRunParams, len(mmRun.callArgs))
	copy(argCopy, mmRun.callArgs)

	mmRun.mutex.RUnlock()

	return argCopy
}

// MinimockRunDone returns true if the count of the Run invocations corresponds
// the number of defined expectations
func (m *CliCommandMock) MinimockRunDone() bool {
	for _, e := range m.RunMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RunMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRunCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRun != nil && mm_atomic.LoadUint64(&m.afterRunCounter) < 1 {
		return false
	}
	return true
}

// MinimockRunInspect logs each unmet expectation
func (m *CliCommandMock) MinimockRunInspect() {
	for _, e := range m.RunMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CliCommandMock.Run with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RunMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRunCounter) < 1 {
		if m.RunMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CliCommandMock.Run")
		} else {
			m.t.Errorf("Expected call to CliCommandMock.Run with params: %#v", *m.RunMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRun != nil && mm_atomic.LoadUint64(&m.afterRunCounter) < 1 {
		m.t.Error("Expected call to CliCommandMock.Run")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CliCommandMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDescriptionInspect()

		m.MinimockInitInspect()

		m.MinimockNameInspect()

		m.MinimockRunInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CliCommandMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CliCommandMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDescriptionDone() &&
		m.MinimockInitDone() &&
		m.MinimockNameDone() &&
		m.MinimockRunDone()
}
