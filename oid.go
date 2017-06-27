package oid

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

var (
	machineID  []byte
	processID  int
	oidCounter uint32
)

// readRandomUint32 returns a random objectIdCounter.
func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot read random object id: %v", err))
	}
	return uint32((uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24))
}

// readMachineId generates and returns a machine id.
// If this function fails to get the hostname it will cause a runtime error.
func readMachineID() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

func init() {
	machineID = readMachineID()
	processID = os.Getpid()
	oidCounter = readRandomUint32()
}

// OID represents unique identifier id for each object in mongodb.
type OID [12]byte

// NewOID returns a new object id.
func NewOID() OID {
	oid := OID{}
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(oid[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	oid[4] = machineID[0]
	oid[5] = machineID[1]
	oid[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	oid[7] = byte(processID >> 8)
	oid[8] = byte(processID)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&oidCounter, 1)
	oid[9] = byte(i >> 16)
	oid[10] = byte(i >> 8)
	oid[11] = byte(i)
	return oid
}

// Equal returns true if o1 and o2 equals, otherwise returns false.
func Equal(o1, o2 OID) bool {
	return bytes.Equal(o1[:], o2[:])
}

// String returns the string representation of object id.
func (o OID) String() string {
	return hex.EncodeToString(o[:])
}

// Bytes returns the bytes slice representation of object id.
func (o OID) Bytes() []byte {
	return o[:]
}

// Timestamp returns the timestamp part of object id.
func (o OID) Timestamp() int64 {
	tsBytes := []byte(o[0:4])
	ts := int64(binary.BigEndian.Uint32(tsBytes))
	return ts
}
