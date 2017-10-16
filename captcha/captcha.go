// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package captcha implements generation and verification of image and audio
// CAPTCHAs.
//
// A captcha solution is the sequence of digits 0-9 with the defined length.
// There are two captcha representations: image and audio.
//
// An image representation is a PNG-encoded image with the solution printed on
// it in such a way that makes it hard for computers to solve it using OCR.
//
// An audio representation is a WAVE-encoded (8 kHz unsigned 8-bit) sound with
// the spoken solution (currently in English, Russian, Chinese, and Japanese).
// To make it hard for computers to solve audio captcha, the voice that
// pronounces numbers has random speed and pitch, and there is a randomly
// generated background noise mixed into the sound.
//
// This package doesn't require external files or libraries to generate captcha
// representations; it is self-contained.
//
// To make captchas one-time, the package includes a memory storage that stores
// captcha ids, their solutions, and expiration time. Used captchas are removed
// from the store immediately after calling Verify or VerifyString, while
// unused captchas (user loaded a page with captcha, but didn't submit the
// form) are collected automatically after the predefined expiration time.
// Developers can also provide custom store (for example, which saves captcha
// ids and solutions in database) by implementing Store interface and
// registering the object with SetCustomStore.
//
// Captchas are created by calling New, which returns the captcha id.  Their
// representations, though, are created on-the-fly by calling WriteImage or
// WriteAudio functions. Created representations are not stored anywhere, but
// subsequent calls to these functions with the same id will write the same
// captcha solution. Reload function will create a new different solution for
// the provided captcha, allowing users to "reload" captcha if they can't solve
// the displayed one without reloading the whole page.  Verify and VerifyString
// are used to verify that the given solution is the right one for the given
// captcha id.
//
// Server provides an http.Handler which can serve image and audio
// representations of captchas automatically from the URL. It can also be used
// to reload captchas.  Refer to Server function documentation for details, or
// take a look at the example in "capexample" subdirectory.
package captcha

import (
	"errors"
	"io"
	"time"
	"fmt"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)

var (
	ErrNotFound = errors.New("captcha: id not found")
)

// New creates a new captcha with the standard length, saves it in the internal
// storage and returns its id.
func New() string {
	return NewLen(DefaultLen)
}

// NewLen is just like New, but accepts length of a captcha solution as the
// argument.
func NewLen(length int) (code string) {
	codeByte:=RandomDigits(length)
	for _, value := range codeByte {
		code+=fmt.Sprintf("%d",value)
	}
	return
}

func codeToByte(code string) []byte {
	ns := make([]byte, len(code))
	for i := range ns {
		d := code[i]
		switch {
		case '0' <= d && d <= '9':
			ns[i] = d - '0'
		case d == ' ' || d == ',':
			// ignore
		default:
		}
		}
		return ns
}

// WriteImage writes PNG-encoded image representation of the captcha with the
// given code. The image will have the given width and height.
func WriteImage(w io.Writer,code string, width, height int) error {
	_, err := NewImage(codeToByte(code), width, height).WriteTo(w)
	return err
}

// WriteAudio writes WAV-encoded audio representation of the captcha with the
// given id and the given language. If there are no sounds for the given
// language, English is used.
func WriteAudio(w io.Writer,code string, lang string) error {
	_, err := NewAudio(codeToByte(code), lang).WriteTo(w)
	return err
}