// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Includes the broad cateogry of machine used for this specific task execution.
type TaskExecutionMetadata_InstanceClass int32

const (
	// The default instance class configured for the flyte application platform.
	TaskExecutionMetadata_DEFAULT TaskExecutionMetadata_InstanceClass = 0
	// The instance class configured for interruptible tasks.
	TaskExecutionMetadata_INTERRUPTIBLE TaskExecutionMetadata_InstanceClass = 1
)

var TaskExecutionMetadata_InstanceClass_name = map[int32]string{
	0: "DEFAULT",
	1: "INTERRUPTIBLE",
}

var TaskExecutionMetadata_InstanceClass_value = map[string]int32{
	"DEFAULT":       0,
	"INTERRUPTIBLE": 1,
}

func (x TaskExecutionMetadata_InstanceClass) String() string {
	return proto.EnumName(TaskExecutionMetadata_InstanceClass_name, int32(x))
}

func (TaskExecutionMetadata_InstanceClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{8, 0}
}

type WorkflowExecutionEvent struct {
	// Workflow execution id
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                       `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.WorkflowExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the workflow.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*WorkflowExecutionEvent_OutputUri
	//	*WorkflowExecutionEvent_Error
	OutputResult         isWorkflowExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkflowExecutionEvent) Reset()         { *m = WorkflowExecutionEvent{} }
func (m *WorkflowExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEvent) ProtoMessage()    {}
func (*WorkflowExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{0}
}

func (m *WorkflowExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEvent.Unmarshal(m, b)
}
func (m *WorkflowExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEvent.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEvent.Merge(m, src)
}
func (m *WorkflowExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEvent.Size(m)
}
func (m *WorkflowExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEvent proto.InternalMessageInfo

func (m *WorkflowExecutionEvent) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetPhase() core.WorkflowExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecution_UNDEFINED
}

func (m *WorkflowExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isWorkflowExecutionEvent_OutputResult interface {
	isWorkflowExecutionEvent_OutputResult()
}

type WorkflowExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,5,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type WorkflowExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*WorkflowExecutionEvent_OutputUri) isWorkflowExecutionEvent_OutputResult() {}

func (*WorkflowExecutionEvent_Error) isWorkflowExecutionEvent_OutputResult() {}

func (m *WorkflowExecutionEvent) GetOutputResult() isWorkflowExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkflowExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkflowExecutionEvent_OutputUri)(nil),
		(*WorkflowExecutionEvent_Error)(nil),
	}
}

type NodeExecutionEvent struct {
	// Unique identifier for this node execution
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                   `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.NodeExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the node.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	InputUri   string               `protobuf:"bytes,5,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionEvent_OutputUri
	//	*NodeExecutionEvent_Error
	OutputResult isNodeExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Additional metadata to do with this event's node target based
	// on the node type
	//
	// Types that are valid to be assigned to TargetMetadata:
	//	*NodeExecutionEvent_WorkflowNodeMetadata
	//	*NodeExecutionEvent_TaskNodeMetadata
	TargetMetadata isNodeExecutionEvent_TargetMetadata `protobuf_oneof:"target_metadata"`
	// [To be deprecated] Specifies which task (if any) launched this node.
	ParentTaskMetadata *ParentTaskExecutionMetadata `protobuf:"bytes,9,opt,name=parent_task_metadata,json=parentTaskMetadata,proto3" json:"parent_task_metadata,omitempty"`
	// Specifies the parent node of the current node execution. Node executions at level zero will not have a parent node.
	ParentNodeMetadata *ParentNodeExecutionMetadata `protobuf:"bytes,10,opt,name=parent_node_metadata,json=parentNodeMetadata,proto3" json:"parent_node_metadata,omitempty"`
	// Retry group to indicate grouping of nodes by retries
	RetryGroup string `protobuf:"bytes,11,opt,name=retry_group,json=retryGroup,proto3" json:"retry_group,omitempty"`
	// Identifier of the node in the original workflow/graph
	// This maps to value of WorkflowTemplate.nodes[X].id
	SpecNodeId string `protobuf:"bytes,12,opt,name=spec_node_id,json=specNodeId,proto3" json:"spec_node_id,omitempty"`
	// Friendly readable name for the node
	NodeName             string   `protobuf:"bytes,13,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEvent) Reset()         { *m = NodeExecutionEvent{} }
func (m *NodeExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEvent) ProtoMessage()    {}
func (*NodeExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{1}
}

func (m *NodeExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEvent.Unmarshal(m, b)
}
func (m *NodeExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEvent.Marshal(b, m, deterministic)
}
func (m *NodeExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEvent.Merge(m, src)
}
func (m *NodeExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEvent.Size(m)
}
func (m *NodeExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEvent proto.InternalMessageInfo

func (m *NodeExecutionEvent) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *NodeExecutionEvent) GetPhase() core.NodeExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecution_UNDEFINED
}

func (m *NodeExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isNodeExecutionEvent_OutputResult interface {
	isNodeExecutionEvent_OutputResult()
}

type NodeExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,6,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionEvent_OutputUri) isNodeExecutionEvent_OutputResult() {}

func (*NodeExecutionEvent_Error) isNodeExecutionEvent_OutputResult() {}

func (m *NodeExecutionEvent) GetOutputResult() isNodeExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

type isNodeExecutionEvent_TargetMetadata interface {
	isNodeExecutionEvent_TargetMetadata()
}

type NodeExecutionEvent_WorkflowNodeMetadata struct {
	WorkflowNodeMetadata *WorkflowNodeMetadata `protobuf:"bytes,8,opt,name=workflow_node_metadata,json=workflowNodeMetadata,proto3,oneof"`
}

type NodeExecutionEvent_TaskNodeMetadata struct {
	TaskNodeMetadata *TaskNodeMetadata `protobuf:"bytes,14,opt,name=task_node_metadata,json=taskNodeMetadata,proto3,oneof"`
}

func (*NodeExecutionEvent_WorkflowNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (*NodeExecutionEvent_TaskNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (m *NodeExecutionEvent) GetTargetMetadata() isNodeExecutionEvent_TargetMetadata {
	if m != nil {
		return m.TargetMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetWorkflowNodeMetadata() *WorkflowNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_WorkflowNodeMetadata); ok {
		return x.WorkflowNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetTaskNodeMetadata() *TaskNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_TaskNodeMetadata); ok {
		return x.TaskNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentTaskMetadata() *ParentTaskExecutionMetadata {
	if m != nil {
		return m.ParentTaskMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentNodeMetadata() *ParentNodeExecutionMetadata {
	if m != nil {
		return m.ParentNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetRetryGroup() string {
	if m != nil {
		return m.RetryGroup
	}
	return ""
}

func (m *NodeExecutionEvent) GetSpecNodeId() string {
	if m != nil {
		return m.SpecNodeId
	}
	return ""
}

func (m *NodeExecutionEvent) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeExecutionEvent_OutputUri)(nil),
		(*NodeExecutionEvent_Error)(nil),
		(*NodeExecutionEvent_WorkflowNodeMetadata)(nil),
		(*NodeExecutionEvent_TaskNodeMetadata)(nil),
	}
}

// For Workflow Nodes we need to send information about the workflow that's launched
type WorkflowNodeMetadata struct {
	ExecutionId          *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowNodeMetadata) Reset()         { *m = WorkflowNodeMetadata{} }
func (m *WorkflowNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowNodeMetadata) ProtoMessage()    {}
func (*WorkflowNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{2}
}

func (m *WorkflowNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNodeMetadata.Unmarshal(m, b)
}
func (m *WorkflowNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNodeMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNodeMetadata.Merge(m, src)
}
func (m *WorkflowNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowNodeMetadata.Size(m)
}
func (m *WorkflowNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNodeMetadata proto.InternalMessageInfo

func (m *WorkflowNodeMetadata) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

type TaskNodeMetadata struct {
	// Captures the status of caching for this execution.
	CacheStatus core.CatalogCacheStatus `protobuf:"varint,1,opt,name=cache_status,json=cacheStatus,proto3,enum=flyteidl.core.CatalogCacheStatus" json:"cache_status,omitempty"`
	// This structure carries the catalog artifact information
	CatalogKey           *core.CatalogMetadata `protobuf:"bytes,2,opt,name=catalog_key,json=catalogKey,proto3" json:"catalog_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskNodeMetadata) Reset()         { *m = TaskNodeMetadata{} }
func (m *TaskNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskNodeMetadata) ProtoMessage()    {}
func (*TaskNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{3}
}

func (m *TaskNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNodeMetadata.Unmarshal(m, b)
}
func (m *TaskNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNodeMetadata.Marshal(b, m, deterministic)
}
func (m *TaskNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNodeMetadata.Merge(m, src)
}
func (m *TaskNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskNodeMetadata.Size(m)
}
func (m *TaskNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNodeMetadata proto.InternalMessageInfo

func (m *TaskNodeMetadata) GetCacheStatus() core.CatalogCacheStatus {
	if m != nil {
		return m.CacheStatus
	}
	return core.CatalogCacheStatus_CACHE_DISABLED
}

func (m *TaskNodeMetadata) GetCatalogKey() *core.CatalogMetadata {
	if m != nil {
		return m.CatalogKey
	}
	return nil
}

type ParentTaskExecutionMetadata struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ParentTaskExecutionMetadata) Reset()         { *m = ParentTaskExecutionMetadata{} }
func (m *ParentTaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentTaskExecutionMetadata) ProtoMessage()    {}
func (*ParentTaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{4}
}

func (m *ParentTaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentTaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *ParentTaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentTaskExecutionMetadata.Merge(m, src)
}
func (m *ParentTaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Size(m)
}
func (m *ParentTaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentTaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentTaskExecutionMetadata proto.InternalMessageInfo

func (m *ParentTaskExecutionMetadata) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

type ParentNodeExecutionMetadata struct {
	// Unique identifier of the parent node id within the execution
	// This is value of core.NodeExecutionIdentifier.node_id of the parent node
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentNodeExecutionMetadata) Reset()         { *m = ParentNodeExecutionMetadata{} }
func (m *ParentNodeExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentNodeExecutionMetadata) ProtoMessage()    {}
func (*ParentNodeExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{5}
}

func (m *ParentNodeExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentNodeExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *ParentNodeExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentNodeExecutionMetadata.Merge(m, src)
}
func (m *ParentNodeExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Size(m)
}
func (m *ParentNodeExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentNodeExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentNodeExecutionMetadata proto.InternalMessageInfo

func (m *ParentNodeExecutionMetadata) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

// Plugin specific execution event information. For tasks like Python, Hive, Spark, DynamicJob.
type TaskExecutionEvent struct {
	// ID of the task. In combination with the retryAttempt this will indicate
	// the task execution uniquely for a given parent node execution.
	TaskId *core.Identifier `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A task execution is always kicked off by a node execution, the event consumer
	// will use the parent_id to relate the task to it's parent node execution
	ParentNodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=parent_node_execution_id,json=parentNodeExecutionId,proto3" json:"parent_node_execution_id,omitempty"`
	// retry attempt number for this task, ie., 2 for the second attempt
	RetryAttempt uint32 `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	// Phase associated with the event
	Phase core.TaskExecution_Phase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.TaskExecution_Phase" json:"phase,omitempty"`
	// id of the process that sent this event, mainly for trace debugging
	ProducerId string `protobuf:"bytes,5,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	// log information for the task execution
	Logs []*core.TaskLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the task.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// URI of the input file, it encodes all the information
	// including Cloud source provider. ie., s3://...
	InputUri string `protobuf:"bytes,8,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionEvent_OutputUri
	//	*TaskExecutionEvent_Error
	OutputResult isTaskExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.
	CustomInfo *_struct.Struct `protobuf:"bytes,11,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	// Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc)
	// that should be recorded regardless of the lack of phase change.
	// The version field should be incremented when metadata changes across the duration of an individual phase.
	PhaseVersion uint32 `protobuf:"varint,12,opt,name=phase_version,json=phaseVersion,proto3" json:"phase_version,omitempty"`
	// If there is an explanation for this phase transition, the reason will capture it.
	Reason string `protobuf:"bytes,13,opt,name=reason,proto3" json:"reason,omitempty"`
	// A predefined yet extensible Task type identifier. If the task definition is already registered in flyte admin
	// this type will be identical, but not all task executions necessarily use pre-registered definitions and this
	// type is useful to render the task in the UI, filter task executions, etc.
	TaskType string `protobuf:"bytes,14,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata around how a task was executed.
	Metadata             *TaskExecutionMetadata `protobuf:"bytes,16,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TaskExecutionEvent) Reset()         { *m = TaskExecutionEvent{} }
func (m *TaskExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEvent) ProtoMessage()    {}
func (*TaskExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{6}
}

func (m *TaskExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEvent.Unmarshal(m, b)
}
func (m *TaskExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEvent.Marshal(b, m, deterministic)
}
func (m *TaskExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEvent.Merge(m, src)
}
func (m *TaskExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEvent.Size(m)
}
func (m *TaskExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEvent proto.InternalMessageInfo

func (m *TaskExecutionEvent) GetTaskId() *core.Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionEvent) GetParentNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.ParentNodeExecutionId
	}
	return nil
}

func (m *TaskExecutionEvent) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

func (m *TaskExecutionEvent) GetPhase() core.TaskExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecution_UNDEFINED
}

func (m *TaskExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *TaskExecutionEvent) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isTaskExecutionEvent_OutputResult interface {
	isTaskExecutionEvent_OutputResult()
}

type TaskExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,9,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionEvent_OutputUri) isTaskExecutionEvent_OutputResult() {}

func (*TaskExecutionEvent_Error) isTaskExecutionEvent_OutputResult() {}

func (m *TaskExecutionEvent) GetOutputResult() isTaskExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionEvent) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

func (m *TaskExecutionEvent) GetPhaseVersion() uint32 {
	if m != nil {
		return m.PhaseVersion
	}
	return 0
}

func (m *TaskExecutionEvent) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TaskExecutionEvent) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *TaskExecutionEvent) GetMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskExecutionEvent_OutputUri)(nil),
		(*TaskExecutionEvent_Error)(nil),
	}
}

// This message holds task execution metadata specific to resource allocation used to manage concurrent
// executions for a project namespace.
type ManagedResourceInfo struct {
	// Unique resource ID used to identify this execution when allocating a token.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedResourceInfo) Reset()         { *m = ManagedResourceInfo{} }
func (m *ManagedResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ManagedResourceInfo) ProtoMessage()    {}
func (*ManagedResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{7}
}

func (m *ManagedResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedResourceInfo.Unmarshal(m, b)
}
func (m *ManagedResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedResourceInfo.Marshal(b, m, deterministic)
}
func (m *ManagedResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedResourceInfo.Merge(m, src)
}
func (m *ManagedResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ManagedResourceInfo.Size(m)
}
func (m *ManagedResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedResourceInfo proto.InternalMessageInfo

func (m *ManagedResourceInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Holds metadata around how a task was executed.
// TODO(katrogan): Extend to include freeform fields (https://github.com/flyteorg/flyte/issues/325).
type TaskExecutionMetadata struct {
	InstanceClass TaskExecutionMetadata_InstanceClass `protobuf:"varint,16,opt,name=instance_class,json=instanceClass,proto3,enum=flyteidl.event.TaskExecutionMetadata_InstanceClass" json:"instance_class,omitempty"`
	// Generated unique name for this task execution used by the backend.
	GeneratedName string `protobuf:"bytes,1,opt,name=generated_name,json=generatedName,proto3" json:"generated_name,omitempty"`
	// Includes information about how resource token allocation (if applicable).
	ManagedResourceInfo  *ManagedResourceInfo `protobuf:"bytes,2,opt,name=managed_resource_info,json=managedResourceInfo,proto3" json:"managed_resource_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskExecutionMetadata) Reset()         { *m = TaskExecutionMetadata{} }
func (m *TaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionMetadata) ProtoMessage()    {}
func (*TaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{8}
}

func (m *TaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionMetadata.Unmarshal(m, b)
}
func (m *TaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *TaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionMetadata.Merge(m, src)
}
func (m *TaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionMetadata.Size(m)
}
func (m *TaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionMetadata proto.InternalMessageInfo

func (m *TaskExecutionMetadata) GetInstanceClass() TaskExecutionMetadata_InstanceClass {
	if m != nil {
		return m.InstanceClass
	}
	return TaskExecutionMetadata_DEFAULT
}

func (m *TaskExecutionMetadata) GetGeneratedName() string {
	if m != nil {
		return m.GeneratedName
	}
	return ""
}

func (m *TaskExecutionMetadata) GetManagedResourceInfo() *ManagedResourceInfo {
	if m != nil {
		return m.ManagedResourceInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.event.TaskExecutionMetadata_InstanceClass", TaskExecutionMetadata_InstanceClass_name, TaskExecutionMetadata_InstanceClass_value)
	proto.RegisterType((*WorkflowExecutionEvent)(nil), "flyteidl.event.WorkflowExecutionEvent")
	proto.RegisterType((*NodeExecutionEvent)(nil), "flyteidl.event.NodeExecutionEvent")
	proto.RegisterType((*WorkflowNodeMetadata)(nil), "flyteidl.event.WorkflowNodeMetadata")
	proto.RegisterType((*TaskNodeMetadata)(nil), "flyteidl.event.TaskNodeMetadata")
	proto.RegisterType((*ParentTaskExecutionMetadata)(nil), "flyteidl.event.ParentTaskExecutionMetadata")
	proto.RegisterType((*ParentNodeExecutionMetadata)(nil), "flyteidl.event.ParentNodeExecutionMetadata")
	proto.RegisterType((*TaskExecutionEvent)(nil), "flyteidl.event.TaskExecutionEvent")
	proto.RegisterType((*ManagedResourceInfo)(nil), "flyteidl.event.ManagedResourceInfo")
	proto.RegisterType((*TaskExecutionMetadata)(nil), "flyteidl.event.TaskExecutionMetadata")
}

func init() { proto.RegisterFile("flyteidl/event/event.proto", fileDescriptor_4b035d24120b1b12) }

var fileDescriptor_4b035d24120b1b12 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xef, 0x6e, 0xdb, 0xb6,
	0x17, 0x8d, 0xdd, 0xd8, 0x89, 0xaf, 0x6c, 0x37, 0x65, 0x93, 0x54, 0xbf, 0xe4, 0xd7, 0xc6, 0x73,
	0xd7, 0x21, 0x68, 0x31, 0x19, 0x70, 0xd1, 0xae, 0xc0, 0x06, 0x0c, 0x49, 0x9a, 0x2d, 0xc6, 0x92,
	0x20, 0x50, 0x9d, 0x15, 0x28, 0x36, 0x08, 0x8c, 0x74, 0xad, 0x08, 0xb1, 0x45, 0x81, 0xa2, 0x9a,
	0xf9, 0x51, 0xf6, 0x7d, 0x6f, 0xb3, 0x57, 0x19, 0xb0, 0x57, 0x18, 0x48, 0xca, 0x4a, 0x24, 0x0b,
	0x5e, 0x33, 0xf4, 0x8b, 0x01, 0x1e, 0x5e, 0x1e, 0x9e, 0xfb, 0x87, 0xc7, 0x82, 0xad, 0xd1, 0x78,
	0x2a, 0x30, 0xf0, 0xc6, 0x3d, 0xfc, 0x88, 0xa1, 0xd0, 0xbf, 0x56, 0xc4, 0x99, 0x60, 0xa4, 0x3d,
	0xdb, 0xb3, 0x14, 0xba, 0xf5, 0x38, 0x8b, 0x75, 0x19, 0xc7, 0x1e, 0xfe, 0x86, 0x6e, 0x22, 0x02,
	0x16, 0xea, 0xf0, 0xad, 0x27, 0xf9, 0xed, 0xc0, 0xc3, 0x50, 0x04, 0xa3, 0x00, 0x79, 0xba, 0xbf,
	0x9d, 0xdf, 0x77, 0xa9, 0xa0, 0x63, 0xe6, 0xa7, 0x9b, 0x3b, 0x3e, 0x63, 0xfe, 0x18, 0x7b, 0x6a,
	0x75, 0x91, 0x8c, 0x7a, 0x22, 0x98, 0x60, 0x2c, 0xe8, 0x24, 0x4a, 0x03, 0xfe, 0x5f, 0x0c, 0x88,
	0x05, 0x4f, 0xdc, 0x54, 0x6a, 0xf7, 0xaf, 0x2a, 0x6c, 0xbe, 0x67, 0xfc, 0x6a, 0x34, 0x66, 0xd7,
	0x87, 0x33, 0x5d, 0x87, 0x52, 0x35, 0x39, 0x81, 0x66, 0xa6, 0xd4, 0x09, 0x3c, 0xb3, 0xd2, 0xa9,
	0xec, 0x1a, 0xfd, 0xe7, 0x56, 0x96, 0x9c, 0x54, 0x63, 0xcd, 0x1d, 0x1e, 0x64, 0xf2, 0x6d, 0x03,
	0x6f, 0x40, 0xb2, 0x03, 0x46, 0xc4, 0x99, 0x97, 0xb8, 0xc8, 0x25, 0x5b, 0xb5, 0x53, 0xd9, 0x6d,
	0xd8, 0x30, 0x83, 0x06, 0x1e, 0xf9, 0x0e, 0x6a, 0xd1, 0x25, 0x8d, 0xd1, 0xbc, 0xd7, 0xa9, 0xec,
	0xb6, 0xfb, 0x5f, 0xfd, 0xdb, 0x45, 0xd6, 0x99, 0x8c, 0xb6, 0xf5, 0x21, 0xf2, 0x2d, 0x18, 0xcc,
	0x75, 0x13, 0xce, 0xd1, 0x73, 0xa8, 0x30, 0x97, 0x95, 0xd8, 0x2d, 0x4b, 0x27, 0x6f, 0xcd, 0x92,
	0xb7, 0x86, 0xb3, 0xea, 0xd8, 0x30, 0x0b, 0xdf, 0x13, 0x64, 0x07, 0x80, 0x25, 0x22, 0x4a, 0x84,
	0x93, 0xf0, 0xc0, 0xac, 0x49, 0x69, 0x47, 0x4b, 0x76, 0x43, 0x63, 0xe7, 0x3c, 0x20, 0xaf, 0xa0,
	0x86, 0x9c, 0x33, 0x6e, 0xd6, 0x15, 0xef, 0xe3, 0x82, 0xb6, 0x9b, 0xca, 0xc9, 0xa0, 0xa3, 0x25,
	0x5b, 0x47, 0xef, 0xdf, 0x87, 0x56, 0xca, 0xcb, 0x31, 0x4e, 0xc6, 0xa2, 0xfb, 0x67, 0x1d, 0xc8,
	0x29, 0xf3, 0xb0, 0x50, 0xea, 0xd7, 0x50, 0xcd, 0x0a, 0x5c, 0xcc, 0x3b, 0x17, 0x7e, 0xab, 0xb8,
	0xd5, 0xe0, 0x13, 0x6a, 0xfa, 0x26, 0x5f, 0xd3, 0xee, 0x22, 0xee, 0xcf, 0x58, 0xcf, 0x6d, 0x68,
	0x04, 0x61, 0xae, 0x9c, 0xf6, 0xaa, 0x02, 0x64, 0x2d, 0xf3, 0xc5, 0xae, 0x2f, 0x28, 0xf6, 0xca,
	0x5d, 0x8a, 0x4d, 0x7e, 0x81, 0xcd, 0xeb, 0x74, 0x46, 0x9c, 0x90, 0x79, 0xe8, 0x4c, 0x50, 0x50,
	0x8f, 0x0a, 0x6a, 0xae, 0x2a, 0x9e, 0x2f, 0xad, 0xfc, 0xb3, 0xcc, 0x26, 0x4a, 0x56, 0xe1, 0x24,
	0x8d, 0x3d, 0xaa, 0xd8, 0xeb, 0xd7, 0x25, 0x38, 0x39, 0x03, 0x22, 0x68, 0x7c, 0x55, 0x60, 0x6e,
	0x2b, 0xe6, 0x4e, 0x91, 0x79, 0x48, 0xe3, 0xab, 0x02, 0xeb, 0x9a, 0x28, 0x60, 0xe4, 0x57, 0x58,
	0x8f, 0x28, 0xc7, 0x50, 0x38, 0x8a, 0x38, 0xe3, 0x6c, 0x28, 0xce, 0x17, 0x45, 0xce, 0x33, 0x15,
	0x2b, 0x99, 0xb3, 0x02, 0xcc, 0xa8, 0x6c, 0x12, 0x65, 0x9b, 0x25, 0xf4, 0x79, 0xc9, 0xb0, 0x88,
	0x3e, 0x37, 0x10, 0x45, 0xfa, 0x9c, 0xfa, 0x1d, 0x30, 0x38, 0x0a, 0x3e, 0x75, 0x7c, 0xce, 0x92,
	0xc8, 0x34, 0xf4, 0xe8, 0x29, 0xe8, 0x47, 0x89, 0x90, 0x0e, 0x34, 0xe3, 0x08, 0x5d, 0x7d, 0x7b,
	0xe0, 0x99, 0x4d, 0x1d, 0x21, 0x31, 0x49, 0x34, 0xf0, 0xe4, 0x94, 0xa8, 0xcd, 0x90, 0x4e, 0xd0,
	0x6c, 0xe9, 0x29, 0x91, 0xc0, 0x29, 0x9d, 0xe0, 0xdc, 0xd3, 0xd9, 0x7f, 0x00, 0xf7, 0x05, 0xe5,
	0x3e, 0x8a, 0x2c, 0x95, 0x2e, 0xc2, 0x7a, 0x59, 0x0f, 0x3f, 0xb3, 0x73, 0x75, 0x7f, 0xaf, 0xc0,
	0x5a, 0xb1, 0xa3, 0xe4, 0x2d, 0x34, 0x5d, 0xea, 0x5e, 0xa2, 0x13, 0x0b, 0x2a, 0x92, 0x58, 0xdd,
	0xd1, 0xee, 0x7f, 0x51, 0xb8, 0xe3, 0x40, 0x7b, 0xf5, 0x81, 0x8c, 0x7c, 0xa7, 0x02, 0x6d, 0xc3,
	0xbd, 0x59, 0x90, 0xef, 0xc1, 0x48, 0xed, 0xdc, 0xb9, 0xc2, 0xa9, 0x7a, 0xc0, 0x46, 0xff, 0x49,
	0x39, 0x49, 0xd6, 0x0e, 0x48, 0x8f, 0xfc, 0x84, 0xd3, 0xee, 0x39, 0x6c, 0x2f, 0x18, 0x8c, 0x85,
	0xc6, 0x92, 0x3b, 0x91, 0x37, 0x96, 0xee, 0xeb, 0x19, 0x6d, 0xe9, 0x40, 0x90, 0x47, 0xb0, 0x32,
	0x6b, 0x6b, 0x45, 0xf5, 0xad, 0x1e, 0xaa, 0x96, 0x76, 0xff, 0xae, 0x01, 0xc9, 0xf1, 0x6a, 0x7f,
	0xeb, 0xc3, 0x8a, 0x9a, 0xf1, 0x4c, 0xcb, 0xff, 0x0a, 0x5a, 0x6e, 0x5d, 0x5f, 0x97, 0x91, 0x03,
	0x8f, 0x38, 0x60, 0xde, 0x9e, 0xdf, 0x5c, 0x43, 0xab, 0x77, 0x72, 0xca, 0x8d, 0x68, 0x3e, 0x95,
	0x81, 0x47, 0x9e, 0x42, 0x4b, 0x4f, 0x30, 0x15, 0x02, 0x27, 0x91, 0x50, 0x1e, 0xd9, 0xb2, 0x9b,
	0x0a, 0xdc, 0xd3, 0xd8, 0x8d, 0x81, 0x2e, 0x97, 0x1a, 0x68, 0x2e, 0xd7, 0xbc, 0x81, 0x16, 0xbc,
	0xb9, 0x36, 0xe7, 0xcd, 0xcf, 0x61, 0x79, 0xcc, 0xfc, 0xd8, 0xac, 0x77, 0xee, 0xed, 0x1a, 0xfd,
	0xcd, 0x12, 0xe6, 0x63, 0xe6, 0xdb, 0x2a, 0xa6, 0xe8, 0xc6, 0x2b, 0xff, 0xdd, 0x8d, 0x57, 0x17,
	0xba, 0x71, 0x63, 0x81, 0x1b, 0xc3, 0x9d, 0xdc, 0xf8, 0x0d, 0x18, 0x6e, 0x12, 0x0b, 0x36, 0x71,
	0x82, 0x70, 0xc4, 0x94, 0x3f, 0x18, 0xfd, 0x47, 0x73, 0x8a, 0xdf, 0xa9, 0x8f, 0x11, 0x1b, 0x74,
	0xec, 0x20, 0x1c, 0x31, 0xd9, 0x17, 0x55, 0x41, 0xe7, 0x23, 0xf2, 0x38, 0x60, 0xa1, 0x72, 0x8e,
	0x96, 0xdd, 0x54, 0xe0, 0xcf, 0x1a, 0x23, 0x9b, 0x50, 0xe7, 0x48, 0x63, 0x16, 0xa6, 0xc6, 0x91,
	0xae, 0x64, 0xae, 0x6a, 0xd2, 0xc4, 0x34, 0x42, 0xe5, 0xce, 0x0d, 0x7b, 0x55, 0x02, 0xc3, 0x69,
	0x84, 0x64, 0x0f, 0x56, 0x33, 0x1b, 0x5c, 0x53, 0x82, 0x9e, 0x95, 0x39, 0xf7, 0xbc, 0x01, 0x66,
	0xc7, 0xe6, 0xff, 0xd1, 0x5f, 0xc0, 0xc3, 0x13, 0x1a, 0x52, 0x1f, 0x3d, 0x1b, 0x63, 0x96, 0x70,
	0x17, 0x55, 0x12, 0xeb, 0x50, 0x13, 0xec, 0x0a, 0xc3, 0xf4, 0x7d, 0xe8, 0x45, 0xf7, 0x8f, 0x2a,
	0x6c, 0x94, 0x3f, 0xd4, 0x0f, 0xd0, 0x0e, 0xc2, 0x58, 0xd0, 0xd0, 0x45, 0xc7, 0x1d, 0xd3, 0x38,
	0x56, 0x02, 0xdb, 0xfd, 0x97, 0x9f, 0x24, 0xd0, 0x1a, 0xa4, 0x67, 0x0f, 0xe4, 0x51, 0xbb, 0x15,
	0xdc, 0x5e, 0x92, 0x67, 0xd0, 0xf6, 0x31, 0x44, 0x4e, 0x05, 0x7a, 0xda, 0x6c, 0xb5, 0xa8, 0x56,
	0x86, 0x4a, 0xc7, 0x25, 0xef, 0x61, 0x63, 0xa2, 0x33, 0x91, 0xb9, 0xa9, 0x54, 0x74, 0xef, 0xf4,
	0x6b, 0x7b, 0x5a, 0x54, 0x52, 0x92, 0xb6, 0xfd, 0x70, 0x32, 0x0f, 0x76, 0x7b, 0xd0, 0xca, 0xe9,
	0x23, 0x06, 0xac, 0xbc, 0x3d, 0xfc, 0x61, 0xef, 0xfc, 0x78, 0xb8, 0xb6, 0x44, 0x1e, 0x40, 0x6b,
	0x70, 0x3a, 0x3c, 0xb4, 0xed, 0xf3, 0xb3, 0xe1, 0x60, 0xff, 0xf8, 0x70, 0xad, 0xb2, 0xff, 0xcd,
	0x87, 0x57, 0x7e, 0x20, 0x2e, 0x93, 0x0b, 0xcb, 0x65, 0x93, 0x9e, 0xba, 0x96, 0x71, 0xbf, 0x97,
	0x7d, 0x06, 0xfb, 0x18, 0xf6, 0xa2, 0x8b, 0xaf, 0x7d, 0xd6, 0xcb, 0x7f, 0x84, 0x5f, 0xd4, 0xd5,
	0x5c, 0xbd, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x27, 0xf3, 0x4f, 0x3c, 0x9d, 0x0b, 0x00, 0x00,
}
