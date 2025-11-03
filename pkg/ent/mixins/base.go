package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// BaseMixin 定义所有表的通用基础字段
// 适用于 PostgreSQL 数据库
type BaseMixin struct {
	ent.Mixin
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		// 主键ID：雪花ID，可手动赋值
		field.Int64("id").
			Unique().
			Positive().
			Immutable().
			Comment("Primary Key | 主键ID，可自定义雪花ID"),
	}
}

// Edges of the BaseMixin.
func (BaseMixin) Edges() []ent.Edge {
	return nil
}

// Indexes of the BaseMixin.
func (BaseMixin) Indexes() []ent.Index {
	return nil
}

// Hooks of the BaseMixin.
func (BaseMixin) Hooks() []ent.Hook {
	return nil
}

// Interceptors of the BaseMixin.
func (BaseMixin) Interceptors() []ent.Interceptor {
	return nil
}

// Policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return nil
}

// Annotations of the BaseMixin.
func (BaseMixin) Annotations() []schema.Annotation {
	return nil
}
