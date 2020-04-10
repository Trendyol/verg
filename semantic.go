package main

import (
	"strconv"
	"strings"
)

type Semantic struct {
	major int
	minor int
	patch int
	pre   string
	isPre bool
}

func New(v string) (*Semantic, error) {

	items := strings.SplitN(v, ".", 3)

	if len(items) < 3 {
		return nil, SemanticErrors.VersionIsNotValid
	}

	major, majorErr := strconv.ParseInt(items[0], 10, 32)

	if majorErr != nil {
		return nil, SemanticErrors.MajorVersionIsNotValid
	}

	minor, minorErr := strconv.ParseInt(items[1], 10, 32)

	if minorErr != nil {
		return nil, SemanticErrors.MinorVersionIsNotValid
	}

	patchItems := strings.SplitN(items[2], "-", 2)

	patch, patchErr := strconv.ParseInt(patchItems[0], 10, 32)

	if patchErr != nil {
		return nil, SemanticErrors.PatchVersionIsNotValid
	}

	semantic := &Semantic{
		major: int(major),
		minor: int(minor),
		patch: int(patch),
	}

	if len(patchItems) > 1 {
		semantic.pre = patchItems[1]
		semantic.isPre = true
	}

	return semantic, nil
}

func Compare(version1, operator, version2 string) (bool, error) {
	v1, err := New(version1)

	if err != nil {
		return false, err
	}

	v2, err := New(version2)

	if err != nil {
		return false, err
	}

	switch operator {
	case ">":
		return GreaterThan(*v1, *v2), nil
	case ">=":
		return Equal(*v1, *v2) || GreaterThan(*v1, *v2), nil
	case "<":
		return LessThan(*v1, *v2), nil
	case "<=":
		return Equal(*v1, *v2) || LessThan(*v1, *v2), nil
	case "==":
		return Equal(*v1, *v2), nil
	}

	return false, nil
}

func Equal(v1, v2 Semantic) bool {
	return v1.Value() == v2.Value() && v1.pre == v2.pre
}

func GreaterThan(v1, v2 Semantic) bool {
	return v1.Value() > v2.Value()
}

func LessThan(v1, v2 Semantic) bool {
	return v1.Value() < v2.Value()
}

func (s Semantic) String() string {

	items := make([]string, 3, 4)

	items[0] = strconv.Itoa(s.major)
	items[1] = strconv.Itoa(s.minor)
	items[2] = strconv.Itoa(s.patch)

	if s.isPre {
		items[2] += "-" + s.pre
	}

	return strings.Join(items, ".")
}

func (s *Semantic) Inc(major, minor, patch, release, beta, alpha bool) {
	if major {
		s.IncMajor()
	}

	if minor {
		s.IncMinor()
	}

	if patch {
		s.IncPatch()
	}

	if release {
		s.IncRelease()
	}

	if beta {
		s.IncBeta()
	}

	if alpha {
		s.IncAlpha()
	}
}

func (s *Semantic) IncMajor() {
	s.major++
	s.minor = 0
	s.patch = 0
	s.pre = ""
	s.isPre = false
}

func (s *Semantic) IncMinor() {
	s.minor++
	s.patch = 0
	s.pre = ""
	s.isPre = false
}

func (s *Semantic) IncPatch() {
	s.patch++
	s.pre = ""
	s.isPre = false
}

func (s *Semantic) incPre(name string) {
	if s.isPre && strings.HasPrefix(s.pre, name) {
		items := strings.SplitN(s.pre, ".", 2)

		if len(items) == 1 {
			items = append(items, "0")
		} else {
			v, _ := strconv.ParseInt(items[1], 10, 32)

			v++

			items[1] = strconv.Itoa(int(v))
		}

		s.pre = strings.Join(items, ".")
	} else {
		s.isPre = true
		s.pre = name + ".0"
	}

}

func (s *Semantic) IncRelease() {
	s.incPre("RELEASE")
}

func (s *Semantic) IncBeta() {
	s.incPre("BETA")
}

func (s *Semantic) IncAlpha() {
	s.incPre("ALPHA")
}

func (s *Semantic) Value() int {
	return (s.major * 100) + (s.minor * 10) + s.patch
}
