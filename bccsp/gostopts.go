package bccsp

const (
	// GOSTR3410_2012_256
	GOSTR3410_2012_256 = "GOSTR3410_2012_256"
	// GOSTR3410_2012_512
	GOSTR3410_2012_512 = "GOSTR3410_2012_512"

	// GOSTR3411_2012_256
	GOSTR3411_2012_256 = "GOSTR3411_2012_256"
	// GOSTR3411_2012_512
	GOSTR3411_2012_512 = "GOSTR3411_2012_512"

	// GOST28147
	GOST28147 = "GOST28147"
)

// GOST28147KeyGenOpts contains options GOST28147KeyGenOpts
type GOST28147KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *GOST28147KeyGenOpts) Algorithm() string {
	return GOST28147
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *GOST28147KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// GOSTR3410_2012_256Opts contains GOSTR3410_2012_256Opts
type GOSTR3410_2012_256Opts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *GOSTR3410_2012_256Opts) Algorithm() string {
	return GOSTR3410_2012_256
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *GOSTR3410_2012_256Opts) Ephemeral() bool {
	return opts.Temporary
}

// GOSTR3410_2012_512Opts contains options for GOSTR3410_2012_512Opts.
type GOSTR3410_2012_512Opts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *GOSTR3410_2012_512Opts) Algorithm() string {
	return GOSTR3410_2012_256
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *GOSTR3410_2012_512Opts) Ephemeral() bool {
	return opts.Temporary
}

// GOSTR3411_2012_256Opts contains options relating to GOSTR3411_2012_256Opts.
type GOSTR3411_2012_256Opts struct {
}

// Algorithm returns the hash algorithm identifier (to be used).
func (opts *GOSTR3411_2012_256Opts) Algorithm() string {
	return GOSTR3411_2012_256
}

// GOSTR3411_2012_512Opts contains options relating to GOSTR3411_2012_512Opts.
type GOSTR3411_2012_512Opts struct {
}

// Algorithm returns the hash algorithm identifier (to be used).
func (opts *GOSTR3411_2012_512Opts) Algorithm() string {
	return GOSTR3411_2012_512
}
