// Code generated by fastssz. DO NOT EDIT.
// Hash: abf19c47c7e904a319d06970c1195ce4f2d1fd46ca45c01aee7a20b8073ed87b
// Version: 0.1.2
package p2p

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Checkpoint object
func (c *Checkpoint) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Checkpoint object to a target array
func (c *Checkpoint) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, c.Epoch)

	// Field (1) 'Root'
	dst = append(dst, c.Root[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Checkpoint object
func (c *Checkpoint) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 40 {
		return ssz.ErrSize
	}

	// Field (0) 'Epoch'
	c.Epoch = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Root'
	copy(c.Root[:], buf[8:40])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Checkpoint object
func (c *Checkpoint) SizeSSZ() (size int) {
	size = 40
	return
}

// HashTreeRoot ssz hashes the Checkpoint object
func (c *Checkpoint) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Checkpoint object with a hasher
func (c *Checkpoint) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Epoch'
	hh.PutUint64(c.Epoch)

	// Field (1) 'Root'
	hh.PutBytes(c.Root[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Checkpoint object
func (c *Checkpoint) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}

// MarshalSSZ ssz marshals the ENRForkID object
func (e *ENRForkID) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ENRForkID object to a target array
func (e *ENRForkID) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentForkDigest'
	if size := len(e.CurrentForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("ENRForkID.CurrentForkDigest", size, 4)
		return
	}
	dst = append(dst, e.CurrentForkDigest...)

	// Field (1) 'NextForkVersion'
	if size := len(e.NextForkVersion); size != 4 {
		err = ssz.ErrBytesLengthFn("ENRForkID.NextForkVersion", size, 4)
		return
	}
	dst = append(dst, e.NextForkVersion...)

	// Field (2) 'NextForkEpoch'
	dst = ssz.MarshalUint64(dst, uint64(e.NextForkEpoch))

	return
}

// UnmarshalSSZ ssz unmarshals the ENRForkID object
func (e *ENRForkID) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentForkDigest'
	if cap(e.CurrentForkDigest) == 0 {
		e.CurrentForkDigest = make([]byte, 0, len(buf[0:4]))
	}
	e.CurrentForkDigest = append(e.CurrentForkDigest, buf[0:4]...)

	// Field (1) 'NextForkVersion'
	if cap(e.NextForkVersion) == 0 {
		e.NextForkVersion = make([]byte, 0, len(buf[4:8]))
	}
	e.NextForkVersion = append(e.NextForkVersion, buf[4:8]...)

	// Field (2) 'NextForkEpoch'
	e.NextForkEpoch = Epoch(ssz.UnmarshallUint64(buf[8:16]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ENRForkID object
func (e *ENRForkID) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the ENRForkID object
func (e *ENRForkID) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ENRForkID object with a hasher
func (e *ENRForkID) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentForkDigest'
	if size := len(e.CurrentForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("ENRForkID.CurrentForkDigest", size, 4)
		return
	}
	hh.PutBytes(e.CurrentForkDigest)

	// Field (1) 'NextForkVersion'
	if size := len(e.NextForkVersion); size != 4 {
		err = ssz.ErrBytesLengthFn("ENRForkID.NextForkVersion", size, 4)
		return
	}
	hh.PutBytes(e.NextForkVersion)

	// Field (2) 'NextForkEpoch'
	hh.PutUint64(uint64(e.NextForkEpoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ENRForkID object
func (e *ENRForkID) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ForkData object
func (f *ForkData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(f)
}

// MarshalSSZ ssz marshals the Goodbye object
func (g *Goodbye) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(g)
}

// MarshalSSZTo ssz marshals the Goodbye object to a target array
func (g *Goodbye) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Reason'
	dst = ssz.MarshalUint64(dst, g.Reason)

	return
}

// UnmarshalSSZ ssz unmarshals the Goodbye object
func (g *Goodbye) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 8 {
		return ssz.ErrSize
	}

	// Field (0) 'Reason'
	g.Reason = ssz.UnmarshallUint64(buf[0:8])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Goodbye object
func (g *Goodbye) SizeSSZ() (size int) {
	size = 8
	return
}

// HashTreeRoot ssz hashes the Goodbye object
func (g *Goodbye) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(g)
}

// HashTreeRootWith ssz hashes the Goodbye object with a hasher
func (g *Goodbye) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Reason'
	hh.PutUint64(g.Reason)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Goodbye object
func (g *Goodbye) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(g)
}

// MarshalSSZ ssz marshals the Ping object
func (p *Ping) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the Ping object to a target array
func (p *Ping) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Id'
	dst = ssz.MarshalUint64(dst, p.Id)

	// Field (1) 'Syncnets'
	dst = ssz.MarshalUint64(dst, uint64(p.Syncnets))

	return
}

// UnmarshalSSZ ssz unmarshals the Ping object
func (p *Ping) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Id'
	p.Id = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Syncnets'
	p.Syncnets = Bitvector64(ssz.UnmarshallUint64(buf[8:16]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Ping object
func (p *Ping) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the Ping object
func (p *Ping) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the Ping object with a hasher
func (p *Ping) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Id'
	hh.PutUint64(p.Id)

	// Field (1) 'Syncnets'
	hh.PutUint64(uint64(p.Syncnets))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Ping object
func (p *Ping) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the SingleRoot object
func (s *SingleRoot) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SingleRoot object to a target array
func (s *SingleRoot) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	dst = append(dst, s.Root[:]...)

	// Field (1) 'BodyRoot'
	dst = append(dst, s.BodyRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SingleRoot object
func (s *SingleRoot) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	copy(s.Root[:], buf[0:32])

	// Field (1) 'BodyRoot'
	copy(s.BodyRoot[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SingleRoot object
func (s *SingleRoot) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SingleRoot object
func (s *SingleRoot) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SingleRoot object with a hasher
func (s *SingleRoot) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	hh.PutBytes(s.Root[:])

	// Field (1) 'BodyRoot'
	hh.PutBytes(s.BodyRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SingleRoot object
func (s *SingleRoot) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the Status object
func (s *Status) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Status object to a target array
func (s *Status) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ForkDigest'
	if size := len(s.ForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("Status.ForkDigest", size, 4)
		return
	}
	dst = append(dst, s.ForkDigest...)

	// Field (1) 'FinalizedRoot'
	if size := len(s.FinalizedRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("Status.FinalizedRoot", size, 32)
		return
	}
	dst = append(dst, s.FinalizedRoot...)

	// Field (2) 'FinalizedEpoch'
	dst = ssz.MarshalUint64(dst, uint64(s.FinalizedEpoch))

	// Field (3) 'HeadRoot'
	if size := len(s.HeadRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("Status.HeadRoot", size, 32)
		return
	}
	dst = append(dst, s.HeadRoot...)

	// Field (4) 'HeadSlot'
	dst = ssz.MarshalUint64(dst, uint64(s.HeadSlot))

	return
}

// UnmarshalSSZ ssz unmarshals the Status object
func (s *Status) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 84 {
		return ssz.ErrSize
	}

	// Field (0) 'ForkDigest'
	if cap(s.ForkDigest) == 0 {
		s.ForkDigest = make([]byte, 0, len(buf[0:4]))
	}
	s.ForkDigest = append(s.ForkDigest, buf[0:4]...)

	// Field (1) 'FinalizedRoot'
	if cap(s.FinalizedRoot) == 0 {
		s.FinalizedRoot = make([]byte, 0, len(buf[4:36]))
	}
	s.FinalizedRoot = append(s.FinalizedRoot, buf[4:36]...)

	// Field (2) 'FinalizedEpoch'
	s.FinalizedEpoch = Epoch(ssz.UnmarshallUint64(buf[36:44]))

	// Field (3) 'HeadRoot'
	if cap(s.HeadRoot) == 0 {
		s.HeadRoot = make([]byte, 0, len(buf[44:76]))
	}
	s.HeadRoot = append(s.HeadRoot, buf[44:76]...)

	// Field (4) 'HeadSlot'
	s.HeadSlot = Slot(ssz.UnmarshallUint64(buf[76:84]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Status object
func (s *Status) SizeSSZ() (size int) {
	size = 84
	return
}

// HashTreeRoot ssz hashes the Status object
func (s *Status) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Status object with a hasher
func (s *Status) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ForkDigest'
	if size := len(s.ForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("Status.ForkDigest", size, 4)
		return
	}
	hh.PutBytes(s.ForkDigest)

	// Field (1) 'FinalizedRoot'
	if size := len(s.FinalizedRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("Status.FinalizedRoot", size, 32)
		return
	}
	hh.PutBytes(s.FinalizedRoot)

	// Field (2) 'FinalizedEpoch'
	hh.PutUint64(uint64(s.FinalizedEpoch))

	// Field (3) 'HeadRoot'
	if size := len(s.HeadRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("Status.HeadRoot", size, 32)
		return
	}
	hh.PutBytes(s.HeadRoot)

	// Field (4) 'HeadSlot'
	hh.PutUint64(uint64(s.HeadSlot))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Status object
func (s *Status) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}