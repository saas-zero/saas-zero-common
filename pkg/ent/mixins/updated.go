package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UpdatedMixin 定义所有表的通用基础字段
// 适用于 PostgreSQL 数据库
type UpdatedMixin struct {
	ent.Mixin
}

func (UpdatedMixin) Fields() []ent.Field {
	return []ent.Field{
		// 时间字段
		field.Time("updated_at").
			Default(time.Now).
			Comment("updated Time | 更新时间"),

		// 操作人 ID
		field.Int64("updated_id").
			Positive().
			Comment("updated ID | 更新人ID"),

		// 操作人名称
		field.String("updated_by").
			NotEmpty().
			MaxLen(64).
			Comment("updated Name | 更新人名称"),
	}
}

// Edges of the UpdatedMixin.
func (UpdatedMixin) Edges() []ent.Edge {
	return nil
}

// Indexes of the UpdatedMixin.
func (UpdatedMixin) Indexes() []ent.Index {
	return nil
}

// Hooks of the UpdatedMixin.
func (UpdatedMixin) Hooks() []ent.Hook {
	return nil
}

// Interceptors of the UpdatedMixin.
func (UpdatedMixin) Interceptors() []ent.Interceptor {
	return nil
}

// Policy of the UpdatedMixin.
func (UpdatedMixin) Policy() ent.Policy {
	return nil
}

// Annotations of the UpdatedMixin.
func (UpdatedMixin) Annotations() []schema.Annotation {
	return nil
}
