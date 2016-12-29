package bireader

import (
	"encoding/binary"
	"io"
)

type endian struct {
	// r      io.Reader
	binary.ByteOrder
}

var BigEndian = endian{binary.BigEndian}
var LittleEndian = endian{binary.LittleEndian}

func (e *endian) ReadBool(r io.Reader) (bool, error) {
	var c uint8
	err := binary.Read(r, e, &c)
	return c != 0, err
}

func (e *endian) ReadUint8(r io.Reader) (uint8, error) {
	var c uint8
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadUint16(r io.Reader) (uint16, error) {
	var c uint16
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadUint32(r io.Reader) (uint32, error) {
	var c uint32
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadUint64(r io.Reader) (uint64, error) {
	var c uint64
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadInt8(r io.Reader) (int8, error) {
	var c int8
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadInt16(r io.Reader) (int16, error) {
	var c int16
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadInt32(r io.Reader) (int32, error) {
	var c int32
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadInt(r io.Reader) (int, error) {
	var c int
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadFloat32(r io.Reader) (float32, error) {
	var c float32
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadInt64(r io.Reader) (int64, error) {
	var c int64
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadString(r io.Reader) (string, error) {
	c := make([]byte, 0)
	var err error
	for {
		var b byte
		err = binary.Read(r, e, &b)
		if b == byte(0x0) || err != nil {
			break
		}
		c = append(c, b)
	}
	return string(c), err
}

func (e *endian) ReadStringWithLength(r io.Reader, length int) (string, error) {
	b := make([]byte, length)
	err := binary.Read(r, e, &b)
	return string(b), err
}

func (e *endian) ReadFloat32ArrayWithLength(r io.Reader, length int) ([]float32, error) {
	b := make([]float32, length)
	err := binary.Read(r, e, &b)
	return b, err
}

func (e *endian) ReadInt32ArrayWithLength(r io.Reader, length int) ([]int32, error) {
	b := make([]int32, length)
	err := binary.Read(r, e, &b)
	return b, err
}

func (e *endian) ReadUint8ArrayWithLength(r io.Reader, length int) ([]uint8, error) {
	b := make([]uint8, length)
	err := binary.Read(r, e, &b)
	return b, err
}

func (e *endian) ReadByte(r io.Reader) (byte, error) {
	var c byte
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) ReadBytes(r io.Reader, num int32) ([]byte, error) {
	c := make([]byte, num)
	err := binary.Read(r, e, &c)
	return c, err
}

func (e *endian) WriteBool(w io.Writer, b bool) error {
	var err error
	if b {
		_, err = w.Write([]byte{1})
	} else {
		_, err = w.Write([]byte{0})
	}
	return err
}
