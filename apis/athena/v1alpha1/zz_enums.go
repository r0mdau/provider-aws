/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type ColumnNullable string

const (
	ColumnNullable_NOT_NULL ColumnNullable = "NOT_NULL"
	ColumnNullable_NULLABLE ColumnNullable = "NULLABLE"
	ColumnNullable_UNKNOWN  ColumnNullable = "UNKNOWN"
)

type DataCatalogType string

const (
	DataCatalogType_LAMBDA DataCatalogType = "LAMBDA"
	DataCatalogType_GLUE   DataCatalogType = "GLUE"
	DataCatalogType_HIVE   DataCatalogType = "HIVE"
)

type EncryptionOption string

const (
	EncryptionOption_SSE_S3  EncryptionOption = "SSE_S3"
	EncryptionOption_SSE_KMS EncryptionOption = "SSE_KMS"
	EncryptionOption_CSE_KMS EncryptionOption = "CSE_KMS"
)

type QueryExecutionState string

const (
	QueryExecutionState_QUEUED    QueryExecutionState = "QUEUED"
	QueryExecutionState_RUNNING   QueryExecutionState = "RUNNING"
	QueryExecutionState_SUCCEEDED QueryExecutionState = "SUCCEEDED"
	QueryExecutionState_FAILED    QueryExecutionState = "FAILED"
	QueryExecutionState_CANCELLED QueryExecutionState = "CANCELLED"
)

type StatementType string

const (
	StatementType_DDL     StatementType = "DDL"
	StatementType_DML     StatementType = "DML"
	StatementType_UTILITY StatementType = "UTILITY"
)

type ThrottleReason string

const (
	ThrottleReason_CONCURRENT_QUERY_LIMIT_EXCEEDED ThrottleReason = "CONCURRENT_QUERY_LIMIT_EXCEEDED"
)

type WorkGroupState string

const (
	WorkGroupState_ENABLED  WorkGroupState = "ENABLED"
	WorkGroupState_DISABLED WorkGroupState = "DISABLED"
)