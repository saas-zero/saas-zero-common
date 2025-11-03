package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DeletedMixin 定义所有表的通用基础字段
// 适用于 PostgreSQL 数据库
type DeletedMixin struct {
	ent.Mixin
}

func (DeletedMixin) Fields() []ent.Field {
	return []ent.Field{
		// 时间字段
		field.Time("deleted_at").
			Optional().
			Comment("deleted Time | 删除时间"),

		// 操作人 ID
		field.Int64("deleted_id").
			Optional().
			Comment("deleted ID | 删除人ID，逻辑删除用"),

		// 操作人名称
		field.String("deleted_by").
			Optional().
			MaxLen(64).
			Comment("deleted Name | 删除人名称"),
	}
}

// Edges of the DeletedMixin.
func (DeletedMixin) Edges() []ent.Edge {
	return nil
}

// Indexes of the DeletedMixin.
func (DeletedMixin) Indexes() []ent.Index {
	return nil
}

// Hooks of the DeletedMixin.
func (DeletedMixin) Hooks() []ent.Hook {
	return nil
}

// Interceptors of the DeletedMixin.
func (DeletedMixin) Interceptors() []ent.Interceptor {
	return nil
}

// Policy of the DeletedMixin.
func (DeletedMixin) Policy() ent.Policy {
	return nil
}

// Annotations of the DeletedMixin.
func (DeletedMixin) Annotations() []schema.Annotation {
	return nil
}
