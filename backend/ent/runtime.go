// Code generated by entc, DO NOT EDIT.

package ent

import (
	"matsuokashuhei/go-todo/ent/schema"
	"matsuokashuhei/go-todo/ent/tag"
	"matsuokashuhei/go-todo/ent/task"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tagMixin := schema.Tag{}.Mixin()
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreateTime is the schema descriptor for create_time field.
	tagDescCreateTime := tagMixinFields0[0].Descriptor()
	// tag.DefaultCreateTime holds the default value on creation for the create_time field.
	tag.DefaultCreateTime = tagDescCreateTime.Default.(func() time.Time)
	// tagDescUpdateTime is the schema descriptor for update_time field.
	tagDescUpdateTime := tagMixinFields0[1].Descriptor()
	// tag.DefaultUpdateTime holds the default value on creation for the update_time field.
	tag.DefaultUpdateTime = tagDescUpdateTime.Default.(func() time.Time)
	// tag.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	tag.UpdateDefaultUpdateTime = tagDescUpdateTime.UpdateDefault.(func() time.Time)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreateTime is the schema descriptor for create_time field.
	taskDescCreateTime := taskMixinFields0[0].Descriptor()
	// task.DefaultCreateTime holds the default value on creation for the create_time field.
	task.DefaultCreateTime = taskDescCreateTime.Default.(func() time.Time)
	// taskDescUpdateTime is the schema descriptor for update_time field.
	taskDescUpdateTime := taskMixinFields0[1].Descriptor()
	// task.DefaultUpdateTime holds the default value on creation for the update_time field.
	task.DefaultUpdateTime = taskDescUpdateTime.Default.(func() time.Time)
	// task.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	task.UpdateDefaultUpdateTime = taskDescUpdateTime.UpdateDefault.(func() time.Time)
}