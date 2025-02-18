// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/monitoring/dashboard/v1/table.proto

package dashboardpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A table that displays time series data.
type TimeSeriesTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The data displayed in this table.
	DataSets []*TimeSeriesTable_TableDataSet `protobuf:"bytes,1,rep,name=data_sets,json=dataSets,proto3" json:"data_sets,omitempty"`
}

func (x *TimeSeriesTable) Reset() {
	*x = TimeSeriesTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesTable) ProtoMessage() {}

func (x *TimeSeriesTable) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesTable.ProtoReflect.Descriptor instead.
func (*TimeSeriesTable) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0}
}

func (x *TimeSeriesTable) GetDataSets() []*TimeSeriesTable_TableDataSet {
	if x != nil {
		return x.DataSets
	}
	return nil
}

// Groups a time series query definition with table options.
type TimeSeriesTable_TableDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Fields for querying time series data from the
	// Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `protobuf:"bytes,1,opt,name=time_series_query,json=timeSeriesQuery,proto3" json:"time_series_query,omitempty"`
	// Optional. A template string for naming `TimeSeries` in the resulting data set.
	// This should be a string with interpolations of the form `${label_name}`,
	// which will resolve to the label's value i.e.
	// "${resource.labels.project_id}."
	TableTemplate string `protobuf:"bytes,2,opt,name=table_template,json=tableTemplate,proto3" json:"table_template,omitempty"`
	// Optional. The lower bound on data point frequency for this data set, implemented by
	// specifying the minimum alignment period to use in a time series query
	// For example, if the data is published once every 10 minutes, the
	// `min_alignment_period` should be at least 10 minutes. It would not
	// make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *durationpb.Duration `protobuf:"bytes,3,opt,name=min_alignment_period,json=minAlignmentPeriod,proto3" json:"min_alignment_period,omitempty"`
	// Optional. Table display options for configuring how the table is rendered.
	TableDisplayOptions *TableDisplayOptions `protobuf:"bytes,4,opt,name=table_display_options,json=tableDisplayOptions,proto3" json:"table_display_options,omitempty"`
}

func (x *TimeSeriesTable_TableDataSet) Reset() {
	*x = TimeSeriesTable_TableDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesTable_TableDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesTable_TableDataSet) ProtoMessage() {}

func (x *TimeSeriesTable_TableDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesTable_TableDataSet.ProtoReflect.Descriptor instead.
func (*TimeSeriesTable_TableDataSet) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TimeSeriesTable_TableDataSet) GetTimeSeriesQuery() *TimeSeriesQuery {
	if x != nil {
		return x.TimeSeriesQuery
	}
	return nil
}

func (x *TimeSeriesTable_TableDataSet) GetTableTemplate() string {
	if x != nil {
		return x.TableTemplate
	}
	return ""
}

func (x *TimeSeriesTable_TableDataSet) GetMinAlignmentPeriod() *durationpb.Duration {
	if x != nil {
		return x.MinAlignmentPeriod
	}
	return nil
}

func (x *TimeSeriesTable_TableDataSet) GetTableDisplayOptions() *TableDisplayOptions {
	if x != nil {
		return x.TableDisplayOptions
	}
	return nil
}

var File_google_monitoring_dashboard_v1_table_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_table_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x03, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x73, 0x1a, 0xdc, 0x02, 0x0a, 0x0c,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x60, 0x0a, 0x11,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a,
	0x0a, 0x0e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x14, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x41, 0x6c, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x6c, 0x0a, 0x15,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_dashboard_v1_table_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_table_proto_rawDescData = file_google_monitoring_dashboard_v1_table_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_table_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_table_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_table_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_table_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_table_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_monitoring_dashboard_v1_table_proto_goTypes = []interface{}{
	(*TimeSeriesTable)(nil),              // 0: google.monitoring.dashboard.v1.TimeSeriesTable
	(*TimeSeriesTable_TableDataSet)(nil), // 1: google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet
	(*TimeSeriesQuery)(nil),              // 2: google.monitoring.dashboard.v1.TimeSeriesQuery
	(*durationpb.Duration)(nil),          // 3: google.protobuf.Duration
	(*TableDisplayOptions)(nil),          // 4: google.monitoring.dashboard.v1.TableDisplayOptions
}
var file_google_monitoring_dashboard_v1_table_proto_depIdxs = []int32{
	1, // 0: google.monitoring.dashboard.v1.TimeSeriesTable.data_sets:type_name -> google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet
	2, // 1: google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.time_series_query:type_name -> google.monitoring.dashboard.v1.TimeSeriesQuery
	3, // 2: google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.min_alignment_period:type_name -> google.protobuf.Duration
	4, // 3: google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.table_display_options:type_name -> google.monitoring.dashboard.v1.TableDisplayOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_table_proto_init() }
func file_google_monitoring_dashboard_v1_table_proto_init() {
	if File_google_monitoring_dashboard_v1_table_proto != nil {
		return
	}
	file_google_monitoring_dashboard_v1_metrics_proto_init()
	file_google_monitoring_dashboard_v1_table_display_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesTable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_monitoring_dashboard_v1_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesTable_TableDataSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_monitoring_dashboard_v1_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_table_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_table_proto_depIdxs,
		MessageInfos:      file_google_monitoring_dashboard_v1_table_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_table_proto = out.File
	file_google_monitoring_dashboard_v1_table_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_table_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_table_proto_depIdxs = nil
}
