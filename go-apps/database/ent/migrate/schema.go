// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminColumns holds the columns for the "admin" table.
	AdminColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// AdminTable holds the schema information for the "admin" table.
	AdminTable = &schema.Table{
		Name:       "admin",
		Columns:    AdminColumns,
		PrimaryKey: []*schema.Column{AdminColumns[0]},
	}
	// GenerationColumns holds the columns for the "generation" table.
	GenerationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GenerationTable holds the schema information for the "generation" table.
	GenerationTable = &schema.Table{
		Name:       "generation",
		Columns:    GenerationColumns,
		PrimaryKey: []*schema.Column{GenerationColumns[0]},
	}
	// GenerationGColumns holds the columns for the "generation_g" table.
	GenerationGColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GenerationGTable holds the schema information for the "generation_g" table.
	GenerationGTable = &schema.Table{
		Name:       "generation_g",
		Columns:    GenerationGColumns,
		PrimaryKey: []*schema.Column{GenerationGColumns[0]},
	}
	// GenerationRealtimeColumns holds the columns for the "generation_realtime" table.
	GenerationRealtimeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GenerationRealtimeTable holds the schema information for the "generation_realtime" table.
	GenerationRealtimeTable = &schema.Table{
		Name:       "generation_realtime",
		Columns:    GenerationRealtimeColumns,
		PrimaryKey: []*schema.Column{GenerationRealtimeColumns[0]},
	}
	// ModelColumns holds the columns for the "model" table.
	ModelColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ModelTable holds the schema information for the "model" table.
	ModelTable = &schema.Table{
		Name:       "model",
		Columns:    ModelColumns,
		PrimaryKey: []*schema.Column{ModelColumns[0]},
	}
	// NegativePromptColumns holds the columns for the "negative_prompt" table.
	NegativePromptColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NegativePromptTable holds the schema information for the "negative_prompt" table.
	NegativePromptTable = &schema.Table{
		Name:       "negative_prompt",
		Columns:    NegativePromptColumns,
		PrimaryKey: []*schema.Column{NegativePromptColumns[0]},
	}
	// PromptColumns holds the columns for the "prompt" table.
	PromptColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PromptTable holds the schema information for the "prompt" table.
	PromptTable = &schema.Table{
		Name:       "prompt",
		Columns:    PromptColumns,
		PrimaryKey: []*schema.Column{PromptColumns[0]},
	}
	// SchedulerColumns holds the columns for the "scheduler" table.
	SchedulerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SchedulerTable holds the schema information for the "scheduler" table.
	SchedulerTable = &schema.Table{
		Name:       "scheduler",
		Columns:    SchedulerColumns,
		PrimaryKey: []*schema.Column{SchedulerColumns[0]},
	}
	// ServerColumns holds the columns for the "server" table.
	ServerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ServerTable holds the schema information for the "server" table.
	ServerTable = &schema.Table{
		Name:       "server",
		Columns:    ServerColumns,
		PrimaryKey: []*schema.Column{ServerColumns[0]},
	}
	// UpscaleColumns holds the columns for the "upscale" table.
	UpscaleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UpscaleTable holds the schema information for the "upscale" table.
	UpscaleTable = &schema.Table{
		Name:       "upscale",
		Columns:    UpscaleColumns,
		PrimaryKey: []*schema.Column{UpscaleColumns[0]},
	}
	// UpscaleRealtimeColumns holds the columns for the "upscale_realtime" table.
	UpscaleRealtimeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UpscaleRealtimeTable holds the schema information for the "upscale_realtime" table.
	UpscaleRealtimeTable = &schema.Table{
		Name:       "upscale_realtime",
		Columns:    UpscaleRealtimeColumns,
		PrimaryKey: []*schema.Column{UpscaleRealtimeColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminTable,
		GenerationTable,
		GenerationGTable,
		GenerationRealtimeTable,
		ModelTable,
		NegativePromptTable,
		PromptTable,
		SchedulerTable,
		ServerTable,
		UpscaleTable,
		UpscaleRealtimeTable,
		UserTable,
	}
)

func init() {
	AdminTable.Annotation = &entsql.Annotation{
		Table: "admin",
	}
	GenerationTable.Annotation = &entsql.Annotation{
		Table: "generation",
	}
	GenerationGTable.Annotation = &entsql.Annotation{
		Table: "generation_g",
	}
	GenerationRealtimeTable.Annotation = &entsql.Annotation{
		Table: "generation_realtime",
	}
	ModelTable.Annotation = &entsql.Annotation{
		Table: "model",
	}
	NegativePromptTable.Annotation = &entsql.Annotation{
		Table: "negative_prompt",
	}
	PromptTable.Annotation = &entsql.Annotation{
		Table: "prompt",
	}
	SchedulerTable.Annotation = &entsql.Annotation{
		Table: "scheduler",
	}
	ServerTable.Annotation = &entsql.Annotation{
		Table: "server",
	}
	UpscaleTable.Annotation = &entsql.Annotation{
		Table: "upscale",
	}
	UpscaleRealtimeTable.Annotation = &entsql.Annotation{
		Table: "upscale_realtime",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
