package main

import (
	"reflect"
	"testing"
)

func TestSemantic_New(t *testing.T) {

	tests := []struct {
		value string
		want  *Semantic
		err   error
	}{
		{
			value: "1.0.0",
			want:  &Semantic{major: 1, minor: 0, patch: 0},
			err:   nil,
		},
		{
			value: "1.0.0-alpha.0",
			want:  &Semantic{major: 1, minor: 0, patch: 0, pre: "alpha.0", isPre: true},
			err:   nil,
		},
		{
			value: "x.0.0",
			want:  nil,
			err:   SemanticErrors.MajorVersionIsNotValid,
		},
		{
			value: "1.y.0",
			want:  nil,
			err:   SemanticErrors.MinorVersionIsNotValid,
		},
		{
			value: "1.0.z",
			want:  nil,
			err:   SemanticErrors.PatchVersionIsNotValid,
		},
		{
			value: "1.0.0.alpha.0",
			want:  nil,
			err:   SemanticErrors.PatchVersionIsNotValid,
		},
		{
			value: "10.0",
			want:  nil,
			err:   SemanticErrors.VersionIsNotValid,
		},
	}

	for _, test := range tests {
		if s, err := New(test.value); !reflect.DeepEqual(s, test.want) || !reflect.DeepEqual(err, test.err) {
			t.Errorf("New(%s) = %+v, want %+v, err: %+v", test.value, s, test.want, err)
		}
	}

}

func TestSemantic_Compare(t *testing.T) {

	tests := []struct {
		v1       string
		operator string
		v2       string
		want     bool
		err      error
	}{
		{
			v1:       "1.0.0",
			v2:       "1.0.0",
			operator: "==",
			want:     true,
			err:      nil,
		},
		{
			v1:       "1.0.0",
			v2:       "1.0.0",
			operator: ">=",
			want:     true,
			err:      nil,
		},
		{
			v1:       "1.0.1",
			v2:       "1.0.0",
			operator: ">",
			want:     true,
			err:      nil,
		},
		{
			v1:       "1.0.0",
			v2:       "1.0.1",
			operator: ">",
			want:     false,
			err:      nil,
		},
		{
			v1:       "1.0.0",
			v2:       "1.0.1",
			operator: "<",
			want:     true,
			err:      nil,
		},
		{
			v1:       "1.0.1",
			v2:       "1.0.1",
			operator: "<=",
			want:     true,
			err:      nil,
		},
	}

	for _, test := range tests {
		if result, err := Compare(test.v1, test.operator, test.v2); result != test.want || !reflect.DeepEqual(test.err, err) {
			t.Errorf("Compare(%s, %s, %s) = %+v, want %+v", test.v1, test.operator, test.v2, result, test.want)
		}
	}

}

func TestSemantic_Equal(t *testing.T) {

	tests := []struct {
		v1   Semantic
		v2   Semantic
		want bool
	}{
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0, pre: "alpha.0", isPre: true},
			v2:   Semantic{major: 1, minor: 0, patch: 0, pre: "alpha.0", isPre: true},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 1},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0, pre: "alpha.0", isPre: true},
			v2:   Semantic{major: 1, minor: 0, patch: 0, pre: "alpha.1", isPre: true},
			want: false,
		},
	}

	for _, test := range tests {
		if result := Equal(test.v1, test.v2); result != test.want {
			t.Errorf("Equal(%+v, %+v) = %+v, want %+v", test.v1, test.v2, result, test.want)
		}
	}

}

func TestSemantic_GreaterThan(t *testing.T) {

	tests := []struct {
		v1   Semantic
		v2   Semantic
		want bool
	}{
		{
			v1:   Semantic{major: 2, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 1, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 1},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 2, minor: 0, patch: 0},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 1, patch: 0},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 1},
			want: false,
		},
	}

	for _, test := range tests {
		if result := GreaterThan(test.v1, test.v2); result != test.want {
			t.Errorf("Equal(%+v, %+v) = %+v, want %+v", test.v1, test.v2, result, test.want)
		}
	}

}

func TestSemantic_LessThan(t *testing.T) {

	tests := []struct {
		v1   Semantic
		v2   Semantic
		want bool
	}{
		{
			v1:   Semantic{major: 2, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 1, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 1},
			v2:   Semantic{major: 1, minor: 0, patch: 0},
			want: false,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 2, minor: 0, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 1, patch: 0},
			want: true,
		},
		{
			v1:   Semantic{major: 1, minor: 0, patch: 0},
			v2:   Semantic{major: 1, minor: 0, patch: 1},
			want: true,
		},
	}

	for _, test := range tests {
		if result := LessThan(test.v1, test.v2); result != test.want {
			t.Errorf("Equal(%+v, %+v) = %+v, want %+v", test.v1, test.v2, result, test.want)
		}
	}

}

func TestSemantic_String(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.0",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 0},
			want:     "2.1.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 1},
			want:     "2.0.1",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "RELEASE.0", isPre: true},
			want:     "2.0.0-RELEASE.0",
		},
	}

	for _, test := range tests {
		if result := test.semantic.String(); result != test.want {
			t.Errorf("String() = %s, want %s", result, test.want)
		}
	}

}

func TestSemantic_Inc(t *testing.T) {

	tests := []struct {
		major    bool
		minor    bool
		patch    bool
		release  bool
		beta     bool
		alpha    bool
		semantic Semantic
		want     string
	}{
		{
			major:    true,
			semantic: Semantic{major: 1, minor: 1, patch: 1, pre: "RELEASE.0", isPre: true},
			want:     "2.0.0",
		},
		{
			minor:    true,
			semantic: Semantic{major: 1, minor: 0, patch: 1, pre: "RELEASE.0", isPre: true},
			want:     "1.1.0",
		},
		{
			patch:    true,
			semantic: Semantic{major: 1, minor: 0, patch: 0, pre: "RELEASE.0", isPre: true},
			want:     "1.0.1",
		},
		{
			release:  true,
			semantic: Semantic{major: 1, minor: 0, patch: 0, pre: "RELEASE.0", isPre: true},
			want:     "1.0.0-RELEASE.1",
		},
		{
			beta:     true,
			semantic: Semantic{major: 1, minor: 0, patch: 0, pre: "BETA.0", isPre: true},
			want:     "1.0.0-BETA.1",
		},
		{
			alpha:    true,
			semantic: Semantic{major: 1, minor: 0, patch: 0, pre: "ALPHA.0", isPre: true},
			want:     "1.0.0-ALPHA.1",
		},
		{
			release:  true,
			semantic: Semantic{major: 1, minor: 0, patch: 0},
			want:     "1.0.0-RELEASE.0",
		},
		{
			beta:     true,
			semantic: Semantic{major: 1, minor: 0, patch: 0},
			want:     "1.0.0-BETA.0",
		},
		{
			alpha:    true,
			semantic: Semantic{major: 1, minor: 0, patch: 0},
			want:     "1.0.0-ALPHA.0",
		},
	}

	for _, test := range tests {
		if test.semantic.Inc(test.major, test.minor, test.patch, test.release, test.beta, test.alpha); test.semantic.String() != test.want {
			t.Errorf("Inc(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncMajor(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "3.0.0",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 0},
			want:     "3.0.0",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 1},
			want:     "3.0.0",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 1, pre: "RELEASE.0", isPre: true},
			want:     "3.0.0",
		},
	}

	for _, test := range tests {
		if test.semantic.IncMajor(); test.semantic.String() != test.want {
			t.Errorf("IncMajor(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncMinor(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.1.0",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 0},
			want:     "2.2.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 1},
			want:     "2.1.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 1, pre: "RELEASE.0", isPre: true},
			want:     "2.1.0",
		},
	}

	for _, test := range tests {
		if test.semantic.IncMinor(); test.semantic.String() != test.want {
			t.Errorf("IncMinor(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncPatch(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.1",
		},
		{
			semantic: Semantic{major: 2, minor: 1, patch: 0},
			want:     "2.1.1",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 1},
			want:     "2.0.2",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 1, pre: "RELEASE.0", isPre: true},
			want:     "2.0.2",
		},
	}

	for _, test := range tests {
		if test.semantic.IncPatch(); test.semantic.String() != test.want {
			t.Errorf("IncPatch(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncPre(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.0-pre.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "pre", isPre: true},
			want:     "2.0.0-pre.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "pre.0", isPre: true},
			want:     "2.0.0-pre.1",
		},
	}

	for _, test := range tests {
		if test.semantic.incPre("pre"); test.semantic.String() != test.want {
			t.Errorf("incPre(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncRelease(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.0-RELEASE.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "TEST", isPre: true},
			want:     "2.0.0-RELEASE.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "RELEASE.0", isPre: true},
			want:     "2.0.0-RELEASE.1",
		},
	}

	for _, test := range tests {
		if test.semantic.IncRelease(); test.semantic.String() != test.want {
			t.Errorf("IncRelease(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncAlpha(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.0-ALPHA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "ALPHA", isPre: true},
			want:     "2.0.0-ALPHA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "TEST", isPre: true},
			want:     "2.0.0-ALPHA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "ALPHA.0", isPre: true},
			want:     "2.0.0-ALPHA.1",
		},
	}

	for _, test := range tests {
		if test.semantic.IncAlpha(); test.semantic.String() != test.want {
			t.Errorf("IncAlpha(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_IncBeta(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     string
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     "2.0.0-BETA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "BETA", isPre: true},
			want:     "2.0.0-BETA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "TEST", isPre: true},
			want:     "2.0.0-BETA.0",
		},
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0, pre: "BETA.0", isPre: true},
			want:     "2.0.0-BETA.1",
		},
	}

	for _, test := range tests {
		if test.semantic.IncBeta(); test.semantic.String() != test.want {
			t.Errorf("IncBeta(%+v) = %s, want %s", test.semantic, test.semantic.String(), test.want)
		}
	}

}

func TestSemantic_Value(t *testing.T) {

	tests := []struct {
		semantic Semantic
		want     int
	}{
		{
			semantic: Semantic{major: 2, minor: 0, patch: 0},
			want:     200,
		},
		{
			semantic: Semantic{major: 1, minor: 1, patch: 1},
			want:     111,
		},
		{
			semantic: Semantic{major: 0, minor: 1, patch: 1},
			want:     11,
		},
		{
			semantic: Semantic{major: 0, minor: 0, patch: 1},
			want:     1,
		},
	}

	for _, test := range tests {
		if test.semantic.Value(); test.semantic.Value() != test.want {
			t.Errorf("Value(%+v) = %d, want %d", test.semantic, test.semantic.Value(), test.want)
		}
	}

}
