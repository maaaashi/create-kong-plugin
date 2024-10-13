package handler

import "testing"

func TestCreatePluginTemplate(t *testing.T) {
	type args struct {
		pluginName string
		language   string
		mkdir      bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePluginTemplate(tt.args.pluginName, tt.args.language, tt.args.mkdir)
		})
	}
}
