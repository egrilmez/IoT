package common

import (
	"time"

	"github.com/brocaar/loraserver/api/as"
	"github.com/brocaar/loraserver/api/nc"
	"github.com/brocaar/loraserver/internal/backend"
	"github.com/brocaar/lorawan"
	"github.com/brocaar/lorawan/band"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// NodeSessionTTL defines the TTL of a node session (will be renewed on each
// activity)
var NodeSessionTTL = time.Hour * 24 * 31

// Band is the ISM band configuration to use
var Band band.Band

// BandName is the name of the used ISM band
var BandName band.Name

// DeduplicationDelay holds the time to wait for uplink de-duplication
var DeduplicationDelay = time.Millisecond * 200

// GetDownlinkDataDelay holds the delay between uplink delivery to the app server and getting the downlink data from the app server (if any)
var GetDownlinkDataDelay = time.Millisecond * 100

// TimeLocation holds the timezone location
var TimeLocation = time.Local

// CreateGatewayOnStats defines if non-existing gateways should be created
// automatically when receiving stats.
var CreateGatewayOnStats = false

// SpreadFactorToRequiredSNRTable contains the required SNR to demodulate a
// LoRa frame for the given spreadfactor.
// These values are taken from the SX1276 datasheet.
var SpreadFactorToRequiredSNRTable = map[int]float64{
	6:  -5,
	7:  -7.5,
	8:  -10,
	9:  -12.5,
	10: -15,
	11: -17.5,
	12: -20,
}

// LogNodeFrames defines if uplink and downlink frames must be logged to
// the database.
var LogNodeFrames bool

// GatewayServerJWTSecret contains the JWT secret used by the gateway API
// server for gateway authentication.
var GatewayServerJWTSecret string

// RedisPool holds the Redis connection pool.
var RedisPool *redis.Pool

// DB holds the PostgreSQL database connection.
var DB *sqlx.DB

// NetID contains the LoRaWAN NetID.
var NetID lorawan.NetID

// Application holds the connection to the application-server.
var Application as.ApplicationServerClient

// Controller holds the connection to the network-controller.
var Controller nc.NetworkControllerClient

// Gateway holds the gateway backend.
var Gateway backend.Gateway
