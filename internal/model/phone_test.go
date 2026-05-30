package model

import (
	"testing"
)

func strPtr(s string) *string {
	return &s
}

func uint8Ptr(n uint8) *uint8 {
	return &n
}

func TestHrefValue_FullUSNumber(t *testing.T) {
	p := &PhoneNumber{
		CountryCode: uint8Ptr(1),
		AreaCode:    512,
		Prefix:      555,
		Subscriber:  1234,
	}
	got := p.HrefValue()
	want := "tel:+15125551234"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHrefValue_NoCountryCode(t *testing.T) {
	p := &PhoneNumber{
		AreaCode:   512,
		Prefix:     555,
		Subscriber: 1234,
	}
	got := p.HrefValue()
	want := "tel:5125551234"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHrefValue_WithExitCode(t *testing.T) {
	p := &PhoneNumber{
		ExitCode:    strPtr("011"),
		CountryCode: uint8Ptr(44),
		AreaCode:    20,
		Prefix:      555,
		Subscriber:  1234,
	}
	got := p.HrefValue()
	want := "tel:011+44205551234"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHrefValue_LeadingZeros(t *testing.T) {
	p := &PhoneNumber{
		CountryCode: uint8Ptr(1),
		AreaCode:    512,
		Prefix:      555,
		Subscriber:  99,
	}
	got := p.HrefValue()
	want := "tel:+15125550099"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHrefValue_AllOptionalNil(t *testing.T) {
	p := &PhoneNumber{
		AreaCode:   512,
		Prefix:     555,
		Subscriber: 1234,
	}
	got := p.HrefValue()
	want := "tel:5125551234"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
