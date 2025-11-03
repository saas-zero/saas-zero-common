package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CreatedMixin 定义所有表的通用基础字段
// 适用于 PostgreSQL 数据库
type CreatedMixin struct {
	ent.Mixin
}

func (CreatedMixin) Fields() []ent.Field {
	return []ent.Field{
		// 时间字段
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Create Time | 创建时间"),

		// 操作人 ID
		field.Int64("created_id").
			Positive().
			Immutable().
			Comment("Creator ID | 创建人ID"),

		// 操作人名称
		field.String("created_by").
			NotEmpty().
			MaxLen(64).
			Immutable().
			Comment("Creator Name | 创建人名称"),
	}
}

// Edges of the CreatedMixin.
func (CreatedMixin) Edges() []ent.Edge {
	return nil
}

// Indexes of the CreatedMixin.
func (CreatedMixin) Indexes() []ent.Index {
	return nil
}

// Hooks of the CreatedMixin.
func (CreatedMixin) Hooks() []ent.Hook {
	return nil
}

// Interceptors of the CreatedMixin.
func (CreatedMixin) Interceptors() []ent.Interceptor {
	return nil
}

// Policy of the CreatedMixin.
func (CreatedMixin) Policy() ent.Policy {
	return nil
}

// Annotations of the CreatedMixin.
func (CreatedMixin) Annotations() []schema.Annotation {
	return nil
}
