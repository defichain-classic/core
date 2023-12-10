package params

// DefiChainBootNodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Defichain Classic network.
var DefiChainBootNodes = []string{
	"enode://6533cf0e8984ac309e507e570aa56be42830f0c7b41aacbb0022a02180475785629d466c9eab645d780047324dc4404a04ef67c3fb73e35712439398dfc2e283@95.179.166.25:30303",  // Frankfurt
	"enode://0ed95d6d8a8eedb251b6b5b7bbf44c8c276555ec22b8483a06cd2e9a1a4e8d9cbf270c38cc8d45432b649509af89d91ceba43c37e04978c75aa67bd0b3c442e9@70.34.205.196:30303",  // Stockholm
	"enode://7717489274bffc9cdceb635a911f2aec86ddfdee6bb14a2fba281ffaf142d1fa1e7ddb7e21dedc35380df67c7e94a48f7910dd618e2e532f9f61dcce90009067@207.148.72.102:30303", // Singapore
	"enode://a760d61976e506f465e11aa7145de14764f260af52614e6d67cde1c77de0bcfa1d2324f6f4f348edb2bd28b117fc8d71f6856337b2fada21db7eeb1b253b294f@45.63.28.152:30303",   // Sydney
}

// Once DefiChain Classic network has DNS discovery set up,
// this value can be configured.
// var MintMeDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
